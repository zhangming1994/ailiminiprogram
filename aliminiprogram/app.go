package aliminiprogram

import (
	"context"
)

// AppService 应用服务
// Docs: https://opendocs.alipay.com/apis/api_49
type AppService Service

// CreateAppMemberBiz 应用添加成员
type CreateAppMemberBiz struct {
	LogonID string `json:"logon_id"` // 支付宝登录账号ID
	Role    string `json:"role"`     //成员的角色类型，DEVELOPER-开发者，EXPERIENCER-体验者
}

// CreateMember 应用添加成员，目前只支持小程序类型的应用使用
func (s *AppService) CreateMember(ctx context.Context, biz *CreateAppMemberBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.app.members.create"
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

// DeleteMemberBiz 应用删除成员
type DeleteMemberBiz struct {
	UserID string `json:"user_id"` // 蚂蚁统一会员ID
	Role   string `json:"role"`    //成员的角色类型，DEVELOPER-开发者，EXPERIENCER-体验者
}

// DeleteMember 应用删除成员，目前只支持小程序类型的应用使用
func (s *AppService) DeleteMember(ctx context.Context, biz *DeleteMemberBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.app.members.delete"
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

// QueryAppMembersBiz 应用查询成员列表
type QueryAppMembersBiz struct {
	Role string `json:"role"` //成员的角色类型，DEVELOPER-开发者，EXPERIENCER-体验者
}

type QueryAppMembersResp struct {
	AlipayOpenAppMembersQueryResponse struct {
		Code              string `json:"code"`
		Msg               string `json:"msg"`
		AppMemberInfoList []struct {
			UserID    string `json:"user_id"`
			NickName  string `json:"nick_name"`
			Portrait  string `json:"portrait"`
			Status    string `json:"status"`
			GmtJoin   string `json:"gmt_join"`
			LogonID   string `json:"logon_id"`
			GmtInvite string `json:"gmt_invite"`
			Role      string `json:"role"`
		} `json:"app_member_info_list"`
	} `json:"alipay_open_app_members_query_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// QueryAppMembers 应用查询成员列表，目前只支持小程序类型的应用
func (s *AppService) QueryAppMembers(ctx context.Context, biz *QueryAppMembersBiz, opts ...ValueOptions) (*QueryAppMembersResp, error) {
	apiMethod := "alipay.open.app.members.query"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(QueryAppMembersResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateAppQRCodeResp 生成小程序推广二维码返回值
type CreateAppQRCodeResp struct {
	AlipayOpenAppQrcodeCreateResponse struct {
		Code                 string `json:"code"`
		Msg                  string `json:"msg"`
		QrCodeUrl            string `json:"qr_code_url"`
		QrCodeUrlCircleWhite string `json:"qr_code_url_circle_white"`
		QrCodeUrlCircleBlue  string `json:"qr_code_url_circle_blue"`
	} `json:"alipay_open_app_qrcode_create_response"`
	Sign string `json:"sign"`
}

// CreateAppQRCodeBiz 生成小程序推广二维码
type CreateAppQRCodeBiz struct {
	URLParam   string `json:"url_param"`   //小程序中能访问到的页面路径。
	QueryParam string `json:"query_param"` //小程序的启动参数，打开小程序的query，在小程序onLaunch的方法中获取。
	Describe   string `json:"describe"`    //对应的二维码描述。
}

// CreateAppQRCode 生成小程序推广二维码
func (s *AppService) CreateAppQRCode(ctx context.Context, biz *CreateAppQRCodeBiz, opts ...ValueOptions) (*CreateAppQRCodeResp, error) {
	apiMethod := "alipay.open.app.qrcode.create"
	req, err := s.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(CreateAppQRCodeResp)
	_, err = s.Client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
