package go_qywechat

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type tokenInfo struct {
	token     string
	expiresIn time.Duration
}

type token struct {
	mutex *sync.RWMutex
	tokenInfo
	lastRefresh  time.Time
	getTokenFunc func() (tokenInfo, error)
}

func (c *QyWechatSystemApp) getAccessToken() (tokenInfo, error) {
	var accessToken string
	var expiresIn time.Duration

	cacheKey := c.getAccessTokenCacheKey()

	// 从缓存中获取accessToken
	if c.cache != nil {
		var ctx = context.Background()
		// 获取缓存剩余时间
		ttl, err := c.cache.TTL(ctx, cacheKey).Result()
		if err != nil {
			return tokenInfo{}, err
		}

		ttlSec := int64(ttl.Seconds())
		if ttlSec > 0 {
			accessToken, err = c.cache.Get(ctx, cacheKey).Result()
			if err != nil {
				return tokenInfo{}, err
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
			return tokenInfo{}, err
		}

		accessToken = get.AccessToken
		expiresIn = time.Duration(get.ExpiresInSecs)

		// 设置accessToken缓存
		if c.cache != nil {
			var ctx = context.Background()
			_, err = c.cache.SetNX(ctx, cacheKey, accessToken, expiresIn*time.Second).Result()
			if err != nil {
				return tokenInfo{}, err
			}
		}
	}

	return tokenInfo{token: accessToken, expiresIn: expiresIn}, nil
}

func (c *QyWechatSystemApp) getAccessTokenCacheKey() string {
	return fmt.Sprintf("qywechat:accessToken:%s:%s", c.QyWechat.CorpID, c.AppSecret)
}

func (c *QyWechatApp) getAccessToken() (tokenInfo, error) {
	var accessToken string
	var expiresIn time.Duration

	cacheKey := c.getAccessTokenCacheKey()

	// 从缓存中获取accessToken
	if c.cache != nil {
		var ctx = context.Background()
		// 获取缓存剩余时间
		ttl, err := c.cache.TTL(ctx, cacheKey).Result()
		if err != nil {
			return tokenInfo{}, err
		}

		ttlSec := int64(ttl.Seconds())
		if ttlSec > 0 {
			accessToken, err = c.cache.Get(ctx, cacheKey).Result()
			if err != nil {
				return tokenInfo{}, err
			}
			expiresIn = time.Duration(ttlSec * int64(time.Second))
		}
	}

	// 请求微信获取
	if accessToken == "" {
		get, err := c.execGetAccessToken(reqAccessToken{
			CorpID:     c.CorpID,
			CorpSecret: c.CorpSecret,
		})
		if err != nil {
			return tokenInfo{}, err
		}

		accessToken = get.AccessToken
		expiresIn = time.Duration(get.ExpiresInSecs * int64(time.Second))

		// 设置accessToken缓存
		if c.cache != nil {
			var ctx = context.Background()
			_, err = c.cache.SetNX(ctx, cacheKey, accessToken, expiresIn).Result()
			if err != nil {
				return tokenInfo{}, err
			}
		}
	}

	return tokenInfo{token: accessToken, expiresIn: expiresIn}, nil
}

func (c *QyWechatApp) getAccessTokenCacheKey() string {
	return fmt.Sprintf("qywechat:accessToken:%s:%d", c.QyWechat.CorpID, c.AgentID)
}

// getJSAPITicket 获取 JSAPI_ticket
func (c *QyWechatApp) getJSAPITicket() (tokenInfo, error) {
	get, err := c.execGetJSAPITicket(reqJSAPITicket{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// getJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *QyWechatApp) getJSAPITicketAgentConfig() (tokenInfo, error) {
	get, err := c.execGetJSAPITicketAgentConfig(reqJSAPITicketAgentConfig{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
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
	t.lastRefresh = time.Now()
	return nil
}

func (t *token) getToken() string {
	// intensive mutex juggling action
	t.mutex.RLock()
	if t.token == "" {
		t.mutex.RUnlock() // RWMutex doesn't like recursive locking
		// TODO: what to do with the possible error?
		_ = t.syncToken()
		t.mutex.RLock()
	}
	tokenToUse := t.token
	t.mutex.RUnlock()
	return tokenToUse
}

func (t *token) setGetTokenFunc(f func() (tokenInfo, error)) {
	t.getTokenFunc = f
}
