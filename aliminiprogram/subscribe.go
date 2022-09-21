package aliminiprogram

type AuditPassResp struct {
	MiniAppId string `json:"mini_app_id"`
}
type AuditRejectResp struct {
	MiniAppId          string             `json:"mini_app_id"`
	BundleId           string             `json:"bundle_id"`
	MiniAppVersion     string             `json:"mini_app_version"` // 小程序提审版本号
	AuditReason        string             `json:"audit_reason"`
	BaseAudit          string             `json:"base_audit"`
	PromoteAudit       string             `json:"promote_audit"`
	BaseAuditReason    MiniAppAuditReason `json:"base_audit_reason"`
	PromoteAuditReason MiniAppAuditReason `json:"promote_audit_reason"`
}

type MiniAppAuditReason struct {
	AuditImages []string                 `json:"audit_images"`
	Memos       []MiniAppAuditReasonMemo `json:"memos"`
}

type MiniAppAuditReasonMemo struct {
	Memo          string   `json:"memo"`
	MemoImageList []string `json:"memo_image_list"`
}

// SportVenueAuditNotify  场馆审核消息属性
type SportVenueAuditNotify struct {
	VenueId       string   `json:"venue_id"`         // 支付宝对应的场馆id
	OutVenueId    string   `json:"out_venue_id"`     // 服务商对应的场馆id
	OutSubVenueId []string `json:"out_sub_venue_id"` // 服务商场馆id列表
	/*
		安全审核中：infosec-audit
		安全审核不通过：infosec-unpass
		云验收中： cloud-audit
		云验收不通过： cloud-unpass
		上架： online
		下架： offline
		人工下架： manual-offline
	*/
	VenueStatus  string   `json:"venue_status"`  // 场馆当前状态
	Desc         string   `json:"desc"`          // 场馆描述
	SubVenueId   []string `json:"sub_venue_id"`  // 子场馆id列表
	UtcTimeStamp int64    `json:"utc_timestamp"` //
}

type VenueAvailable struct {
	Response struct {
		Code   string `json:"code"`
		Msg    string `json:"msg"`
		Enable bool   `json:"enable"`
	} `json:"response"`
	AppCertSn string `json:"app_cert_sn"`
}

type VenueAvailableErr struct {
	Response  AliMiniResponse `json:"response"`
	AppCertSn string          `json:"app_cert_sn"`
}

type AliMiniResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

// IsvCreateMerchantConfirm isv代商户创建小程序返回参数
type IsvCreateMerchantConfirm struct {
	OrderNO    string `json:"order_no"`     // 服务商代商家创建小程序，由支付宝开放平台返回的订单号
	OutOrderNO string `json:"out_order_no"` // 外部订单号，开发者2088帐号+外部订单号维度来控制请求幂等
	MinAppId   string `json:"min_app_id"`   // 小程序应用id，商家确认同意创建非空，商家拒绝或超时返回空
	Status     string `json:"status"`       // 状态 PROCESS处理中，TIMEOUT超时，AGREED同意， REJECTED拒绝
	Pid        string `json:"pid"`          // 商户PID
}
