// Code generated by sdkcodegen; DO NOT EDIT.

package go_qywechat

// ExternalContact 外部联系人
type ExternalContact struct {
	// ExternalUserid 外部联系人的userid
	ExternalUserid string `json:"external_userid"`
	// Name 外部联系人的名称，如果外部联系人为微信用户，则返回外部联系人的名称为其微信昵称；如果外部联系人为企业微信用户，则会按照以下优先级顺序返回：此外部联系人或管理员设置的昵称、认证的实名和账号名称。
	Name string `json:"name"`
	// Position 外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，仅当联系人类型是企业微信用户时有此字段
	Position string `json:"position"`
	// Avatar 外部联系人头像，第三方不可获取
	Avatar string `json:"avatar"`
	// CorpName 外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
	CorpName string `json:"corp_name"`
	// Type 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Type ExternalUserType `json:"type"`
	// Gender 外部联系人性别 0-未知 1-男性 2-女性
	Gender UserGender `json:"gender"`
	// Unionid 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。查看绑定方法 关于返回的unionid，如果是第三方应用调用该接口，则返回的unionid是该第三方服务商所关联的微信开放者帐号下的unionid。也就是说，同一个企业客户，企业自己调用，与第三方服务商调用，所返回的unionid不同；不同的服务商调用，所返回的unionid也不同。
	Unionid string `json:"unionid"`
	// ExternalProfile 成员对外信息
	ExternalProfile ExternalProfile `json:"external_profile"`
}

// ExternalProfile 成员对外信息
type ExternalProfile struct {
	// ExternalCorpName 企业简称
	ExternalCorpName string `json:"external_corp_name"`
	// ExternalAttr 属性列表，目前支持文本、网页、小程序三种类型
	ExternalAttr []ExternalAttr `json:"external_attr"`
}

// ExternalAttr 属性列表，目前支持文本、网页、小程序三种类型
type ExternalAttr struct {
	// Type 属性类型: 0-文本 1-网页 2-小程序
	Type int `json:"type"`
	// Name 属性名称： 需要先确保在管理端有创建该属性，否则会忽略
	Name string `json:"name"`
	// Text 文本类型的属性 ，type为0时必填
	Text ExternalAttrText `json:"text"`
	// Web 网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空 ，type为1时必填
	Web ExternalAttrWeb `json:"web"`
	// Miniprogram 小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空 ，type为2时必填
	Miniprogram ExternalAttrMiniprogram `json:"miniprogram"`
}

// ExternalAttrText 文本类型的属性
type ExternalAttrText struct {
	// Value 文本属性内容,长度限制12个UTF8字符
	Value string `json:"value"`
}

// ExternalAttrWeb 网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空 ，type为1时必填
type ExternalAttrWeb struct {
	// Url 网页的url,必须包含http或者https头
	Url string `json:"url"`
	// Title 网页的展示标题,长度限制12个UTF8字符
	Title string `json:"title"`
}

// ExternalAttrMiniprogram 小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空 ，type为2时必填
type ExternalAttrMiniprogram struct {
	// Appid 小程序appid，必须是有在本企业安装授权的小程序，否则会被忽略
	Appid string `json:"appid"`
	// Pagepath 小程序的页面路径
	Pagepath string `json:"pagepath"`
	// Title 企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。
	Title string `json:"title"`
}

// ExternalUserType 外部联系人的类型
//
// 1表示该外部联系人是微信用户
// 2表示该外部联系人是企业微信用户
type ExternalUserType int

const (
	// ExternalUserTypeWeChat 微信用户
	ExternalUserTypeWeChat ExternalUserType = 1
	// ExternalUserTypeWorkWeChat 企业微信用户
	ExternalUserTypeWorkWeChat ExternalUserType = 2
)

// FollowUser 添加了外部联系人的企业成员
type FollowUser struct {
	//  添加了外部联系人的企业成员
	FollowUserInfo
	// Tags 该成员添加此外部联系人所打标签
	Tags []FollowUserTag `json:"tags"`
}

// FollowInfo 企业成员客户跟进信息，可以参考获取客户详情，但标签信息只会返回企业标签的tag_id，个人标签将不再返回
type FollowInfo struct {
	//  添加了外部联系人的企业成员
	FollowUserInfo
	// TagID 该成员添加此外部联系人所打标签
	TagID []string `json:"tag_id"`
}

