package go_qywechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/url"
	"sync"
)

// QyWechat 企业微信客户端
type QyWechat struct {
	opts  options
	cache *redis.Client

	CorpID string // CorpID 企业 ID，必填
}

// QyWechatSystemApp 企业微信客户端（系统应用:通信录、外部联系人等）
type QyWechatSystemApp struct {
	*QyWechat

	AppSecret string // 系统应用的密钥

	accessToken *token
}

// QyWechatApp 企业微信客户端（分应用）
type QyWechatApp struct {
	*QyWechat

	// CorpSecret 应用的凭证密钥，必填
	CorpSecret string

	// AgentID 应用 ID，必填
	AgentID                int64
	accessToken            *token
	jsapiTicket            *token
	jsapiTicketAgentConfig *token
}

// New 构造一个 QyWechat 客户端对象，需要提供企业 ID
func New(corpId string, opts ...CtorOption) *QyWechat {
	optionsObj := defaultOptions()

	for _, o := range opts {
		o.applyTo(&optionsObj)
	}

	var rdb *redis.Client
	if optionsObj.Cache != nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     optionsObj.Cache.Host,
			Password: optionsObj.Cache.Password, // no password set
			DB:       optionsObj.Cache.Db,       // use default DB
		})
	}

	return &QyWechat{
		opts:  optionsObj,
		cache: rdb,

		CorpID: corpId,
	}
}

// WithSystemApp 构造本企业下系统app的客户端（系统应用:通信录、外部联系人等）
func (c *QyWechat) WithSystemApp(appSecret string) *QyWechatSystemApp {
	app := QyWechatSystemApp{
		QyWechat: c,

		AppSecret: appSecret,

		accessToken: &token{mutex: &sync.RWMutex{}},
	}
	app.accessToken.setGetTokenFunc(app.GetAccessToken)
	return &app
}

// WithApp 构造本企业下某自建 app 的客户端
func (c *QyWechat) WithApp(corpSecret string, agentID int64) *QyWechatApp {
	app := QyWechatApp{
		QyWechat: c,

		CorpSecret: corpSecret,
		AgentID:    agentID,

		accessToken:            &token{mutex: &sync.RWMutex{}},
		jsapiTicket:            &token{mutex: &sync.RWMutex{}},
		jsapiTicketAgentConfig: &token{mutex: &sync.RWMutex{}},
	}
	app.accessToken.setGetTokenFunc(app.GetAccessToken)
	app.jsapiTicket.setGetTokenFunc(app.getJSAPITicket)
	app.jsapiTicketAgentConfig.setGetTokenFunc(app.getJSAPITicketAgentConfig)
	return &app
}

func (c *QyWechatSystemApp) executeQyapiGet(path string, req urlValuer, respObj interface{}, withAccessToken bool) error {
	qyUrl := c.composeQyapiURLWithToken(path, req, withAccessToken)
	urlStr := qyUrl.String()

	resp, err := c.opts.HTTP.Get(urlStr)
	if err != nil {
		// TODO: error_chain
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}

func (c *QyWechatSystemApp) composeQyapiURLWithToken(path string, req interface{}, withAccessToken bool) *url.URL {
	qyUrl := c.composeQyapiURL(path, req)

	if !withAccessToken {
		return qyUrl
	}

	q := qyUrl.Query()
	q.Set("access_token", c.accessToken.getToken())
	qyUrl.RawQuery = q.Encode()

	return qyUrl
}

func (c *QyWechatSystemApp) composeQyapiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	// TODO: refactor
	base, err := url.Parse(c.opts.QYAPIHost)
	if err != nil {
		// TODO: error_chain
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", c.opts.QYAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

func (c *QyWechatApp) executeQyapiGet(path string, req urlValuer, respObj interface{}, withAccessToken bool) error {
	qyUrl := c.composeQyapiURLWithToken(path, req, withAccessToken)
	urlStr := qyUrl.String()

	resp, err := c.opts.HTTP.Get(urlStr)
	if err != nil {
		// TODO: error_chain
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}

func (c *QyWechatApp) composeQyapiURLWithToken(path string, req interface{}, withAccessToken bool) *url.URL {
	qyUrl := c.composeQyapiURL(path, req)

	if !withAccessToken {
		return qyUrl
	}

	q := qyUrl.Query()
	q.Set("access_token", c.accessToken.getToken())
	qyUrl.RawQuery = q.Encode()

	return qyUrl
}

func (c *QyWechatApp) composeQyapiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	// TODO: refactor
	base, err := url.Parse(c.opts.QYAPIHost)
	if err != nil {
		// TODO: error_chain
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", c.opts.QYAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

func (c *QyWechatSystemApp) executeQyapiJSONPost(path string, req bodyer, respObj interface{}, withAccessToken bool) error {
	qyUrl := c.composeQyapiURLWithToken(path, req, withAccessToken)
	urlStr := qyUrl.String()

	body, err := req.intoBody()
	if err != nil {
		// TODO: error_chain
		return err
	}

	resp, err := c.opts.HTTP.Post(urlStr, "application/json", bytes.NewReader(body))
	if err != nil {
		// TODO: error_chain
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(respObj)
	if err != nil {
		// TODO: error_chain
		return err
	}

	return nil
}
