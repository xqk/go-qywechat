package go_qywechat

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

func marshalIntoJSONBody(x interface{}) ([]byte, error) {
	y, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		return nil, makeReqMarshalErr(err)
	}

	return y, nil
}

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

// reqMessage 消息发送请求
type reqMessage struct {
	ToUser  []string
	ToParty []string
	ToTag   []string
	ChatID  string
	AgentID int64
	MsgType string
	Content map[string]interface{}
	IsSafe  bool
}

var _ bodyer = reqMessage{}

func (x reqMessage) intoBody() ([]byte, error) {
	// fuck
	safeInt := 0
	if x.IsSafe {
		safeInt = 1
	}

	obj := map[string]interface{}{
		"msgtype": x.MsgType,
		"agentid": x.AgentID,
		"safe":    safeInt,
	}

	// msgtype polymorphism
	if x.MsgType != "template_card" {
		obj[x.MsgType] = x.Content
	} else {
		obj[x.MsgType] = x.Content["template_card"]
	}

	// 复用这个结构体，因为是 package-private 的所以这么做没风险
	if x.ChatID != "" {
		obj["chatid"] = x.ChatID
	} else {
		obj["touser"] = strings.Join(x.ToUser, "|")
		obj["toparty"] = strings.Join(x.ToParty, "|")
		obj["totag"] = strings.Join(x.ToTag, "|")
	}

	return marshalIntoJSONBody(obj)
}

