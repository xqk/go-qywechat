package go_qywechat

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

// ITokenProvider 是鉴权 token 的外部提供者需要实现的 interface。可用于官方所谓
// 使用“中控服务”集中提供、刷新 token 的场景。
//
// 不同类型的 tokens（如 access token、JSAPI token 等）都是这个 interface 提供，
// 实现方需要自行掌握 token 的类别，避免在 client 构造函数的选项中传入错误的种类。
type ITokenProvider interface {
	// GetToken 取回一个 token。有可能被并发调用。
	GetToken(context.Context) (string, error)
}

type TokenInfo struct {
	Token          string
	ExpiresIn      time.Duration
	ExpirationTime time.Time
}

type token struct {
	mutex *sync.RWMutex
	TokenInfo
	lastRefresh      time.Time
	getTokenFunc     func() (TokenInfo, error)
	externalProvider ITokenProvider
	Cache            *redis.Client
	CacheKey         string
}

func (c *QyWechatSystemApp) GetAccessToken() (TokenInfo, error) {
	var accessToken string
	var expiresIn time.Duration

	cacheKey := c.GetAccessTokenCacheKey()
	currentTime := time.Now()

	// 从缓存中获取accessToken
	if c.cache != nil {
		var ctx = context.Background()
		// 获取缓存剩余时间
		ttl, err := c.cache.TTL(ctx, cacheKey).Result()
		if err != nil {
			return TokenInfo{}, err
		}

		ttlSec := int64(ttl.Seconds())
		if ttlSec > 0 {
			accessToken, err = c.cache.Get(ctx, cacheKey).Result()
			if err != nil {
				return TokenInfo{}, err
			}
			expiresIn = time.Duration(ttlSec)
		}
	}

	// 请求微信获取
	if accessToken == "" {
		get, err := c.execGetAccessToken(reqAccessToken{
			CorpID:     c.CorpID,
			CorpSecret: c.AppSecret,
		})
		if err != nil {
			return TokenInfo{}, err
		}

		accessToken = get.AccessToken
		expiresIn = time.Duration(get.ExpiresInSecs)

		// 设置accessToken缓存
		if c.cache != nil {
			var ctx = context.Background()
			_, err = c.cache.SetNX(ctx, cacheKey, accessToken, expiresIn*time.Second).Result()
			if err != nil {
				return TokenInfo{}, err
			}
		}
	}

	expirationTime := currentTime.Add(expiresIn * time.Second)

	return TokenInfo{Token: accessToken, ExpiresIn: expiresIn, ExpirationTime: expirationTime}, nil
}

func (c *QyWechatSystemApp) GetAccessTokenCacheKey() string {
	return fmt.Sprintf("qywechat:accessToken:%s:%s", c.QyWechat.CorpID, c.AppSecret)
}

func (c *QyWechatApp) GetAccessToken() (TokenInfo, error) {
	var accessToken string
	var expiresIn time.Duration

	cacheKey := c.GetAccessTokenCacheKey()
	currentTime := time.Now()

	// 从缓存中获取accessToken
	if c.cache != nil {
		var ctx = context.Background()
		// 获取缓存剩余时间
		ttl, err := c.cache.TTL(ctx, cacheKey).Result()
		if err != nil {
			return TokenInfo{}, err
		}

		ttlSec := int64(ttl.Seconds())
		if ttlSec > 0 {
			accessToken, err = c.cache.Get(ctx, cacheKey).Result()
			if err != nil {
				return TokenInfo{}, err
			}
			expiresIn = time.Duration(ttlSec)
		}
	}

	// 请求微信获取
	if accessToken == "" {
		get, err := c.execGetAccessToken(reqAccessToken{
			CorpID:     c.CorpID,
			CorpSecret: c.CorpSecret,
		})
		if err != nil {
			return TokenInfo{}, err
		}

		accessToken = get.AccessToken
		expiresIn = time.Duration(get.ExpiresInSecs)

		// 设置accessToken缓存
		if c.cache != nil {
			var ctx = context.Background()
			_, err = c.cache.SetNX(ctx, cacheKey, accessToken, expiresIn*time.Second).Result()
			if err != nil {
				return TokenInfo{}, err
			}
		}
	}

	expirationTime := currentTime.Add(expiresIn * time.Second)

	return TokenInfo{Token: accessToken, ExpiresIn: expiresIn, ExpirationTime: expirationTime}, nil
}

func (c *QyWechatApp) GetAccessTokenCacheKey() string {
	return fmt.Sprintf("qywechat:accessToken:%s:%d", c.QyWechat.CorpID, c.AgentID)
}

// SpawnAccessTokenRefresher 启动该 app 的 access token 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *QyWechatSystemApp) SpawnAccessTokenRefresher() {
	ctx := context.Background()
	c.SpawnAccessTokenRefresherWithContext(ctx)
}

// SpawnAccessTokenRefresherWithContext 启动该 app 的 access token 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *QyWechatSystemApp) SpawnAccessTokenRefresherWithContext(ctx context.Context) {
	go c.accessToken.tokenRefresher(ctx)
}

