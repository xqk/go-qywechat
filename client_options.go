package go_qywechat

import "net/http"

// DefaultQYAPIHost 默认企业微信 API Host
const DefaultQYAPIHost = "https://qyapi.weixin.qq.com"

type redisConf struct {
	Host     string
	Password string
	Db       int
}

type options struct {
	QYAPIHost string
	HTTP      *http.Client
	Cache     *redisConf
}

// CtorOption 客户端对象构造参数
type CtorOption interface {
	applyTo(*options)
}

// 构造默认配置
func defaultOptions() options {
	return options{
		QYAPIHost: DefaultQYAPIHost,
		HTTP:      &http.Client{},
	}
}

type withQYAPIHost struct {
	x string
}

// WithQYAPIHost 覆盖默认企业微信 API 域名
func WithQYAPIHost(host string) CtorOption {
	return &withQYAPIHost{x: host}
}

var _ CtorOption = (*withQYAPIHost)(nil)

func (x *withQYAPIHost) applyTo(y *options) {
	y.QYAPIHost = x.x
}

type withHTTPClient struct {
	x *http.Client
}

// WithHTTPClient 使用给定的 http.Client 作为 HTTP 客户端
func WithHTTPClient(client *http.Client) CtorOption {
	return &withHTTPClient{x: client}
}

var _ CtorOption = (*withHTTPClient)(nil)

func (x *withHTTPClient) applyTo(y *options) {
	y.HTTP = x.x
}

type withCache struct {
	x *redisConf
}

// WithCache 使用给定的 redis 作为 缓存 客户端
func WithCache(host, password string, db int) CtorOption {
	return &withCache{x: &redisConf{
		Host:     host,
		Password: password,
		Db:       db,
	}}
}

var _ CtorOption = (*withCache)(nil)

func (x *withCache) applyTo(y *options) {
	y.Cache = x.x
}
