package go_qywechat

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"sync"
	"time"
)

type TokenInfo struct {
	token          string
	expiresIn      time.Duration
	expirationTime time.Time
}

type token struct {
	mutex *sync.RWMutex
	TokenInfo
	lastRefresh  time.Time
	getTokenFunc func() (TokenInfo, error)
}

func (c *QyWechatSystemApp) GetAccessToken() (TokenInfo, error) {
	var accessToken string
	var expiresIn time.Duration

	cacheKey := c.getAccessTokenCacheKey()
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

	return TokenInfo{token: accessToken, expiresIn: expiresIn, expirationTime: expirationTime}, nil
}

func (c *QyWechatSystemApp) getAccessTokenCacheKey() string {
	return fmt.Sprintf("qywechat:accessToken:%s:%s", c.QyWechat.CorpID, c.AppSecret)
}

func (c *QyWechatApp) GetAccessToken() (TokenInfo, error) {
	var accessToken string
	var expiresIn time.Duration

	cacheKey := c.getAccessTokenCacheKey()
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

	return TokenInfo{token: accessToken, expiresIn: expiresIn, expirationTime: expirationTime}, nil
}

func (c *QyWechatApp) getAccessTokenCacheKey() string {
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

// getJSAPITicket 获取 JSAPI_ticket
func (c *QyWechatApp) getJSAPITicket() (TokenInfo, error) {
	get, err := c.execGetJSAPITicket(reqJSAPITicket{})
	if err != nil {
		return TokenInfo{}, err
	}
	return TokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// getJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *QyWechatApp) getJSAPITicketAgentConfig() (TokenInfo, error) {
	get, err := c.execGetJSAPITicketAgentConfig(reqJSAPITicketAgentConfig{})
	if err != nil {
		return TokenInfo{}, err
	}
	return TokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

func (t *token) syncToken() error {
	get, err := t.getTokenFunc()
	if err != nil {
		return err
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.token = get.token
	t.expiresIn = get.expiresIn * time.Second
	t.expirationTime = get.expirationTime
	t.lastRefresh = time.Now()
	return nil
}

func (t *token) getToken() string {
	// intensive mutex juggling action
	t.mutex.RLock()
	if t.token == "" || time.Now().After(t.expirationTime) {
		t.mutex.RUnlock() // RWMutex doesn't like recursive locking
		// TODO: what to do with the possible error?
		_ = t.syncToken()
		t.mutex.RLock()
	}
	tokenToUse := t.token
	t.mutex.RUnlock()
	return tokenToUse
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

			waitUntilTime := t.lastRefresh.Add(t.expiresIn).Add(-refreshTimeWindow)
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
