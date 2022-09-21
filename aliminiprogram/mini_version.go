package aliminiprogram

import (
	"context"
	"encoding/json"
	"io"
)

type QueryVersionListBiz struct {
	BundleId      string `json:"bundle_id"`
	VersionStatus string `json:"version_status"`
}

// QueryVersionListResp 查询小程序列表
type QueryVersionListResp struct {
	AlipayOpenMiniVersionListQueryResponse struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		AppVersionInfos []struct {
			AppVersion    string `json:"app_version"`
			CreateTime    string `json:"create_time"`
			BundleId      string `json:"bundle_id"`
			VersionStatus string `json:"version_status"`
		} `json:"app_version_infos"`
		AppVersions []string `json:"app_versions"`
	} `json:"alipay_open_mini_version_list_query_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

func (s *MiniService) QueryVersionList(ctx context.Context, biz QueryVersionListBiz, opts ...ValueOptions) (*QueryVersionListResp, error) {
	apiMethod := "alipay.open.mini.version.list.query"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(QueryVersionListResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteVersionBiz 小程序删除版本
type DeleteVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号
	BundleID   string `json:"bundle_id"`   //小程序投放的端参数，例如投放到支付宝钱包是支付宝端。该参数可选，默认支付宝端，目前仅支持支付宝端，枚举列举：com.alipay.alipaywallet:支付宝端
}

// DeleteVersion 小程序删除版本
func (s *MiniService) DeleteVersion(ctx context.Context, biz *DeleteVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.delete"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

type ApplyVersionAuditResp struct {
	AlipayOpenMiniVersionAuditApplyResponse struct {
		Code        string `json:"code"`
		Msg         string `json:"msg"`
		SpeedUpMemo string `json:"speed_up_memo"`
		SpeedUp     string `json:"speed_up"`
	} `json:"alipay_open_mini_version_audit_apply_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// ApplyVersionAuditBiz 小程序提交审核
type ApplyVersionAuditBiz struct {
	FirstSpecialLicensePic  string        `json:"first_special_license_pic,omitempty"`
	SecondSpecialLicensePic string        `json:"second_special_license_pic,omitempty"`
	ThirdSpecialLicensePic  string        `json:"third_special_license_pic,omitempty"`
	FirstLicensePic         string        `json:"first_license_pic,omitempty"`
	SecondLicensePic        string        `json:"second_license_pic,omitempty"`
	ThirdLicensePic         string        `json:"third_license_pic,omitempty"`
	FourthLicensePic        string        `json:"fourth_license_pic,omitempty"`
	FifthLicensePic         string        `json:"fifth_license_pic,omitempty"`
	OutDoorPic              string        `json:"out_door_pic,omitempty"`
	FirstScreenShot         string        `json:"first_screen_shot,omitempty"`
	SecondScreenShot        string        `json:"second_screen_shot,omitempty"`
	ThirdScreenShot         string        `json:"third_screen_shot,omitempty"`
	FourthScreenShot        string        `json:"fourth_screen_shot,omitempty"`
	FifthScreenShot         string        `json:"fifth_screen_shot,omitempty"`
	TestFileName            []byte        `json:"test_file_name,omitempty"`
	LicenseValidDate        string        `json:"license_valid_date,omitempty"`
	AppVersion              string        `json:"app_version"`
	LicenseName             string        `json:"license_name,omitempty"`
	AutoOnline              string        `json:"auto_online"`
	SpeedUp                 string        `json:"speed_up"`
	AuditRule               string        `json:"audit_rule,omitempty"` // 审核类型
	AppName                 string        `json:"app_name,omitempty"`
	AppEnglishName          string        `json:"app_english_name,omitempty"`
	AppSlogan               string        `json:"app_slogan,omitempty"`
	AppLogo                 string        `json:"app_logo,omitempty"`
	AppCategoryIDs          string        `json:"app_category_ids,omitempty"`
	AppDesc                 string        `json:"app_desc,omitempty"`
	ServicePhone            string        `json:"service_phone,omitempty"`
	ServiceEmail            string        `json:"service_email,omitempty"`
	VersionDesc             string        `json:"version_desc"`
	Memo                    string        `json:"memo,omitempty"`
	RegionType              string        `json:"region_type"`
	ServiceRegionInfo       []*RegionInfo `json:"service_region_info,omitempty"`
	LicenseNo               string        `json:"license_no,omitempty"`
	MiniCategoryIDs         string        `json:"mini_category_ids,omitempty"`
	TestAccount             string        `json:"test_accout,omitempty"` // 官方拼写错误
	TestPassword            string        `json:"test_password,omitempty"`
	BundleID                string        `json:"bundle_id,omitempty"`
}