// respMessageSend 消息发送响应
type respMessageSend struct {
	respCommon

	InvalidUsers   string `json:"invaliduser"`
	InvalidParties string `json:"invalidparty"`
	InvalidTags    string `json:"invalidtag"`
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

type reqAddContactExternalContact struct {
	ExternalContactWay
}

var _ bodyer = reqAddContactExternalContact{}

func (x reqAddContactExternalContact) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type respAddContactExternalContact struct {
	respCommon
	ExternalContactAddContact
}

type ExternalContactAddContact struct {
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}

type reqGetContactWayExternalContact struct {
	ConfigID string `json:"config_id"`
}

var _ bodyer = reqGetContactWayExternalContact{}

func (x reqGetContactWayExternalContact) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type respGetContactWayExternalContact struct {
	respCommon
	ContactWay ExternalContactContactWay `json:"contact_way"`
}

type ExternalContactContactWay struct {
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
	ExternalContactWay
}

var _ bodyer = reqListContactWayExternalContact{}

func (x reqListContactWayExternalContact) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type respListContactWayChatExternalContact struct {
	respCommon
	ExternalContactListContactWayChat
}

type ExternalContactListContactWayChat struct {
	NextCursor string       `json:"next_cursor"`
	ContactWay []contactWay `json:"contact_way"`
}

type contactWay struct {
	ConfigID string `json:"config_id"`
}

var _ bodyer = reqUpdateContactWayExternalContact{}

func (x reqUpdateContactWayExternalContact) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// TaskCardBtn 任务卡片消息按钮
type TaskCardBtn struct {
	// Key 按钮key值，用户点击后，会产生任务卡片回调事件，回调事件会带上该key值，只能由数字、字母和“_-@”组成，最长支持128字节
	Key string `json:"key"`
	// Name 按钮名称
	Name string `json:"name"`
	// ReplaceName 点击按钮后显示的名称，默认为“已处理”
	ReplaceName string `json:"replace_name"`
	// Color 按钮字体颜色，可选“red”或者“blue”,默认为“blue”
	Color string `json:"color"`
	// IsBold 按钮字体是否加粗，默认false
	IsBold bool `json:"is_bold"`
}

// Article news 类型的文章
type Article struct {
	// 标题，不超过128个字节，超过会自动截断（支持id转译）
	Title string `json:"title"`
	// 描述，不超过512个字节，超过会自动截断（支持id转译）
	Description string `json:"description"`
	// 点击后跳转的链接。 最长2048字节，请确保包含了协议头(http/https)，小程序或者url必须填写一个
	URL string `json:"url"`
	// 图文消息的图片链接，最长2048字节，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150
	PicURL string `json:"picurl"`
	// 小程序appid，必须是与当前应用关联的小程序，appid和pagepath必须同时填写，填写后会忽略url字段
	AppID string `json:"appid"`
	// 点击消息卡片后的小程序页面，最长128字节，仅限本小程序内的页面。appid和pagepath必须同时填写，填写后会忽略url字段
	PagePath string `json:"pagepath"`
}

// MPArticle mpnews 类型的文章
type MPArticle struct {
	// 标题，不超过128个字节，超过会自动截断（支持id转译）
	Title string `json:"title"`
	// 图文消息缩略图的media_id, 可以通过素材管理接口获得。此处thumb_media_id即上传接口返回的media_id
	ThumbMediaID string `json:"thumb_media_id"`
	// 图文消息的作者，不超过64个字节
	Author string `json:"author"`
	// 图文消息点击“阅读原文”之后的页面链接
	ContentSourceURL string `json:"content_source_url"`
	// 图文消息的内容，支持html标签，不超过666 K个字节（支持id转译）
	Content string `json:"content"`
	// 图文消息的描述，不超过512个字节，超过会自动截断（支持id转译）
	Digest string `json:"digest"`
}

// Source 卡片来源样式信息，不需要来源样式可不填写
type Source struct {
	// 来源图片的url，来源图片的尺寸建议为72*72
	IconURL string `json:"icon_url"`
	// 来源图片的描述，建议不超过20个字，（支持id转译）
	Desc string `json:"desc"`
	// 来源文字的颜色，目前支持：0(默认) 灰色，1 黑色，2 红色，3 绿色
	DescColor int `json:"desc_color"`
}

// ActionList 操作列表，列表长度取值范围为 [1, 3]
type ActionList struct {
	// 操作的描述文案
	Text string `json:"text"`
	// 操作key值，用户点击后，会产生回调事件将本参数作为EventKey返回，回调事件会带上该key值，最长支持1024字节，不可重复
	Key string `json:"key"`
}

// ActionMenu 卡片右上角更多操作按钮
type ActionMenu struct {
	// 更多操作界面的描述
	Desc       string       `json:"desc"`
	ActionList []ActionList `json:"action_list"`
}

// MainTitle 一级标题
type MainTitle struct {
	// 一级标题，建议不超过36个字，文本通知型卡片本字段非必填，但不可本字段和sub_title_text都不填，（支持id转译）
	Title string `json:"title"`
	// 标题辅助信息，建议不超过160个字，（支持id转译）
	Desc string `json:"desc"`
}

// QuoteArea 引用文献样式
type QuoteArea struct {
	// 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	Type int `json:"type"`
	// 点击跳转的url，quote_area.type是1时必填
	URL string `json:"url"`
	// 引用文献样式的标题
	Title string `json:"title"`
	// 引用文献样式的引用文案
	QuoteText string `json:"quote_text"`
	// 小程序appid，必须是与当前应用关联的小程序，appid和pagepath必须同时填写，填写后会忽略url字段
	AppID string `json:"appid"`
	// 点击消息卡片后的小程序页面，最长128字节，仅限本小程序内的页面。appid和pagepath必须同时填写，填写后会忽略url字段
	PagePath string `json:"pagepath"`
}

// EmphasisContent 关键数据样式
type EmphasisContent struct {
	// 关键数据样式的数据内容，建议不超过14个字
	Title string `json:"title"`
	// 关键数据样式的数据描述内容，建议不超过22个字
	Desc string `json:"desc"`
}

// HorizontalContentList 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
type HorizontalContentList struct {
	// 二级标题，建议不超过5个字
	KeyName string `json:"keyname"`
	// 二级文本，如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过30个字，（支持id转译）
	Value string `json:"value"`
	// 链接类型，0或不填代表不是链接，1 代表跳转url，2 代表下载附件，3 代表点击跳转成员详情
	Type int `json:"type,omitempty"`
	// 链接跳转的url，horizontal_content_list.type是1时必填
	URL string `json:"url,omitempty"`
	// 附件的media_id，horizontal_content_list.type是2时必填
	MediaID string `json:"media_id,omitempty"`
	// 成员详情的userid，horizontal_content_list.type是3时必填
	Userid string `json:"userid,omitempty"`
}

// JumpList 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
type JumpList struct {
	// 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
	Type int `json:"type"`
	// 跳转链接样式的文案内容，建议不超过18个字
	Title string `json:"title"`
	// 跳转链接的url，jump_list.type是1时必填
	URL string `json:"url,omitempty"`
	// 跳转链接的小程序的appid，必须是与当前应用关联的小程序，jump_list.type是2时必填
	Appid string `json:"appid,omitempty"`
	// 跳转链接的小程序的pagepath，jump_list.type是2时选填
	PagePath string `json:"pagepath,omitempty"`
}

// CardAction 整体卡片的点击跳转事件，text_notice必填本字段
type CardAction struct {
	// 跳转事件类型，1 代表跳转url，2 代表打开小程序。text_notice卡片模版中该字段取值范围为[1,2]
	Type int `json:"type"`
	// 跳转事件的url，card_action.type是1时必填
	URL string `json:"url"`
	// 跳转事件的小程序的appid，必须是与当前应用关联的小程序，card_action.type是2时必填
	Appid string `json:"appid"`
	// 跳转事件的小程序的pagepath，card_action.type是2时选填
	Pagepath string `json:"pagepath"`
}

// ImageTextArea 左图右文样式，news_notice类型的卡片，card_image和image_text_area两者必填一个字段，不可都不填
type ImageTextArea struct {
	// 左图右文样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	Type int `json:"type"`
	// 点击跳转的url，image_text_area.type是1时必填
	URL string `json:"url"`
	// 点击跳转的小程序的appid，必须是与当前应用关联的小程序，image_text_area.type是2时必填
	AppID string `json:"appid,omitempty"`
	// 点击跳转的小程序的pagepath，image_text_area.type是2时选填
	PagePath string `json:"pagepath,omitempty"`
	// 左图右文样式的标题
	Title string `json:"title"`
	// 左图右文样式的描述
	Desc string `json:"desc"`
	// 左图右文样式的图片url
	ImageURL string `json:"image_url"`
}

// CardImage 图片样式，news_notice类型的卡片，card_image和image_text_area两者必填一个字段，不可都不填
type CardImage struct {
	// 图片的url
	URL string `json:"url"`
	// 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
	AspectRatio float32 `json:"aspect_ratio"`
}

// ButtonSelection 按钮交互型
type ButtonSelection struct {
	// 下拉式的选择器的key，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
	QuestionKey string `json:"question_key"`
	// 下拉式的选择器的key，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
	Title string `json:"title"`
	// 选项列表，下拉选项不超过 10 个，最少1个
	OptionList []struct {
		// 下拉式的选择器选项的id，用户提交后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
		ID string `json:"id"`
		// 下拉式的选择器选项的文案，建议不超过16个字
		Text string `json:"text"`
	} `json:"option_list"`
	// 默认选定的id，不填或错填默认第一个
	SelectedID string `json:"selected_id"`
}

type Button struct {
	// 按钮点击事件类型，0 或不填代表回调点击事件，1 代表跳转url
	Type int `json:"type,omitempty"`
	// 按钮文案，建议不超过10个字
	Text string `json:"text"`
	// 按钮样式，目前可填1~4，不填或错填默认1
	Style int `json:"style,omitempty"`
	// 按钮key值，用户点击后，会产生回调事件将本参数作为EventKey返回，回调事件会带上该key值，最长支持1024字节，不可重复，button_list.type是0时必填
	Key string `json:"key,omitempty"`
	// 跳转事件的url，button_list.type是1时必填
	URL string `json:"url,omitempty"`
}

// CheckBox 选择题样式
type CheckBox struct {
	// 选择题key值，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
	QuestionKey string `json:"question_key"`
	// 选项list，选项个数不超过 20 个，最少1个
	OptionList []struct {
		// 选项id，用户提交选项后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
		ID string `json:"id"`
		// 选项文案描述，建议不超过17个字
		Text string `json:"text"`
		// 该选项是否要默认选中
		IsChecked bool `json:"is_checked"`
	} `json:"option_list" validate:"required,min=1,max=20"`
	// 选择题模式，单选：0，多选：1，不填默认0
	Mode int `json:"mode" validate:"omitempty,oneof=0 1"`
}

// SubmitButton 提交按钮样式
type SubmitButton struct {
	// 按钮文案，建议不超过10个字，不填默认为提交
	Text string `json:"text"`
	// 提交按钮的key，会产生回调事件将本参数作为EventKey返回，最长支持1024字节
	Key string `json:"key"`
}

// SelectList 下拉式的选择器列表，multiple_interaction类型的卡片该字段不可为空，一个消息最多支持 3 个选择器
type SelectList struct {
	// 下拉式的选择器题目的key，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节，不可重复
	QuestionKey string `json:"question_key"`
	// 下拉式的选择器上面的title
	Title string `json:"title,omitempty"`
	// 默认选定的id，不填或错填默认第一个
	SelectedID string       `json:"selected_id,omitempty"`
	OptionList []OptionList `json:"option_list"`
}

// 项列表，下拉选项不超过 10 个，最少1个
type OptionList struct {
	// 下拉式的选择器选项的id，用户提交选项后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
	ID string `json:"id"`
	// 下拉式的选择器选项的文案，建议不超过16个字
	Text string `json:"text"`
}

type respUpdateContactWayExternalContact struct {
	respCommon
}

type reqDelContactWayExternalContact struct {
	ConfigID string `json:"config_id"`
}

var _ bodyer = reqDelContactWayExternalContact{}

func (x reqDelContactWayExternalContact) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type respDelContactWayExternalContact struct {
	respCommon
}

type reqCloseTempChatExternalContact struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
}