// FollowUserInfo 添加了外部联系人的企业成员
type FollowUserInfo struct {
	// UserID 外部联系人的userid
	UserID string `json:"userid"`
	// Remark 该成员对此外部联系人的备注
	Remark string `json:"remark"`
	// Description 该成员对此外部联系人的描述
	Description string `json:"description"`
	// Createtime 该成员添加此外部联系人的时间
	Createtime int `json:"createtime"`
	// RemarkCorpName 该成员对此客户备注的企业名称
	RemarkCorpName string `json:"remark_corp_name"`
	// RemarkMobiles 该成员对此客户备注的手机号码，第三方不可获取
	RemarkMobiles []string `json:"remark_mobiles"`
	// AddWay 该成员添加此客户的来源
	AddWay FollowUserAddWay `json:"add_way"`
	// OperUserID 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
	OperUserID string `json:"oper_userid"`
	// State 企业自定义的state参数，用于区分客户具体是通过哪个「联系我」添加，由企业通过创建「联系我」方式指定
	State string `json:"state"`
}

// FollowUserTag 该成员添加此外部联系人所打标签
type FollowUserTag struct {
	// GroupName 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	GroupName string `json:"group_name"`
	// TagName 该成员添加此外部联系人所打标签名称
	TagName string `json:"tag_name"`
	// Type 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
	Type FollowUserTagType `json:"type"`
}

// FollowUserTagType 该成员添加此外部联系人所打标签类型
//
// 1-企业设置
// 2-用户自定义
type FollowUserTagType int

const (
	// 企业设置
	FollowUserTagTypeWork FollowUserTagType = 1
	// 用户自定义
	FollowUserTagTypeUser FollowUserTagType = 2
)

// FollowUserAddWay 该成员添加此客户的来源
//
// 具体含义详见[来源定义](https://work.weixin.qq.com/api/doc/90000/90135/92114#13878/%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89)
type FollowUserAddWay int

const (
	// 未知来源
	FollowUserAddWayUnknown FollowUserAddWay = 0
	// 扫描二维码
	FollowUserAddWayQRCode FollowUserAddWay = 1
	// 搜索手机号
	FollowUserAddWayMobile FollowUserAddWay = 2
	// 名片分享
	FollowUserAddWayCard FollowUserAddWay = 3
	// 群聊
	FollowUserAddWayGroupChat FollowUserAddWay = 4
	// 手机通讯录
	FollowUserAddWayAddressBook FollowUserAddWay = 5
	// 微信联系人
	FollowUserAddWayWeChatContact FollowUserAddWay = 6
	// 来自微信的添加好友申请
	FollowUserAddWayWeChatFriendApply FollowUserAddWay = 7
	// 安装第三方应用时自动添加的客服人员
	FollowUserAddWayThirdParty FollowUserAddWay = 8
	// 搜索邮箱
	FollowUserAddWayEmail FollowUserAddWay = 9
	// 内部成员共享
	FollowUserAddWayInternalShare FollowUserAddWay = 201
	// 管理员/负责人分配
	FollowUserAddWayAdmin FollowUserAddWay = 202
)

// ExternalContactRemark 客户备注信息
type ExternalContactRemark struct {
	// Userid 企业成员的userid
	Userid string `json:"userid"`
	// ExternalUserid 外部联系人userid
	ExternalUserid string `json:"external_userid"`
	// Remark 此用户对外部联系人的备注，最多20个字符，remark，description，remark_company，remark_mobiles和remark_pic_mediaid不可同时为空。
	Remark string `json:"remark"`
	// Description 此用户对外部联系人的描述，最多150个字符
	Description string `json:"description"`
	// RemarkCompany 此用户对外部联系人备注的所属公司名称，最多20个字符，remark_company只在此外部联系人为微信用户时有效。
	RemarkCompany string `json:"remark_company"`
	// RemarkMobiles 此用户对外部联系人备注的手机号，如果填写了remark_mobiles，将会覆盖旧的备注手机号。如果要清除所有备注手机号,请在remark_mobiles填写一个空字符串(“”)。
	RemarkMobiles []string `json:"remark_mobiles"`
	// RemarkPicMediaid 备注图片的mediaid，remark_pic_mediaid可以通过素材管理接口获得。
	RemarkPicMediaid string `json:"remark_pic_mediaid"`
}