func (a ApplyVersionAuditBiz) Params() map[string]string {
	params := make(map[string]string)
	if a.LicenseName != "" {
		params["license_name"] = a.LicenseName
	}
	if a.LicenseValidDate != "" {
		params["license_valid_date"] = a.LicenseValidDate
	}
	if a.AppVersion != "" {
		params["app_version"] = a.AppVersion
	}
	if a.AppName != "" {
		params["app_name"] = a.AppName
	}
	if a.AppEnglishName != "" {
		params["app_english_name"] = a.AppEnglishName
	}
	if a.AppSlogan != "" {
		params["app_slogan"] = a.AppSlogan
	}
	if a.AppCategoryIDs != "" {
		params["app_category_ids"] = a.AppCategoryIDs
	}
	if a.AppDesc != "" {
		params["app_desc"] = a.AppDesc
	}
	if a.ServicePhone != "" {
		params["service_phone"] = a.ServicePhone
	}
	if a.ServiceEmail != "" {
		params["service_email"] = a.ServiceEmail
	}
	if a.VersionDesc != "" {
		params["version_desc"] = a.VersionDesc
	}
	if a.Memo != "" {
		params["memo"] = a.Memo
	}
	if a.RegionType != "" {
		params["region_type"] = a.RegionType
	}
	if a.ServiceRegionInfo != nil {
		serviceRegionInfo, _ := json.Marshal(a.ServiceRegionInfo)
		params["service_region_info"] = string(serviceRegionInfo)
	}
	if a.LicenseNo != "" {
		params["license_no"] = a.LicenseNo
	}
	if a.MiniCategoryIDs != "" {
		params["mini_category_ids"] = a.MiniCategoryIDs
	}
	if a.TestAccount != "" {
		params["test_accout"] = a.TestAccount
	}
	if a.TestPassword != "" {
		params["test_password"] = a.TestPassword
	}
	if a.BundleID != "" {
		params["bundle_id"] = a.BundleID
	}
	if a.FirstLicensePic != "" {
		params["first_license_pic"] = a.FirstLicensePic
	}
	if a.SecondLicensePic != "" {
		params["second_license_pic"] = a.SecondLicensePic
	}
	if a.ThirdLicensePic != "" {
		params["third_license_pic"] = a.ThirdLicensePic
	}
	if a.FourthLicensePic != "" {
		params["fourth_license_pic"] = a.FourthLicensePic
	}
	if a.FifthLicensePic != "" {
		params["fifth_license_pic"] = a.FifthLicensePic
	}
	if a.OutDoorPic != "" {
		params["out_door_pic"] = a.OutDoorPic
	}
	if a.FirstScreenShot != "" {
		params["first_screen_shot"] = a.FirstScreenShot
	}
	if a.SecondScreenShot != "" {
		params["second_screen_shot"] = a.SecondScreenShot
	}
	if a.ThirdScreenShot != "" {
		params["third_screen_shot"] = a.ThirdScreenShot
	}
	if a.FourthScreenShot != "" {
		params["fourth_screen_shot"] = a.FourthScreenShot
	}
	if a.FifthScreenShot != "" {
		params["fifth_screen_shot"] = a.FifthScreenShot
	}
	if a.FirstSpecialLicensePic != "" {
		params["first_special_license_pic"] = a.FirstSpecialLicensePic
	}
	if a.SecondSpecialLicensePic != "" {
		params["second_special_license_pic"] = a.SecondSpecialLicensePic
	}
	if a.ThirdSpecialLicensePic != "" {
		params["third_special_license_pic"] = a.ThirdSpecialLicensePic
	}
	return params
}