var _ bodyer = reqCloseTempChatExternalContact{}

func (x reqCloseTempChatExternalContact) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type respCloseTempChatExternalContact struct {
	respCommon
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

// TemplateCardType 模板卡片的类型
type TemplateCardType string

const (
	CardTypeTextNotice          TemplateCardType = "text_notice"          // 文本通知型
	CardTypeNewsNotice          TemplateCardType = "news_notice"          // 图文展示型
	CardTypeButtonInteraction   TemplateCardType = "button_interaction"   // 按钮交互型
	CardTypeVoteInteraction     TemplateCardType = "vote_interaction"     // 投票选择型
	CardTypeMultipleInteraction TemplateCardType = "multiple_interaction" // 多项选择型
)

type TemplateCard struct {
	CardType   TemplateCardType `json:"card_type"`
	Source     Source           `json:"source"`
	ActionMenu *ActionMenu      `json:"action_menu,omitempty" validate:"required_with=TaskID"`
	TaskID     string           `json:"task_id,omitempty" validate:"required_with=ActionMenu"`
	MainTitle  *MainTitle       `json:"main_title"`
	QuoteArea  *QuoteArea       `json:"quote_area,omitempty"`
	// 文本通知型
	EmphasisContent *EmphasisContent `json:"emphasis_content,omitempty"`
	SubTitleText    string           `json:"sub_title_text,omitempty"`
	// 图文展示型
	ImageTextArea         *ImageTextArea          `json:"image_text_area,omitempty"`
	CardImage             *CardImage              `json:"card_image,omitempty"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list"`
	JumpList              []JumpList              `json:"jump_list"`
	CardAction            *CardAction             `json:"card_action,omitempty"`
	// 按钮交互型
	ButtonSelection *ButtonSelection `json:"button_selection,omitempty"`
	ButtonList      []Button         `json:"button_list,omitempty" validate:"omitempty,max=6"`
	// 投票选择型
	CheckBox     *CheckBox     `json:"checkbox,omitempty"`
	SelectList   []SelectList  `json:"select_list,omitempty" validate:"max=3"`
	SubmitButton *SubmitButton `json:"submit_button,omitempty"`
}

type TemplateCardUpdateMessage struct {
	UserIds      []string `json:"userids" validate:"omitempty,max=100"`
	PartyIds     []int64  `json:"partyids" validate:"omitempty,max=100"`
	TagIds       []int32  `json:"tagids" validate:"omitempty,max=100"`
	AtAll        int      `json:"atall,omitempty"`
	ResponseCode string   `json:"response_code"`
	Button       struct {
		ReplaceName string `json:"replace_name"`
	} `json:"button" validate:"required_without=TemplateCard"`
	TemplateCard TemplateCard `json:"template_card" validate:"required_without=Button"`
	ReplaceText  string       `json:"replace_text,omitempty"`
}
