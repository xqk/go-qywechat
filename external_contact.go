package go_qywechat

import "time"

// ExternalContactUnassigned 离职成员的客户
type ExternalContactUnassigned struct {
	// HandoverUserID 离职成员的userid
	HandoverUserID string
	// ExternalUserID 外部联系人userid
	ExternalUserID string
	// DemissionTime 成员离职时间
	DemissionTime time.Time
}

// BatchListExternalContact 批量获取客户详情
func (c *QyWechatSystemApp) BatchListExternalContact(userIDs []string, cursor string, limit int) (*BatchListExternalContactsResp, error) {
	resp, err := c.execExternalContactBatchList(reqExternalContactBatchList{
		UserIDs: userIDs,
		Cursor:  cursor,
		Limit:   limit,
	})
	if err != nil {
		return nil, err
	}
	return &BatchListExternalContactsResp{Result: resp.ExternalContactList, NextCursor: resp.NextCursor}, nil
}

// GetExternalContact 获取客户详情
func (c *QyWechatSystemApp) GetExternalContact(externalUserID string) (*ExternalContactInfo, error) {
	resp, err := c.execExternalContactGet(reqExternalContactGet{
		ExternalUserID: externalUserID,
	})
	if err != nil {
		return nil, err
	}
	return &resp.ExternalContactInfo, nil
}

