package aliminiprogram

import (
	"bytes"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"
)

var alog map[string]string = map[string]string{
	"MD2-RSA":       "MD2WithRSA",
	"MD5-RSA":       "MD5WithRSA",
	"SHA1-RSA":      "SHA1WithRSA",
	"SHA256-RSA":    "SHA256WithRSA",
	"SHA384-RSA":    "SHA384WithRSA",
	"SHA512-RSA":    "SHA512WithRSA",
	"SHA256-RSAPSS": "SHA256WithRSAPSS",
	"SHA384-RSAPSS": "SHA384WithRSAPSS",
	"SHA512-RSAPSS": "SHA512WithRSAPSS",
}

// GetCertSn 获取根证书序列号 alipay_root_cert_sn 应用公钥证书SN  app_cert_sn
func GetCertSn(certData []byte) (string, error) {
	strs := strings.Split(string(certData), "-----END CERTIFICATE-----")
	var cert bytes.Buffer
	for i := 0; i < len(strs); i++ {
		if strs[i] == "" {
			continue
		}
		if blo, _ := pem.Decode([]byte(strs[i] + "-----END CERTIFICATE-----")); blo != nil {
			c, err := x509.ParseCertificate(blo.Bytes)
			if err != nil {
				continue
			}
			if _, ok := alog[c.SignatureAlgorithm.String()]; !ok {
				continue
			}
			si := c.Issuer.String() + c.SerialNumber.String()
			if cert.String() == "" {
				cert.WriteString(md5Encode(si))
			} else {
				cert.WriteString("_")
				cert.WriteString(md5Encode(si))
			}
		}
	}
	return cert.String(), nil
}

func md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// CertProcess 证书处理
func CertProcess(privateKey, alipayCertPublicKeyRSA2, rootCert, appCert string) (private *rsa.PrivateKey, publicKey *rsa.PublicKey, rootCertSN, appCertSN string, err error) {
	encodedKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return private, publicKey, rootCertSN, appCertSN, err
	}
	private, err = x509.ParsePKCS1PrivateKey(encodedKey)
	if err != nil {
		return private, publicKey, rootCertSN, appCertSN, err
	}

	aliMiniCertPublicKey, err := base64.StdEncoding.DecodeString(alipayCertPublicKeyRSA2)
	if err != nil {
		return private, publicKey, rootCertSN, appCertSN, err
	}
	blockCert, _ := pem.Decode(aliMiniCertPublicKey)
	x509Cert, err := x509.ParseCertificate(blockCert.Bytes)
	if err != nil {
		return private, publicKey, rootCertSN, appCertSN, err
	}
	pub, ok := x509Cert.PublicKey.(*rsa.PublicKey)
	if ok {
		//publicKey = base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(pub))
		publicKey = pub
	} else {
		err = fmt.Errorf("x509 publickeyRSA2 assert error")
		return private, publicKey, rootCertSN, appCertSN, err
	}

	// 根证书sn
	aliMinirRootCert, err := base64.StdEncoding.DecodeString(rootCert)
	if err != nil {
		return private, publicKey, rootCertSN, appCertSN, err
	}
	rootCertSN, err = GetCertSn(aliMinirRootCert)
	if err != nil {
		return private, publicKey, rootCertSN, appCertSN, err
	}

	// appcertsn
	aliMiniAppCert, err := base64.StdEncoding.DecodeString(appCert)
	if err != nil {
		return private, publicKey, rootCertSN, appCertSN, err
	}
	appCertSN, err = GetCertSn(aliMiniAppCert)
	if err != nil {
		return private, publicKey, rootCertSN, appCertSN, err
	}
	return
}
