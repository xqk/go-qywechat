package go_qywechat

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type reqAccessToken struct {
	CorpID     string
	CorpSecret string
}

var _ urlValuer = reqAccessToken{}

func (x reqAccessToken) intoURLValues() url.Values {
	return url.Values{
		"corpid":     {x.CorpID},
		"corpsecret": {x.CorpSecret},
	}
}

type respCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type respAccessToken struct {
	respCommon

	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}

// IsOK 响应体是否为一次成功请求的响应
//
// 实现依据: https://work.weixin.qq.com/api/doc#10013
//
// > 企业微信所有接口，返回包里都有errcode、errmsg。
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *respCommon) IsOK() bool {
	return x.ErrCode == 0
}

func (x *respCommon) TryIntoErr() error {
	if x.IsOK() {
		return nil
	}

	return &WorkwxClientError{
		Code: x.ErrCode,
		Msg:  x.ErrMsg,
	}
}

type reqJSAPITicketAgentConfig struct{}

var _ urlValuer = reqJSAPITicketAgentConfig{}

func (x reqJSAPITicketAgentConfig) intoURLValues() url.Values {
	return url.Values{
		"type": {"agent_config"},
	}
}

type reqJSAPITicket struct{}

var _ urlValuer = reqJSAPITicket{}

func (x reqJSAPITicket) intoURLValues() url.Values {
	return url.Values{}
}

type respJSAPITicket struct {
	respCommon

	Ticket        string `json:"ticket"`
	ExpiresInSecs int64  `json:"expires_in"`
}

// respExternalContactList 获取客户列表
type respExternalContactList struct {
	respCommon

	ExternalUserID []string `json:"external_userid"`
}

// reqExternalContactGet 获取客户详情
type reqExternalContactGet struct {
	ExternalUserID string `json:"external_userid"`
}

var _ urlValuer = reqExternalContactGet{}

func (x reqExternalContactGet) intoURLValues() url.Values {
	return url.Values{
		"external_userid": {x.ExternalUserID},
	}
}

// respExternalContactGet 获取客户详情
type respExternalContactGet struct {
	respCommon
	ExternalContactInfo
}

// ExternalContactInfo 外部联系人信息
type ExternalContactInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

// ExternalContactBatchInfo 外部联系人信息
type ExternalContactBatchInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowInfo      `json:"follow_info"`
}

// BatchListExternalContactsResp 外部联系人信息
type BatchListExternalContactsResp struct {
	Result     []ExternalContactBatchInfo
	NextCursor string
}

// reqExternalContactBatchList 批量获取客户详情
type reqExternalContactBatchList struct {
	UserIDs []string `json:"userid_list"`
	Cursor  string   `json:"cursor"`
	Limit   int      `json:"limit"`
}

var _ bodyer = reqExternalContactBatchList{}

func (x reqExternalContactBatchList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respExternalContactBatchList 批量获取客户详情
type respExternalContactBatchList struct {
	respCommon
	NextCursor          string                     `json:"next_cursor"`
	ExternalContactList []ExternalContactBatchInfo `json:"external_contact_list"`
}

// reqExternalContactList 获取客户列表
type reqExternalContactList struct {
	UserID string `json:"userid"`
}

var _ urlValuer = reqExternalContactList{}

func (x reqExternalContactList) intoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// reqJSCode2Session 临时登录凭证校验
type reqJSCode2Session struct {
	JSCode string
}

var _ urlValuer = reqJSCode2Session{}

func (x reqJSCode2Session) intoURLValues() url.Values {
	return url.Values{
		"js_code":    {x.JSCode},
		"grant_type": {"authorization_code"},
	}
}

// respJSCode2Session 临时登录凭证校验
type respJSCode2Session struct {
	respCommon
	JSCodeSession
}

// JSCodeSession 临时登录凭证
type JSCodeSession struct {
	CorpID     string `json:"corpid"`
	UserID     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

type reqListFollowUserExternalContact struct {
}

var _ urlValuer = reqListFollowUserExternalContact{}

func (x reqListFollowUserExternalContact) intoURLValues() url.Values {
	return url.Values{}
}

type respListFollowUserExternalContact struct {
	respCommon
	ExternalContactFollowUserList
}

type reqUserGet struct {
	UserID string
}

var _ urlValuer = reqUserGet{}

func (x reqUserGet) intoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// respUserDetail 成员详细信息的公共字段
type respUserDetail struct {
	UserID         string   `json:"userid"`
	Name           string   `json:"name"`
	DeptIDs        []int64  `json:"department"`
	DeptOrder      []uint32 `json:"order"`
	Position       string   `json:"position"`
	Mobile         string   `json:"mobile"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	IsLeaderInDept []int    `json:"is_leader_in_dept"`
	AvatarURL      string   `json:"avatar"`
	Telephone      string   `json:"telephone"`
	IsEnabled      int      `json:"enable"`
	Alias          string   `json:"alias"`
	Status         int      `json:"status"`
	QRCodeURL      string   `json:"qr_code"`
	// TODO: extattr external_profile external_position
}

// respUserGet 读取成员响应
type respUserGet struct {
	respCommon

	respUserDetail
}

// reqUserList 部门成员请求
type reqUserList struct {
	DeptID     int64
	FetchChild bool
}

var _ urlValuer = reqUserList{}

func (x reqUserList) intoURLValues() url.Values {
	var fetchChild int64
	if x.FetchChild {
		fetchChild = 1
	}

	return url.Values{
		"department_id": {strconv.FormatInt(x.DeptID, 10)},
		"fetch_child":   {strconv.FormatInt(fetchChild, 10)},
	}
}

// respUsersByDeptID 部门成员详情响应
type respUserList struct {
	respCommon

	Users []*respUserDetail `json:"userlist"`
}

// reqUserIDByMobile 手机号获取 userid 请求
type reqUserIDByMobile struct {
	Mobile string `json:"mobile"`
}

var _ bodyer = reqUserIDByMobile{}

func (x reqUserIDByMobile) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// respUserIDByMobile 手机号获取 userid 响应
type respUserIDByMobile struct {
	respCommon

	UserID string `json:"userid"`
}

type reqDeptList struct {
	HaveID bool
	ID     int64
}

var _ urlValuer = reqDeptList{}

func (x reqDeptList) intoURLValues() url.Values {
	if !x.HaveID {
		return url.Values{}
	}

	return url.Values{
		"id": {strconv.FormatInt(x.ID, 10)},
	}
}

// respDeptList 部门列表响应
type respDeptList struct {
	respCommon

	// TODO: 不要懒惰，把 API 层的类型写好
	Department []*DeptInfo `json:"department"`
}

// respExternalContactRemark 获取客户详情
type respExternalContactRemark struct {
	respCommon
}

// reqUserInfoGet 获取访问用户身份
type reqUserInfoGet struct {
	// 通过成员授权获取到的code，最大为512字节。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期。
	Code string
}

var _ urlValuer = reqUserInfoGet{}

func (x reqUserInfoGet) intoURLValues() url.Values {
	return url.Values{
		"code": {x.Code},
	}
}

// respUserInfoGet 部门列表响应
type respUserInfoGet struct {
	respCommon
	UserIdentityInfo
}