// ExternalContactCorpTag 企业客户标签
type ExternalContactCorpTag struct {
	// ID 标签id
	ID string `json:"id"`
	// Name 标签名称
	Name string `json:"name"`
	// CreateTime 标签创建时间
	CreateTime int `json:"create_time"`
	// Order 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Order uint32 `json:"order"`
	// Deleted 标签是否已经被删除，只在指定tag_id进行查询时返回
	Deleted bool `json:"deleted"`
}

// ExternalContactCorpTagGroup 企业客户标签
type ExternalContactCorpTagGroup struct {
	// GroupID 标签组id
	GroupID string `json:"group_id"`
	// GroupName 标签组名称
	GroupName string `json:"group_name"`
	// CreateTime 标签组创建时间
	CreateTime int `json:"create_time"`
	// Order 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Order uint32 `json:"order"`
	// Deleted 标签组是否已经被删除，只在指定tag_id进行查询时返回
	Deleted bool `json:"deleted"`
	// Tag 标签组内的标签列表
	Tag []ExternalContactCorpTag `json:"tag"`
}

// ExternalContactMarkTag 企业标记客户标签
type ExternalContactMarkTag struct {
	// UserID 添加外部联系人的userid
	UserID string `json:"userid"`
	// ExternalUserID 外部联系人userid
	ExternalUserID string `json:"external_userid"`
	// AddTag 要标记的标签列表
	AddTag []string `json:"add_tag"`
	// RemoveTag 要移除的标签列表
	RemoveTag []string `json:"remove_tag"`
}

// ExternalContactUnassignedList 离职成员的客户列表
type ExternalContactUnassignedList struct {
	// Info 离职成员的客户
	Info []ExternalContactUnassigned `json:"info"`
	// IsLast 是否是最后一条记录
	IsLast bool `json:"is_last"`
	// NextCursor 分页查询游标,已经查完则返回空("")
	NextCursor string `json:"next_cursor"`
}

// ExternalContactTransferStatus 客户接替结果状态
type ExternalContactTransferStatus uint8

const (
	// ExternalContactTransferStatusSuccess 1-接替完毕
	ExternalContactTransferStatusSuccess ExternalContactTransferStatus = 1
	// ExternalContactTransferStatusWait 2-等待接替
	ExternalContactTransferStatusWait ExternalContactTransferStatus = 2
	// ExternalContactTransferStatusRefused 3-客户拒绝
	ExternalContactTransferStatusRefused ExternalContactTransferStatus = 3
	// ExternalContactTransferStatusExhausted 4-接替成员客户达到上限
	ExternalContactTransferStatusExhausted ExternalContactTransferStatus = 4
	// ExternalContactTransferStatusNoData 5-无接替记录
	ExternalContactTransferStatusNoData ExternalContactTransferStatus = 5
)

// ExternalContactGroupChatTransferFailed 离职成员的群再分配失败
type ExternalContactGroupChatTransferFailed struct {
	// ChatID 没能成功继承的群ID
	ChatID string `json:"chat_id"`
	// ErrCode 没能成功继承的群，错误码
	ErrCode int `json:"errcode"`
	// ErrMsg 没能成功继承的群，错误描述
	ErrMsg string `json:"errmsg"`
}

// ExternalContactFollowUserList 配置了客户联系功能的成员列表
type ExternalContactFollowUserList struct {
	// FollowUser 配置了客户联系功能的成员userid列表
	FollowUser []string `json:"follow_user"`
}

