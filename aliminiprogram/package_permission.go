package aliminiprogram

import (
	"context"
)

type AppApiQueryResp struct {
	AlipayOpenAppApiQueryResponse struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
		Apis []struct {
			ApiName     string `json:"api_name"`
			FieldName   string `json:"field_name"`
			PackageCode string `json:"package_code"`
		} `json:"apis"`
	} `json:"alipay_open_app_api_query_response"`
	Sign string `json:"sign"`
}

// AppApiQuery 查询可申请权限
func (s *MiniService) AppApiQuery(ctx context.Context, opts ...ValueOptions) (*AppApiQueryResp, error) {
	apiMethod := "alipay.open.app.api.query"
	req, err := s.Client.NewRequest(apiMethod, nil, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(AppApiQueryResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApplyPermission struct {
	AuthFieldApply AuthFieldApply `json:"auth_field_apply"`
	Picture1       []byte         `json:"picture_1"`
	Picture2       []byte         `json:"picture_2"`
	Picture3       []byte         `json:"picture_3"`
	Picture4       []byte         `json:"picture_4"`
	Picture5       []byte         `json:"picture_5"`
}

type AuthFieldApply struct {
	ApiName        string `json:"api_name"`        // 接口英文名称，通过alipay.open.app.api.query接口查询获得。
	FieldName      string `json:"field_name"`      // 字段英文名称，通过alipay.open.app.api.query接口查询获得。
	PackageCode    string `json:"package_code"`    // 功能code，通过alipay.open.app.api.query接口查询获得
	SceneCode      string `json:"scene_code"`      // 场景code，alipay.open.app.api.scene.query 接口查询获得
	QpsAnswer      string `json:"qps_answer"`      // 接入后一年内预计接口秒级调用量峰值是多少？（最高峰值：1000QPS）
	CustomerAnswer string `json:"customer_answer"` // 贵公司是否有自己的客服渠道，能及时响应和处理舆情数量是多少？
	Memo           string `json:"memo"`            // 获取用途
}

// ApplyPermission 申请权限
func (s *MiniService) ApplyPermission(ctx context.Context, biz *ApplyPermission, opts ...ValueOptions) error {
	apiMethod := "alipay.open.app.api.field.apply"
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

type QueryFieldScene struct {
	FieldName string `json:"field_name"`
	ApiName   string `json:"api_name"`
}

type AuthFieldSceneDTO struct {
	AlipayOpenAppApiSceneQueryResponse struct {
		Code           string `json:"code"`
		Msg            string `json:"msg"`
		AuthFieldScene []struct {
			SceneCode string `json:"scene_code"`
			SceneDesc string `json:"scene_desc"`
		} `json:"auth_field_scene"`
	} `json:"alipay_open_app_api_scene_query_response"`
	Sign string `json:"sign"`
}

// QueryFieldScene 查询接口字段使用场景
func (s *MiniService) QueryFieldScene(ctx context.Context, biz *QueryFieldScene, opts ...ValueOptions) (AuthFieldSceneDTO, error) {
	apiMethod := "alipay.open.app.api.scene.query"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return AuthFieldSceneDTO{}, err
	}
	resp := new(AuthFieldSceneDTO)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return AuthFieldSceneDTO{}, err
	}
	return *resp, nil
}

type FieldQueryRecordResp struct {
	AlipayOpenAppApiFieldQueryResponse struct {
		Code              string `json:"code"`
		Msg               string `json:"msg"`
		AuthFieldResponse struct {
			Records []struct {
				UserAppId   string `json:"user_app_id"`
				ApiName     string `json:"api_name"`
				FieldName   string `json:"field_name"`
				Status      string `json:"status"`
				Reason      string `json:"reason"`
				PackageCode string `json:"package_code"`
			} `json:"records"`
		} `json:"auth_field_response"`
	} `json:"alipay_open_app_api_field_query_response"`
	Sign string `json:"sign"`
}

// FieldQueryRecord 用户信息申请记录查询
func (s *MiniService) FieldQueryRecord(ctx context.Context, opts ...ValueOptions) (FieldQueryRecordResp, error) {
	apiMethod := "alipay.open.app.api.field.query"
	req, err := s.Client.NewRequest(apiMethod, nil, opts...)
	if err != nil {
		return FieldQueryRecordResp{}, err
	}
	resp := new(FieldQueryRecordResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return FieldQueryRecordResp{}, err
	}
	return *resp, nil
}

type ApiFieldChangedNotify struct {
	UserAppId string `json:"user_app_id"` // 用户app_id
	ApiName   string `json:"api_name"`    // 接口英文名称
	FieldName string `json:"field_name"`  // 接口字段英文名称
	/*
		审核结果。结果枚举：
		AGREE：通过。
		REJECT：驳回。
		INVALID无效（isv代申请场景）。
	*/
	Status      string `json:"status"`       //
	Reason      string `json:"reason"`       /// 驳回原因
	PackageCode string `json:"package_code"` // 接口所属功能code
}

type MerchantAesSetResp struct {
	AlipayOpenAuthAppAesSetResponse struct {
		Code   string `json:"code"`
		Msg    string `json:"msg"`
		AesKey string `json:"aes_key"`
	} `json:"alipay_open_auth_app_aes_set_response"`
	Sign string `json:"sign"`
}

type MerchantAesSetReq struct {
	MerchantAppId string `json:"merchant_app_id"`
}

// MerchantAesKeySet 运营商小程序AESKEY设置
func (s *MiniService) MerchantAesKeySet(ctx context.Context, biz MerchantAesSetReq, opts ...ValueOptions) (MerchantAesSetResp, error) {
	apiMethod := "alipay.open.auth.app.aes.set"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return MerchantAesSetResp{}, err
	}
	resp := new(MerchantAesSetResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return MerchantAesSetResp{}, err
	}
	return *resp, nil
}

type MerchantAesKeyQueryResp struct {
	AlipayOpenAuthAppAesGetResponse struct {
		Code   string `json:"code"`
		Msg    string `json:"msg"`
		AesKey string `json:"aes_key"`
	} `json:"alipay_open_auth_app_aes_get_response"`
	Sign string `json:"sign"`
}

// MerchantAesKeyQuery 商户AESKEY查询
func (s *MiniService) MerchantAesKeyQuery(ctx context.Context, biz MerchantAesSetReq, opts ...ValueOptions) (MerchantAesSetResp, error) {
	apiMethod := "alipay.open.auth.app.aes.get"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return MerchantAesSetResp{}, err
	}
	resp := new(MerchantAesSetResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return MerchantAesSetResp{}, err
	}
	return *resp, nil
}

type AliMiniCategory struct {
	AlipayOpenMiniCategoryQueryResponse struct {
		Code             string `json:"code"`
		Msg              string `json:"msg"`
		MiniCategoryList []struct {
			CategoryId         string `json:"category_id"`
			CategoryName       string `json:"category_name"`
			ParentCategoryId   string `json:"parent_category_id"`
			HasChild           bool   `json:"has_child"`
			NeedLicense        bool   `json:"need_license"`
			NeedOutDoorPic     bool   `json:"need_out_door_pic"`
			NeedSpecialLicense bool   `json:"need_special_license"`
		} `json:"mini_category_list"`
	} `json:"alipay_open_mini_category_query_response"`
	Sign string `json:"sign"`
}
