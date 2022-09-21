package aliminiprogram

import (
	"crypto/rsa"
	"fmt"
)

// AliMiniProgram  支付宝小程序
type AliMiniProgram struct {
	PrivateKey *rsa.PrivateKey // 私钥
	PublicKey  *rsa.PublicKey  // 公钥
	AppId      string          // 开发者appid
	RootCertSn string          // 根证书SN
	AppCertSn  string          // 应用公钥证书SN
}

// NewMiniProgramClient 传入原始的base64值初始化
// private 私钥 alipayPublic 支付宝公钥证书 rootCert 支付宝根证书 appCert 应用公钥证书
func NewMiniProgramClient(appId, private, aliPublicKey, rootCert, appCert string) (*Client, error) {
	if appId == "" || private == "" || aliPublicKey == "" || rootCert == "" || appCert == "" {
		return nil, fmt.Errorf("initialization failurethe,certificate is not complete")
	}
	pri, pub, rootCertSN, appCertSN, err := CertProcess(private, aliPublicKey, rootCert, appCert)
	if err != nil {
		return nil, err
	}
	return NewClient(nil, pri, pub, AppID(appId), AppCertSN(appCertSN), RootCertSN(rootCertSN)), nil
}