func (a ApplyVersionAuditBiz) MultipartParams() map[string]io.Reader {
	params := make(map[string]io.Reader)
	return params
}

// RegionInfo 省市区信息，当区域类型为LOCATION时，不能为空
type RegionInfo struct {
	ProvinceCode string `json:"province_code"`
	ProvinceName string `json:"province_name"`
	CityCode     string `json:"city_code"`
	CityName     string `json:"city_name"`
	AreaCode     string `json:"area_code"`
	AreaName     string `json:"area_name"`
}

// ApplyVersionAudit 小程序提交审核
func (s *MiniService) ApplyVersionAudit(ctx context.Context, biz *ApplyVersionAuditBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.audit.apply"
	params := make(map[string]interface{})
	if biz.LicenseName != "" {
		params["license_name"] = biz.LicenseName
	}
	if biz.AppLogo != "" {
		params["app_logo"] = biz.AppLogo
	}
	if biz.SpeedUp != "" {
		params["speed_up"] = biz.SpeedUp
	}
	if biz.AutoOnline != "" {
		params["auto_online"] = biz.AutoOnline
	}
	if biz.LicenseValidDate != "" {
		params["license_valid_date"] = biz.LicenseValidDate
	}
	if biz.AppVersion != "" {
		params["app_version"] = biz.AppVersion
	}
	if biz.AppName != "" {
		params["app_name"] = biz.AppName
	}
	if biz.AppEnglishName != "" {
		params["app_english_name"] = biz.AppEnglishName
	}
	if biz.AppSlogan != "" {
		params["app_slogan"] = biz.AppSlogan
	}
	if biz.AppCategoryIDs != "" {
		params["app_category_ids"] = biz.AppCategoryIDs
	}
	if biz.AppDesc != "" {
		params["app_desc"] = biz.AppDesc
	}
	if biz.ServicePhone != "" {
		params["service_phone"] = biz.ServicePhone
	}
	if biz.ServiceEmail != "" {
		params["service_email"] = biz.ServiceEmail
	}
	if biz.VersionDesc != "" {
		params["version_desc"] = biz.VersionDesc
	}
	if biz.Memo != "" {
		params["memo"] = biz.Memo
	}
	if biz.RegionType != "" {
		params["region_type"] = biz.RegionType
	}
	if biz.ServiceRegionInfo != nil {
		serviceRegionInfo, _ := json.Marshal(biz.ServiceRegionInfo)
		params["service_region_info"] = string(serviceRegionInfo)
	}
	if biz.LicenseNo != "" {
		params["license_no"] = biz.LicenseNo
	}
	if biz.MiniCategoryIDs != "" {
		params["mini_category_ids"] = biz.MiniCategoryIDs
	}
	if biz.TestAccount != "" {
		params["test_accout"] = biz.TestAccount
	}
	if biz.TestPassword != "" {
		params["test_password"] = biz.TestPassword
	}
	if biz.BundleID != "" {
		params["bundle_id"] = biz.BundleID
	}
	if biz.FirstLicensePic != "" {
		params["first_license_pic"] = biz.FirstLicensePic
	}
	if biz.SecondLicensePic != "" {
		params["second_license_pic"] = biz.SecondLicensePic
	}
	if biz.ThirdLicensePic != "" {
		params["third_license_pic"] = biz.ThirdLicensePic
	}
	if biz.FourthLicensePic != "" {
		params["fourth_license_pic"] = biz.FourthLicensePic
	}
	if biz.FifthLicensePic != "" {
		params["fifth_license_pic"] = biz.FifthLicensePic
	}
	if biz.OutDoorPic != "" {
		params["out_door_pic"] = biz.OutDoorPic
	}
	if biz.FirstScreenShot != "" {
		params["first_screen_shot"] = biz.FirstScreenShot
	}
	if biz.SecondScreenShot != "" {
		params["second_screen_shot"] = biz.SecondScreenShot
	}
	if biz.ThirdScreenShot != "" {
		params["third_screen_shot"] = biz.ThirdScreenShot
	}
	if biz.FourthScreenShot != "" {
		params["fourth_screen_shot"] = biz.FourthScreenShot
	}
	if biz.FifthScreenShot != "" {
		params["fifth_screen_shot"] = biz.FifthScreenShot
	}
	if biz.FirstSpecialLicensePic != "" {
		params["first_special_license_pic"] = biz.FirstSpecialLicensePic
	}
	if biz.SecondSpecialLicensePic != "" {
		params["second_special_license_pic"] = biz.SecondSpecialLicensePic
	}
	if biz.ThirdSpecialLicensePic != "" {
		params["third_special_license_pic"] = biz.ThirdSpecialLicensePic
	}
	if len(biz.TestFileName) > 0 {
		params["test_file_name"] = biz.TestFileName
	}
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	applyVersionAuditResp := new(ApplyVersionAuditResp)
	_, err = s.Client.Do(ctx, req, applyVersionAuditResp)
	if err != nil {
		return err
	}
	return nil
}

