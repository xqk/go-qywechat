package go_qywechat

// execGetAccessToken 获取access_token
func (c *QyWechatSystemApp) execGetAccessToken(req reqAccessToken) (respAccessToken, error) {
	var resp respAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return respAccessToken{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAccessToken{}, bizErr
	}

	return resp, nil
}

// execGetAccessToken 获取access_token
func (c *QyWechatApp) execGetAccessToken(req reqAccessToken) (respAccessToken, error) {
	var resp respAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return respAccessToken{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respAccessToken{}, bizErr
	}

	return resp, nil
}

// execGetJSAPITicket 获取企业的jsapi_ticket
func (c *QyWechatApp) execGetJSAPITicket(req reqJSAPITicket) (respJSAPITicket, error) {
	var resp respJSAPITicket
	err := c.executeQyapiGet("/cgi-bin/get_jsapi_ticket", req, &resp, true)
	if err != nil {
		return respJSAPITicket{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSAPITicket{}, bizErr
	}

	return resp, nil
}

// execGetJSAPITicketAgentConfig 获取应用的jsapi_ticket
func (c *QyWechatApp) execGetJSAPITicketAgentConfig(req reqJSAPITicketAgentConfig) (respJSAPITicket, error) {
	var resp respJSAPITicket
	err := c.executeQyapiGet("/cgi-bin/ticket/get", req, &resp, true)
	if err != nil {
		return respJSAPITicket{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSAPITicket{}, bizErr
	}

	return resp, nil
}

// execJSCode2Session 临时登录凭证校验code2Session
func (c *QyWechatApp) execJSCode2Session(req reqJSCode2Session) (respJSCode2Session, error) {
	var resp respJSCode2Session
	err := c.executeQyapiGet("/cgi-bin/miniprogram/jscode2session", req, &resp, true)
	if err != nil {
		return respJSCode2Session{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respJSCode2Session{}, bizErr
	}

	return resp, nil
}

// execExternalContactBatchList 批量获取客户详情
func (c *QyWechatSystemApp) execExternalContactBatchList(req reqExternalContactBatchList) (respExternalContactBatchList, error) {
	var resp respExternalContactBatchList
	err := c.executeQyapiJSONPost("/cgi-bin/externalcontact/batch/get_by_user", req, &resp, true)
	if err != nil {
		return respExternalContactBatchList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactBatchList{}, bizErr
	}

	return resp, nil
}

// execExternalContactGet 获取客户详情
func (c *QyWechatSystemApp) execExternalContactGet(req reqExternalContactGet) (respExternalContactGet, error) {
	var resp respExternalContactGet
	err := c.executeQyapiGet("/cgi-bin/externalcontact/get", req, &resp, true)
	if err != nil {
		return respExternalContactGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactGet{}, bizErr
	}

	return resp, nil
}

// execExternalContactList 获取客户列表
func (c *QyWechatSystemApp) execExternalContactList(req reqExternalContactList) (respExternalContactList, error) {
	var resp respExternalContactList
	err := c.executeQyapiGet("/cgi-bin/externalcontact/list", req, &resp, true)
	if err != nil {
		return respExternalContactList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respExternalContactList{}, bizErr
	}

	return resp, nil
}