// ListExternalContact 获取客户列表
func (c *QyWechatSystemApp) ListExternalContact(userID string) ([]string, error) {
	resp, err := c.execExternalContactList(reqExternalContactList{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExternalUserID, nil
}

// ExternalContactListFollowUser 获取配置了客户联系功能的成员列表
func (c *QyWechatSystemApp) ExternalContactListFollowUser() (*ExternalContactFollowUserList, error) {
	resp, err := c.execListFollowUserExternalContact(reqListFollowUserExternalContact{})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactFollowUserList, nil
}

// ExternalContactAddContact 配置客户联系「联系我」方式
func (c *QyWechatSystemApp) ExternalContactAddContact(t int, scene int, style int, remark string, skipVerify bool, state string, user []string, party []int, isTemp bool, expiresIn int, chatExpiresIn int, unionID string, conclusions Conclusions) (*ExternalContactAddContact, error) {
	resp, err := c.execAddContactExternalContact(
		reqAddContactExternalContact{
			ExternalContactWay{
				Type:          t,
				Scene:         scene,
				Style:         style,
				Remark:        remark,
				SkipVerify:    skipVerify,
				State:         state,
				User:          user,
				Party:         party,
				IsTemp:        isTemp,
				ExpiresIn:     expiresIn,
				ChatExpiresIn: chatExpiresIn,
				UnionID:       unionID,
				Conclusions:   conclusions,
			},
		})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactAddContact, nil
}

// ExternalContactGetContactWay 获取企业已配置的「联系我」方式
func (c *QyWechatSystemApp) ExternalContactGetContactWay(configID string) (*ExternalContactContactWay, error) {
	resp, err := c.execGetContactWayExternalContact(reqGetContactWayExternalContact{ConfigID: configID})
	if err != nil {
		return nil, err
	}

	return &resp.ContactWay, nil
}

// ExternalContactListContactWayChat 获取企业已配置的「联系我」列表
func (c *QyWechatSystemApp) ExternalContactListContactWayChat(startTime int, endTime int, cursor string, limit int) (*ExternalContactListContactWayChat, error) {
	resp, err := c.execListContactWayChatExternalContact(reqListContactWayExternalContact{
		StartTime: startTime,
		EndTime:   endTime,
		Cursor:    cursor,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactListContactWayChat, nil
}

// ExternalContactUpdateContactWay 更新企业已配置的「联系我」成员配置
func (c *QyWechatSystemApp) ExternalContactUpdateContactWay(configID string, remark string, skipVerify bool, style int, state string, user []string, party []int, expiresIn int, chatExpiresIn int, unionid string, conclusions Conclusions) error {
	_, err := c.execUpdateContactWayExternalContact(reqUpdateContactWayExternalContact{
		ConfigID:      configID,
		Remark:        remark,
		SkipVerify:    skipVerify,
		Style:         style,
		State:         state,
		User:          user,
		Party:         party,
		ExpiresIn:     expiresIn,
		ChatExpiresIn: chatExpiresIn,
		UnionID:       unionid,
		Conclusions:   conclusions,
	})

	return err
}

// ExternalContactDelContactWay 删除企业已配置的「联系我」方式
func (c *QyWechatSystemApp) ExternalContactDelContactWay(configID string) error {
	_, err := c.execDelContactWayExternalContact(reqDelContactWayExternalContact{ConfigID: configID})

	return err
}

// ExternalContactCloseTempChat 结束临时会话
func (c *QyWechatSystemApp) ExternalContactCloseTempChat(userID, externalUserID string) error {
	_, err := c.execCloseTempChatExternalContact(reqCloseTempChatExternalContact{
		UserID:         userID,
		ExternalUserID: externalUserID,
	})

	return err
}

// ----

// BatchListExternalContact 批量获取客户详情
func (c *QyWechatApp) BatchListExternalContact(userIDs []string, cursor string, limit int) (*BatchListExternalContactsResp, error) {
	resp, err := c.execExternalContactBatchList(reqExternalContactBatchList{
		UserIDs: userIDs,
		Cursor:  cursor,
		Limit:   limit,
	})
	if err != nil {
		return nil, err
	}
	return &BatchListExternalContactsResp{Result: resp.ExternalContactList, NextCursor: resp.NextCursor}, nil
}

// GetExternalContact 获取客户详情
func (c *QyWechatApp) GetExternalContact(externalUserID string) (*ExternalContactInfo, error) {
	resp, err := c.execExternalContactGet(reqExternalContactGet{
		ExternalUserID: externalUserID,
	})
	if err != nil {
		return nil, err
	}
	return &resp.ExternalContactInfo, nil
}

// ListExternalContact 获取客户列表
func (c *QyWechatApp) ListExternalContact(userID string) ([]string, error) {
	resp, err := c.execExternalContactList(reqExternalContactList{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExternalUserID, nil
}

// ExternalContactListFollowUser 获取配置了客户联系功能的成员列表
func (c *QyWechatApp) ExternalContactListFollowUser() (*ExternalContactFollowUserList, error) {
	resp, err := c.execListFollowUserExternalContact(reqListFollowUserExternalContact{})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactFollowUserList, nil
}

// ExternalContactAddContact 配置客户联系「联系我」方式
func (c *QyWechatApp) ExternalContactAddContact(t int, scene int, style int, remark string, skipVerify bool, state string, user []string, party []int, isTemp bool, expiresIn int, chatExpiresIn int, unionID string, conclusions Conclusions) (*ExternalContactAddContact, error) {
	resp, err := c.execAddContactExternalContact(
		reqAddContactExternalContact{
			ExternalContactWay{
				Type:          t,
				Scene:         scene,
				Style:         style,
				Remark:        remark,
				SkipVerify:    skipVerify,
				State:         state,
				User:          user,
				Party:         party,
				IsTemp:        isTemp,
				ExpiresIn:     expiresIn,
				ChatExpiresIn: chatExpiresIn,
				UnionID:       unionID,
				Conclusions:   conclusions,
			},
		})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactAddContact, nil
}

// ExternalContactGetContactWay 获取企业已配置的「联系我」方式
func (c *QyWechatApp) ExternalContactGetContactWay(configID string) (*ExternalContactContactWay, error) {
	resp, err := c.execGetContactWayExternalContact(reqGetContactWayExternalContact{ConfigID: configID})
	if err != nil {
		return nil, err
	}

	return &resp.ContactWay, nil
}

// ExternalContactListContactWayChat 获取企业已配置的「联系我」列表
func (c *QyWechatApp) ExternalContactListContactWayChat(startTime int, endTime int, cursor string, limit int) (*ExternalContactListContactWayChat, error) {
	resp, err := c.execListContactWayChatExternalContact(reqListContactWayExternalContact{
		StartTime: startTime,
		EndTime:   endTime,
		Cursor:    cursor,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactListContactWayChat, nil
}

// ExternalContactUpdateContactWay 更新企业已配置的「联系我」成员配置
func (c *QyWechatApp) ExternalContactUpdateContactWay(configID string, remark string, skipVerify bool, style int, state string, user []string, party []int, expiresIn int, chatExpiresIn int, unionid string, conclusions Conclusions) error {
	_, err := c.execUpdateContactWayExternalContact(reqUpdateContactWayExternalContact{
		ConfigID:      configID,
		Remark:        remark,
		SkipVerify:    skipVerify,
		Style:         style,
		State:         state,
		User:          user,
		Party:         party,
		ExpiresIn:     expiresIn,
		ChatExpiresIn: chatExpiresIn,
		UnionID:       unionid,
		Conclusions:   conclusions,
	})

	return err
}

// ExternalContactDelContactWay 删除企业已配置的「联系我」方式
func (c *QyWechatApp) ExternalContactDelContactWay(configID string) error {
	_, err := c.execDelContactWayExternalContact(reqDelContactWayExternalContact{ConfigID: configID})

	return err
}

// ExternalContactCloseTempChat 结束临时会话
func (c *QyWechatApp) ExternalContactCloseTempChat(userID, externalUserID string) error {
	_, err := c.execCloseTempChatExternalContact(reqCloseTempChatExternalContact{
		UserID:         userID,
		ExternalUserID: externalUserID,
	})

	return err
}