// CancelVersionAuditBiz 小程序撤销审核
type CancelVersionAuditBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 可不选, 默认撤消正在审核中的版本
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端(com.alipay.alipaywallet:支付宝端)
}

// CancelVersionAudit 小程序撤销审核
func (s *MiniService) CancelVersionAudit(ctx context.Context, biz *CancelVersionAuditBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.audit.cancel"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// CancelVersionAuditedBiz 小程序退回开发
type CancelVersionAuditedBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号
	BundleID   string `json:"bundle_id"`   //小程序投放的端参数，例如投放到支付宝钱包是支付宝端。该参数可选，默认支付宝端，目前仅支持支付宝端，枚举列举：com.alipay.alipaywallet:支付宝端
}

// CancelVersionAudited 小程序退回开发
func (s *MiniService) CancelVersionAudited(ctx context.Context, biz *CancelVersionAuditedBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.audited.cancel"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// OnlineVersionBiz 小程序上架
type OnlineVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// OnlineVersion 小程序上架
func (s *MiniService) OnlineVersion(ctx context.Context, biz *OnlineVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.online"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// OfflineVersionBiz 小程序下架
type OfflineVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// OfflineVersion 小程序下架
func (s *MiniService) OfflineVersion(ctx context.Context, biz *OfflineVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.offline"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// RollbackVersionBiz 小程序回滚
type RollbackVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// RollbackVersion 小程序回滚
func (s *MiniService) RollbackVersion(ctx context.Context, biz *RollbackVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.rollback"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// OnlineGrayVersionBiz 小程序灰度上架
type OnlineGrayVersionBiz struct {
	AppVersion   string `json:"app_version"`   //小程序版本号, 必选
	GrayStrategy string `json:"gray_strategy"` //小程序灰度策略值，支持p10，p30，p50, 代表百分之多少的用户
	BundleID     string `json:"bundle_id"`     //端参数，可不选，默认支付宝端
}

// OnlineGrayVersion 小程序灰度上架
func (s *MiniService) OnlineGrayVersion(ctx context.Context, biz *OnlineGrayVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.gray.online"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// CancelGrayVersionBiz 小程序结束灰度
type CancelGrayVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// CancelGrayVersion 小程序灰度上架
func (s *MiniService) CancelGrayVersion(ctx context.Context, biz *CancelGrayVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.gray.cancel"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// UploadVersionBiz 小程序基于模板上传版本
type UploadVersionBiz struct {
	AppVersion      string `json:"app_version"`      //小程序版本号, 必选
	BundleID        string `json:"bundle_id"`        //端参数，可不选，默认支付宝端
	TemplateID      string `json:"template_id"`      //模板id
	Ext             string `json:"ext"`              //模板的配置参数
	TemplateVersion string `json:"template_version"` //模板版本号，版本号必须满足 x.y.z, 且均为数字。不传默认使用最新在架模板版本。
}

// UploadVersion 小程序基于模板上传版本
func (s *MiniService) UploadVersion(ctx context.Context, biz *UploadVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.upload"
	param := map[string]interface{}{}
	if biz.AppVersion != "" {
		param["app_version"] = biz.AppVersion
	}
	if biz.BundleID != "" {
		param["bundle_id"] = biz.BundleID
	}
	if biz.TemplateID != "" {
		param["template_id"] = biz.TemplateID
	}
	if biz.Ext != "" {
		param["ext"] = biz.Ext
	}
	if biz.TemplateVersion != "" {
		param["template_version"] = biz.TemplateVersion
	}
	req, err := s.Client.NewRequest(apiMethod, param, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// QueryVersionDetailBiz 小程序版本详情查询
type QueryVersionDetailBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// MiniAppCategoryInfo 小程序类目
type MiniAppCategoryInfo struct {
	FirstCategoryID    string `json:"first_category_id"`
	FirstCategoryName  string `json:"first_category_name"`
	SecondCategoryID   string `json:"second_category_id"`
	SecondCategoryName string `json:"second_category_name"`
	ThirdCategoryId    string `json:"third_category_id"`
	ThirdCategoryName  string `json:"third_category_name"`
}

// MiniPackageInfo 小程序功能包
type MiniPackageInfo struct {
	PackageName     string `json:"package_name"`
	PackageDesc     string `json:"package_desc"`
	DocURL          string `json:"doc_url"`
	Status          string `json:"status"`
	PackageOpenType string `json:"package_open_type"`
}

type VersionDetail struct {
	AlipayOpenMiniVersionDetailQueryResponse struct {
		Code                    string                `json:"code"`
		Msg                     string                `json:"msg"`
		AppVersion              string                `json:"app_version"`
		AppName                 string                `json:"app_name"`
		GmtCreate               string                `json:"gmt_create"`
		Status                  string                `json:"status"`
		AppEnglishName          string                `json:"app_english_name"`
		AppLogo                 string                `json:"app_logo"`
		VersionDesc             string                `json:"version_desc"`
		GrayStrategy            string                `json:"gray_strategy"`
		RejectReason            string                `json:"reject_reason"`
		ScanResult              string                `json:"scan_result"`
		GmtApplyAudit           string                `json:"gmt_apply_audit"`
		GmtOnline               string                `json:"gmt_online"`
		GmtOffline              string                `json:"gmt_offline"`
		GmtAuditEnd             string                `json:"gmt_audit_end"`
		AppDesc                 string                `json:"app_desc"`
		ServiceRegionType       string                `json:"service_region_type"`
		AppSlogan               string                `json:"app_slogan"`
		Memo                    string                `json:"memo"`
		ServicePhone            string                `json:"service_phone"`
		ServiceEmail            string                `json:"service_email"`
		MiniAppCategoryInfoList []MiniAppCategoryInfo `json:"mini_app_category_info_list"`
		PackageInfoList         []MiniPackageInfo     `json:"package_info_list"`
		ServiceRegionInfo       []RegionInfo          `json:"service_region_info"`
		ScreenShotList          []string              `json:"screen_shot_list"`
	} `json:"alipay_open_mini_version_detail_query_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// QueryVersionDetail 小程序版本详情查询
func (s *MiniService) QueryVersionDetail(ctx context.Context, biz *QueryVersionDetailBiz, opts ...ValueOptions) (*VersionDetail, error) {
	apiMethod := "alipay.open.mini.version.detail.query"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	versionDetail := new(VersionDetail)
	_, err = s.Client.Do(ctx, req, versionDetail)
	if err != nil {
		return nil, err
	}
	return versionDetail, nil
}

// QueryVersionBuildBiz 小程序查询版本构建状态
type QueryVersionBuildBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

type QueryVersionBuildResp struct {
	AlipayOpenMiniVersionBuildQueryResponse struct {
		Code           string `json:"code"`
		Msg            string `json:"msg"`
		VersionCreated string `json:"version_created"` // 版本是否创建成功
		CreateStatus   string `json:"create_status"`   // 版本构建状态
		NeedRotation   string `json:"need_rotation"`   // 是否需要轮询
		BuildStatus    string `json:"build_status"`    // 构建状态
	} `json:"alipay_open_mini_version_build_query_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// QueryVersionBuild 小程序查询版本构建状态
func (s *MiniService) QueryVersionBuild(ctx context.Context, biz *QueryVersionBuildBiz, opts ...ValueOptions) (*QueryVersionBuildResp, error) {
	apiMethod := "alipay.open.mini.version.build.query"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(QueryVersionBuildResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
