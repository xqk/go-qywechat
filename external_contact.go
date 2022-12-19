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
