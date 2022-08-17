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
	VenueId       string `json:"venue_id"`         // 支付宝对应的场馆id
	OutVenueId    string `json:"out_venue_id"`     // 服务商对应的场馆id
	OutSubVenueId string `json:"out_sub_venue_id"` // 服务商场馆id列表
	/*
		安全审核中：infosec-audit
		安全审核不通过：infosec-unpass
		云验收中： cloud-audit
		云验收不通过： cloud-unpass
		上架： online
		下架： offline
		人工下架： manual-offline
	*/
	VenueStatus string   `json:"venue_status"` // 场馆当前状态
	Desc        string   `json:"desc"`         // 场馆描述
	SubVenueId  []string `json:"sub_venue_id"` // 子场馆id列表
}
