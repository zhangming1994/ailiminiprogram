package aliminiprogram

import (
	"context"
)

// MiniPayMessageConfirm 小程序支付消息确认接口 调用一次就ok 这是小程序内部完成支付触发消息
func (s *MiniService) MiniPayMessageConfirm(ctx context.Context, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.templatemsg.tinypayswitch.confirm"
	req, err := s.Client.NewRequest(apiMethod, nil, opts...)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

type PageDefine struct {
	PageSize   int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
}

type MerchantPID struct {
	AntMerchantExpandApprecommendAvailableQueryResponse struct {
		Code       string   `json:"code"`
		Msg        string   `json:"msg"`
		List       []string `json:"list"`
		PageNumber int      `json:"page_number"`
		PageSize   int      `json:"page_size"`
		TotalPages int      `json:"total_pages"`
		TotalSize  int      `json:"total_size"`
	} `json:"ant_merchant_expand_apprecommend_available_query_response"`
	Sign string `json:"sign"`
}

// MerchantAvailableQuery ====== 下面是非小程序支付消息(线下支付,代扣等触发消息) 先查询可关联的PID接口 然后进行关联[也可以取消关联 查询已经关联的]======
func (s *MiniService) MerchantAvailableQuery(ctx context.Context, biz PageDefine, opts ...ValueOptions) (MerchantPID, error) {
	apiMethod := "ant.merchant.expand.apprecommend.available.query"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return MerchantPID{}, err
	}
	resp := new(MerchantPID)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return MerchantPID{}, err
	}
	return *resp, nil
}

type BindAppPID struct {
	AppNO string `json:"app_no"`
	AccNO string `json:"acc_no"`
}

// BindAppAndPID 关联商户和小程序
func (s *MiniService) BindAppAndPID(ctx context.Context, biz BindAppPID, opts ...ValueOptions) error {
	apiMethod := "ant.merchant.expand.apprecommend.account.create"
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

type BindAccountQueryReq struct {
	AppNO      string `json:"app_no"`
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
}

type BindAccountQueryResp struct {
	AntMerchantExpandApprecommendAccountQueryResponse struct {
		Code       string   `json:"code"`
		Msg        string   `json:"msg"`
		List       []string `json:"list"`
		PageNumber int      `json:"page_number"`
		PageSize   int      `json:"page_size"`
		TotalPages int      `json:"total_pages"`
		TotalSize  string   `json:"total_size"`
	} `json:"ant_merchant_expand_apprecommend_account_query_response"`
	Sign string `json:"sign"`
}

// BindAccountQuery 查询已经关联指定APP的账号列表
func (s *MiniService) BindAccountQuery(ctx context.Context, biz BindAccountQueryReq, opts ...ValueOptions) (BindAccountQueryResp, error) {
	apiMethod := "ant.merchant.expand.apprecommend.account.query"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return BindAccountQueryResp{}, err
	}
	resp := new(BindAccountQueryResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return BindAccountQueryResp{}, err
	}
	return *resp, nil
}

// BindAccountDelete 删除小程序和账号的绑定关系
func (s *MiniService) BindAccountDelete(ctx context.Context, biz BindAppPID, opts ...ValueOptions) error {
	apiMethod := "ant.merchant.expand.apprecommend.account.delete"
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
