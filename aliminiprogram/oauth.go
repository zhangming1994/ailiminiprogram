package aliminiprogram

import (
	"context"
	"git.myarena7.com/arena/lib/log"
	"git.myarena7.com/arena/platform/conf"
	"net/url"
)

type OAuth Service

// OAuthToken 换取授权访问令牌(包含用户授权访问令牌和应用授权令牌)
type OAuthToken struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

// OAuthTokenResp 用户授权访问令牌返回参数
type OAuthTokenResp struct {
	AlipaySystemOauthTokenResponse struct {
		AccessToken  string `json:"access_token"` // 访问令牌。通过该令牌调用需要授权类接口
		AlipayUserId string `json:"alipay_user_id"`
		AuthStart    string `json:"auth_start"`    // 授权token开始时间，作为有效期计算的起点
		ExpiresIn    int    `json:"expires_in"`    //  访问令牌的有效时间，单位是秒。
		ReExpiresIn  int    `json:"re_expires_in"` // 刷新令牌的有效时间，单位是秒。 用redis存储 过期时间就设置为这个
		RefreshToken string `json:"refresh_token"` // 刷新令牌。通过该令牌可以刷新access_token
		UserId       string `json:"user_id"`       // 支付宝用户的唯一标识。以2088开头的16位数字
	} `json:"alipay_system_oauth_token_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// MerchantOAuthTokenResp 换取应用授权令牌的返回参数
type MerchantOAuthTokenResp struct {
	AliPayCertSn string `json:"alipay_cert_sn"`
	AliTokenBack struct {
		Token []Token `json:"tokens"` //
		Code  string  `json:"code"`
		Msg   string  `json:"msg"`
	} `json:"alipay_open_auth_token_app_response"`
	Sign string `json:"sign"`
}

type Token struct {
	UserId          string `json:"user_id"`           // 授权商户的userid
	AuthAppId       string `json:"auth_app_id"`       // 授权商户的appid
	AppAuthToken    string `json:"app_auth_token"`    // 应用授权令牌
	AppRefreshToken string `json:"app_refresh_token"` // 刷新令牌
	ReExpiresIn     int64  `json:"re_expires_in"`     // 刷新令牌的有效时间（从接口调用时间作为起始时间），单位到秒
	ExpiresIn       int64  `json:"expires_in"`        // 该字段已作废，应用令牌长期有效，接入方不需要消费该字段
}

type UserInfoResp struct {
	AlipayUserInfoShareResponse struct {
		Code     string `json:"code"`
		Msg      string `json:"msg"`
		Avatar   string `json:"avatar"`
		NickName string `json:"nick_name"`
		UserId   string `json:"user_id"`
	} `json:"alipay_user_info_share_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// GetUserInfo  获取用户信息
func (o *OAuth) GetUserInfo(ctx context.Context, authToken string, opts ...ValueOptions) (UserInfoResp, error) {
	method := "alipay.user.info.share"
	value := url.Values{}
	value.Set("auth_token", authToken)
	req, err := o.Client.NewGetRequest(method, value, opts...)
	if err != nil {
		log.Ctx(ctx).DebugAsJSON(err.Error(), "获取用户信息失败")
		return UserInfoResp{}, err
	}
	var oAuthTokenResp UserInfoResp
	_, err = o.Client.Do(ctx, req, &oAuthTokenResp)
	if err != nil {
		log.Ctx(ctx).DebugAsJSON(err.Error(), "获取用户信息失败")
		return UserInfoResp{}, err
	}
	return oAuthTokenResp, nil
}

// MerchantOAuthToken 换取应用授权令牌
func (o *OAuth) MerchantOAuthToken(ctx context.Context, biz OAuthToken, opts ...ValueOptions) (MerchantOAuthTokenResp, error) {
	method := "alipay.open.auth.token.app"
	req, err := o.Client.NewRequest(method, biz, opts...)
	if err != nil {
		log.Ctx(ctx).DebugAsJSON(err.Error(), "换取商户应用授权令牌")
		return MerchantOAuthTokenResp{}, err
	}
	var oAuthTokenResp MerchantOAuthTokenResp
	_, err = o.Client.Do(ctx, req, &oAuthTokenResp)
	if err != nil {
		log.Ctx(ctx).DebugAsJSON(err.Error(), "换取商户应用授权令牌失败")
		return MerchantOAuthTokenResp{}, err
	}
	return oAuthTokenResp, nil
}

// UserOAuthToken 换取用户授权访问令牌
func (o *OAuth) UserOAuthToken(ctx context.Context, biz OAuthToken, opts ...ValueOptions) (OAuthTokenResp, error) {
	method := "alipay.system.oauth.token"
	param := url.Values{}
	param.Set("grant_type", "authorization_code")
	param.Set("code", biz.Code)
	req, err := o.Client.NewGetRequest(method, param, opts...)
	if err != nil {
		log.Ctx(ctx).DebugAsJSON(err.Error(), "换取用户应用授权令牌失败")
		return OAuthTokenResp{}, err
	}
	var oAuthTokenResp OAuthTokenResp
	_, err = o.Client.Do(ctx, req, &oAuthTokenResp)
	if err != nil {
		log.Ctx(ctx).DebugAsJSON(err.Error(), "换取用户应用授权令牌失败")
		return OAuthTokenResp{}, err
	}
	return oAuthTokenResp, nil
}

type SetTemplateAesKey struct {
	AesKey string `json:"aes_key"`
}

// SetTemplateAesKey  为模板设置aes FIXME 支持多个模板配置
func (o *OAuth) SetTemplateAesKey(ctx context.Context, opts ...ValueOptions) (bool, error) {
	method := "alipay.open.auth.app.aes.set"
	biz := struct {
		MerchantAppId string `json:"merchant_app_id"`
		EncryptKey    string `json:"encrypt_key"`
	}{
		EncryptKey:    conf.GetAliMiniProgramConf().TemplateAesKey,
		MerchantAppId: conf.GetAliMiniProgramConf().TemplateAppId,
	}
	req, err := o.Client.NewRequest(method, biz, opts...)
	if err != nil {
		log.Ctx(ctx).DebugAsJSON(err.Error(), "为模板设置AESKEY失败")
		return false, err
	}
	var templateAesKey SetTemplateAesKey

	_, err = o.Client.Do(ctx, req, &templateAesKey)
	if err != nil {
		log.Ctx(ctx).DebugAsJSON(err.Error(), "为模板设置AESKEY失败")
		return false, err
	}
	return true, nil
}
