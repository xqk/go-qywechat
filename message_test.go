package go_qywechat

import "testing"

func TestQyWechatApp_SendMessage(t *testing.T) {
	qyWechatCorpId := ""
	qyWechatRedisHost := ""
	qyWechatRedisPassword := ""
	qyWechatRedisDb := 0
	qyWechat := New(qyWechatCorpId, WithCache(qyWechatRedisHost, qyWechatRedisPassword, qyWechatRedisDb))
	agentSecret := ""
	agentId := int64(0)
	app := qyWechat.WithApp(agentSecret, agentId)
	err := app.sendMessage(
		&Recipient{},
		"",
		map[string]interface{}{},
		true,
	)
	if err != nil {
		return
	}
}
