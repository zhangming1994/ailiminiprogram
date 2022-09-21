package aliminiprogram

import (
	"context"
)

// CreateProgramParam isv代商户创建小程序
type CreateProgramParam struct {
	CreateMiniRequest struct {
		AliPayAccount     string `json:"alipay_account"`      // 企业支付宝账号
		LegalPersonalName string `json:"legal_personal_name"` // 商家法人名称
		CertName          string `json:"cert_name"`           // 企业营业执照名称
		CertNO            string `json:"cert_no"`             // 企业营业执照编码
		AppName           string `json:"app_name"`            // 小程序名称
		ContactPhone      string `json:"contact_phone"`       // 商家联系人电话号码
		ContactName       string `json:"contact_name"`        // 商家联系人名称
		OutOrderNO        string `json:"out_order_no"`
	} `json:"create_mini_request"`
}

type CreateProgramResp struct {
	AlipayOpenMiniIsvCreateResponse struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		OrderNo string `json:"order_no"`
	} `json:"alipay_open_mini_isv_create_response"`
	Sign string `json:"sign"`
}

// IsvCreateMini isv创建小程序
func (s *MiniService) IsvCreateMini(ctx context.Context, biz *CreateProgramParam, opts ...ValueOptions) (*CreateProgramResp, error) {
	apiMethod := "alipay.open.mini.isv.create"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(CreateProgramResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
