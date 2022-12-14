package aliminiprogram

import (
	"context"
)

// CreateExperienceBiz 小程序生成体验版
type CreateExperienceBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号
	// 小程序客户端类型，默认为支付宝端。
	// 支付宝端：com.alipay.alipaywallet,
	// DINGDING端：com.alibaba.android.rimet,
	// 高德端:com.amap.app,
	// 天猫精灵端:com.alibaba.ailabs.genie.webapps,
	// 支付宝IOT:com.alipay.iot.xpaas
	BundleID string `json:"bundle_id"`
}

type CreateExperienceReq struct {
	AlipayOpenMiniExperienceCreateResponse struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	} `json:"alipay_open_mini_experience_create_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// CreateExperience 小程序生成体验版
func (s *MiniService) CreateExperience(ctx context.Context, biz *CreateExperienceBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.experience.create"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	createExperience := new(CreateExperienceReq)
	_, err = s.Client.Do(ctx, req, createExperience)
	if err != nil {
		return err
	}
	return nil
}

// QueryExperienceBiz 小程序体验版状态查询
type QueryExperienceBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号
	// 小程序客户端类型，默认为支付宝端。
	// 支付宝端：com.alipay.alipaywallet,
	// DINGDING端：com.alibaba.android.rimet,
	// 高德端:com.amap.app,
	// 天猫精灵端:com.alibaba.ailabs.genie.webapps,
	// 支付宝IOT:com.alipay.iot.xpaas
	BundleID string `json:"bundle_id"`
}

// ExperienceStatus 体验版状态
type ExperienceStatus struct {
	AlipayOpenMiniExperienceQueryResponse struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		ExpQrCodeURL string `json:"exp_qr_code_url"`
		Status       string `json:"status"` // 体验版打包状态，expVersionPackaged-体验版打包成功，expVersionPackaging-体验版打包中，notExpVersion-非体验版
	} `json:"alipay_open_mini_experience_query_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// QueryExperience 小程序体验版状态查询
func (s *MiniService) QueryExperience(ctx context.Context, biz *QueryExperienceBiz, opts ...ValueOptions) (*ExperienceStatus, error) {
	apiMethod := "alipay.open.mini.experience.query"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	experienceStatus := new(ExperienceStatus)
	_, err = s.Client.Do(ctx, req, experienceStatus)
	if err != nil {
		return nil, err
	}
	return experienceStatus, nil
}

// CancelExperienceBiz 小程序取消体验版
type CancelExperienceBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号
	// 小程序客户端类型，默认为支付宝端。
	// 支付宝端：com.alipay.alipaywallet,
	// DINGDING端：com.alibaba.android.rimet,
	// 高德端:com.amap.app,
	// 天猫精灵端:com.alibaba.ailabs.genie.webapps,
	// 支付宝IOT:com.alipay.iot.xpaas
	BundleID string `json:"bundle_id"`
}

// CancelExperience 小程序取消体验版
func (s *MiniService) CancelExperience(ctx context.Context, biz *CancelExperienceBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.experience.cancel"
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
