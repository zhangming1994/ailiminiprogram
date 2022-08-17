package aliminiprogram

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// AliMiniProgram  支付宝小程序
type AliMiniProgram struct {
	PrivateKey *rsa.PrivateKey // 私钥
	PublicKey  *rsa.PublicKey  // 公钥
	AppId      string          // 开发者appid
	RootCertSn string          // 根证书SN
	AppCertSn  string          // 应用公钥证书 SN
}

// NewAliMiniProgram  private 私钥 public 支付宝公钥证书 rootCertSN 根证书SN appCertSN 应用公钥证书
func NewAliMiniProgram(appId, private, public, rootCertSN, appCertSN string) (*AliMiniProgram, error) {
	encodedKey, err := base64.StdEncoding.DecodeString(private)
	if err != nil {
		fmt.Println("decode private key error", err)
		return nil, err
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(encodedKey)
	if err != nil {
		fmt.Println("x509 parse private key error", err)
		return nil, err
	}
	publicKey, err := base64.StdEncoding.DecodeString(public)
	if err != nil {
		fmt.Println("decode public key error", err)
		return nil, err
	}
	pub, err := x509.ParsePKCS1PublicKey(publicKey)
	if err != nil {
		fmt.Println("x509 parse public key error", err)
		return nil, err
	}
	return &AliMiniProgram{
		PrivateKey: privateKey,
		PublicKey:  pub,
		AppId:      appId,
		RootCertSn: rootCertSN,
		AppCertSn:  appCertSN,
	}, nil
}

// NewClient 小程序 文体中心 server等client
func (a *AliMiniProgram) NewClient() *Client {
	return NewClient(nil, a.PrivateKey, a.PublicKey, AppID(a.AppId), AppCertSN(a.AppCertSn), RootCertSN(a.RootCertSn))
}