// ExternalContactWay 配置客户联系「联系我」方式
type ExternalContactWay struct {
	// Type 联系方式类型,1-单人, 2-多人
	Type int `json:"type"`
	// Scene 场景，1-在小程序中联系，2-通过二维码联系
	Scene int `json:"scene"`
	// Style 在小程序中联系时使用的控件样式，详见附表
	Style int `json:"style"`
	// Remark 联系方式的备注信息，用于助记，不超过30个字符
	Remark string `json:"remark"`
	// SkipVerify 外部客户添加时是否无需验证，默认为true
	SkipVerify bool `json:"skip_verify"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符 https://developer.work.weixin.qq.com/document/path/92114
	State string `json:"state"`
	// User 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
	User []string `json:"user"`
	// Party 使用该联系方式的部门id列表，只在type为2时有效
	Party []int `json:"party"`
	// IsTemp 是否临时会话模式，true表示使用临时会话模式，默认为false
	IsTemp bool `json:"is_temp"`
	// ExpiresIn 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天，最多为14天
	ExpiresIn int `json:"expires_in"`
	// ChatExpiresIn 临时会话有效期，以秒为单位。该参数仅在is_temp为true时有效，默认为添加好友后24小时，最多为14天
	ChatExpiresIn int `json:"chat_expires_in"`
	// UnionID 可进行临时会话的客户UnionID，该参数仅在is_temp为true时有效，如不指定则不进行限制
	UnionID string `json:"unionid"`
	// Conclusions 结束语，会话结束时自动发送给客户，可参考“结束语定义”，仅在is_temp为true时有效,https://developer.work.weixin.qq.com/document/path/92572#%E7%BB%93%E6%9D%9F%E8%AF%AD%E5%AE%9A%E4%B9%89
	Conclusions Conclusions `json:"conclusions"`
}

// Conclusions 结束语，会话结束时自动发送给客户
type Conclusions struct {
	// Text 文本消息
	Text Text `json:"text"`
	// Image 图片
	Image Image `json:"image"`
	// Link 链接
	Link Link `json:"link"`
	// MiniProgram 小程序
	MiniProgram MiniProgram `json:"miniprogram"`
}

// Text 结束语，会话结束时自动发送给客户
type Text struct {
	// Content 消息文本内容,最长为4000字节
	Content string `json:"content"`
}

// Image 结束语，会话结束时自动发送给客户
type Image struct {
	// MediaID 图片的media_id
	MediaID string `json:"media_id"`
	// PicURL 图片的url
	PicURL string `json:"pic_url"`
}

// Link 结束语，会话结束时自动发送给客户
type Link struct {
	// Title 图文消息标题，最长为128字节
	Title string `json:"title"`
	// Picurl 图文消息封面的url
	Picurl string `json:"picurl"`
	// Desc 图文消息的描述，最长为512字节
	Desc string `json:"desc"`
	// URL 图文消息的链接
	URL string `json:"url"`
}

// MiniProgram 结束语，会话结束时自动发送给客户
type MiniProgram struct {
	// Title 小程序消息标题，最长为64字节
	Title string `json:"title"`
	// PicMediaID 小程序消息封面的mediaid，封面图建议尺寸为520*416
	PicMediaID string `json:"pic_media_id"`
	// AppID 小程序appid，必须是关联到企业的小程序应用
	AppID string `json:"appid"`
	// Page 小程序page路径
	Page string `json:"page"`
}

// reqListContactWayExternalContact 获取企业已配置的「联系我」列表请求参数
type reqListContactWayExternalContact struct {
	// StartTime 「联系我」创建起始时间戳, 默认为90天前
	StartTime int `json:"start_time"`
	// EndTime 「联系我」创建结束时间戳, 默认为当前时间
	EndTime int `json:"end_time"`
	// Cursor 分页查询使用的游标，为上次请求返回的 next_cursor
	Cursor string `json:"cursor"`
	// Limit 每次查询的分页大小，默认为100条，最多支持1000条
	Limit int `json:"limit"`
}

// reqUpdateContactWayExternalContact 更新企业已配置的「联系我」方式请求参数
type reqUpdateContactWayExternalContact struct {
	// ConfigID 企业联系方式的配置id
	ConfigID string `json:"config_id"`
	// Remark 联系方式的备注信息，不超过30个字符，将覆盖之前的备注
	Remark string `json:"remark"`
	// SkipVerify 外部客户添加时是否无需验证
	SkipVerify bool `json:"skip_verify"`
	// Style 样式，只针对“在小程序中联系”的配置生效
	Style int `json:"style"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符 https://developer.work.weixin.qq.com/document/path/92114
	State string `json:"state"`
	// User 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
	User []string `json:"user"`
	// Party 使用该联系方式的部门id列表，只在type为2时有效
	Party []int `json:"party"`
	// ExpiresIn 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天，最多为14天
	ExpiresIn int `json:"expires_in"`
	// ChatExpiresIn 临时会话有效期，以秒为单位。该参数仅在is_temp为true时有效，默认为添加好友后24小时，最多为14天
	ChatExpiresIn int `json:"chat_expires_in"`
	// UnionID 可进行临时会话的客户UnionID，该参数仅在is_temp为true时有效，如不指定则不进行限制
	UnionID string `json:"unionid"`
	// Conclusions 结束语，会话结束时自动发送给客户，可参考“结束语定义”，仅在is_temp为true时有效,https://developer.work.weixin.qq.com/document/path/92572#%E7%BB%93%E6%9D%9F%E8%AF%AD%E5%AE%9A%E4%B9%89
	Conclusions Conclusions `json:"conclusions"`
}
