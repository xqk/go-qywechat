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

// execListFollowUserExternalContact 获取配置了客户联系功能的成员列表
func (c *QyWechatSystemApp) execListFollowUserExternalContact(req reqListFollowUserExternalContact) (respListFollowUserExternalContact, error) {
	var resp respListFollowUserExternalContact
	err := c.executeQyapiGet("/cgi-bin/externalcontact/get_follow_user_list", req, &resp, true)
	if err != nil {
		return respListFollowUserExternalContact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respListFollowUserExternalContact{}, bizErr
	}

	return resp, nil
}

// execUserGet 读取成员
func (c *QyWechatSystemApp) execUserGet(req reqUserGet) (respUserGet, error) {
	var resp respUserGet
	err := c.executeQyapiGet("/cgi-bin/user/get", req, &resp, true)
	if err != nil {
		return respUserGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserGet{}, bizErr
	}

	return resp, nil
}

// execUserList 获取部门成员详情
func (c *QyWechatSystemApp) execUserList(req reqUserList) (respUserList, error) {
	var resp respUserList
	err := c.executeQyapiGet("/cgi-bin/user/list", req, &resp, true)
	if err != nil {
		return respUserList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserList{}, bizErr
	}

	return resp, nil
}

// execUserIDByMobile 手机号获取userid
func (c *QyWechatSystemApp) execUserIDByMobile(req reqUserIDByMobile) (respUserIDByMobile, error) {
	var resp respUserIDByMobile
	err := c.executeQyapiJSONPost("/cgi-bin/user/getuserid", req, &resp, true)
	if err != nil {
		return respUserIDByMobile{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserIDByMobile{}, bizErr
	}

	return resp, nil
}

// execDeptList 获取部门列表
func (c *QyWechatSystemApp) execDeptList(req reqDeptList) (respDeptList, error) {
	var resp respDeptList
	err := c.executeQyapiGet("/cgi-bin/department/list", req, &resp, true)
	if err != nil {
		return respDeptList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respDeptList{}, bizErr
	}

	return resp, nil
}

// execUserInfoGet 获取访问用户身份
func (c *QyWechatSystemApp) execUserInfoGet(req reqUserInfoGet) (respUserInfoGet, error) {
	var resp respUserInfoGet
	err := c.executeQyapiGet("/cgi-bin/user/getuserinfo", req, &resp, true)
	if err != nil {
		return respUserInfoGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respUserInfoGet{}, bizErr
	}

	return resp, nil
}