// SpawnAccessTokenRefresher 启动该 app 的 access token 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *QyWechatApp) SpawnAccessTokenRefresher() {
	ctx := context.Background()
	c.SpawnAccessTokenRefresherWithContext(ctx)
}

// SpawnAccessTokenRefresherWithContext 启动该 app 的 access token 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *QyWechatApp) SpawnAccessTokenRefresherWithContext(ctx context.Context) {
	go c.accessToken.tokenRefresher(ctx)
}

// GetJSAPITicket 获取 JSAPI_ticket
func (c *QyWechatApp) GetJSAPITicket() (string, error) {
	return c.jsapiTicket.getToken()
}

// getJSAPITicket 获取 JSAPI_ticket
func (c *QyWechatApp) getJSAPITicket() (TokenInfo, error) {
	get, err := c.execGetJSAPITicket(reqJSAPITicket{})
	if err != nil {
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.Ticket, ExpiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// GetJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *QyWechatApp) GetJSAPITicketAgentConfig() (string, error) {
	return c.jsapiTicketAgentConfig.getToken()
}

// getJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *QyWechatApp) getJSAPITicketAgentConfig() (TokenInfo, error) {
	get, err := c.execGetJSAPITicketAgentConfig(reqJSAPITicketAgentConfig{})
	if err != nil {
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.Ticket, ExpiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

func (t *token) syncToken() error {
	get, err := t.getTokenFunc()
	if err != nil {
		return err
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.Token = get.Token
	t.ExpiresIn = get.ExpiresIn * time.Second
	t.ExpirationTime = get.ExpirationTime
	t.lastRefresh = time.Now()
	return nil
}

func (t *token) getToken() (string, error) {
	if t.externalProvider != nil {
		tok, err := t.externalProvider.GetToken(context.TODO())
		if err != nil {
			return "", err
		}
		return tok, nil
	}

	// intensive mutex juggling action
	t.mutex.RLock()
	var accessToken string
	var expiresIn time.Duration
	currentTime := time.Now()
	var step string
	if t.Cache != nil && t.CacheKey != "" { // 有缓存和缓存键（仅限accessToken 禁止jsApiToken等缓存）
		var ctx = context.Background()
		// 获取缓存剩余时间
		ttl, err := t.Cache.TTL(ctx, t.CacheKey).Result()
		if err != nil {
			step = "1"
			return "", err
		}
		step = "2"

		ttlSec := int64(ttl.Seconds())
		if ttlSec > 0 {
			accessToken, err = t.Cache.Get(ctx, t.CacheKey).Result()
			if err != nil {
				step = "3"
				return "", err
			}
			step = "4"
			expiresIn = time.Duration(ttlSec)
		}
	}
	if accessToken == "" {
		if t.Cache != nil && t.CacheKey != "" {
			step = "5"
			t.mutex.RUnlock() // RWMutex doesn't like recursive locking
			err := t.syncToken()
			if err != nil {
				return "", err
			}
			t.mutex.RLock()
		} else if t.Token == "" || time.Now().After(t.ExpirationTime) {
			step = "6"
			t.mutex.RUnlock() // RWMutex doesn't like recursive locking
			err := t.syncToken()
			if err != nil {
				return "", err
			}
			t.mutex.RLock()
		}
	} else {
		step = "7"
		t.Token = accessToken
		t.ExpirationTime = currentTime.Add(expiresIn * time.Second)
		t.ExpiresIn = expiresIn
		t.lastRefresh = currentTime.Add((7200 - expiresIn) * time.Second)
	}

	tokenToUse := t.Token
	println("😄 getToken t.Token", t.Cache != nil, t.CacheKey != "", tokenToUse, step)
	t.mutex.RUnlock()
	return tokenToUse, nil
}

func (t *token) tokenRefresher(ctx context.Context) {
	const refreshTimeWindow = 30 * time.Minute
	const minRefreshDuration = 5 * time.Second

	var waitDuration time.Duration = 0
	for {
		select {
		case <-time.After(waitDuration):
			retryer := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
			if err := backoff.Retry(t.syncToken, retryer); err != nil {
				// TODO: logging
				_ = err
			}

			waitUntilTime := t.lastRefresh.Add(t.ExpiresIn).Add(-refreshTimeWindow)
			waitDuration = time.Until(waitUntilTime)
			if waitDuration < minRefreshDuration {
				waitDuration = minRefreshDuration
			}
		case <-ctx.Done():
			return
		}
	}
}

func (t *token) setGetTokenFunc(f func() (TokenInfo, error)) {
	t.getTokenFunc = f
}

// JSCode2Session 临时登录凭证校验
func (c *QyWechatApp) JSCode2Session(jscode string) (*JSCodeSession, error) {
	resp, err := c.execJSCode2Session(reqJSCode2Session{JSCode: jscode})
	if err != nil {
		return nil, err
	}
	return &resp.JSCodeSession, nil
}
