package aliminiprogram

import (
	"context"
	"log"
	"testing"
)

func initClient() *Client {
	// 应用共要证书 appCertPublicKey_xxxx.crt
	// 支付宝共要证书 alipayCertPublicKey_RSA2.crt
	// private 私钥 public 支付宝公钥证书 rootCertSN 根证书SN appCertSN 应用公钥证书
	aliMiniProgram, err := NewAliMiniProgram("",
		"",
		"",
		"",
		"")
	if err != nil {
		log.Panic(err.Error())
	}
	return aliMiniProgram.NewClient()
}

func TestUploadVersion(t *testing.T) {
	client := initClient()
	var oauth UploadVersionBiz
	oauth.AppVersion = "0.0.3"
	oauth.TemplateID = ""
	oauth.TemplateVersion = "0.0.5"
	err := client.Mini.UploadVersion(context.Background(), &oauth, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetAccessToken(t *testing.T) {
	client := initClient()
	var oauth OAuthToken
	oauth.GrantType = "authorization_code"
	oauth.Code = ""
	resp, err := client.OAuth.UserOAuthToken(context.Background(), oauth)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestBaseInfoQuery(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryBaseInfo(context.Background(), AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestVersionGet(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryVersionList(context.Background(), QueryVersionListBiz{VersionStatus: "INIT"}, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestBuildVersionStatus(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryVersionBuild(context.Background(), &QueryVersionBuildBiz{AppVersion: "0.0.2"}, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestVersionDetail(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryVersionDetail(context.Background(), &QueryVersionDetailBiz{
		AppVersion: "0.0.1",
	}, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestAuditApply(t *testing.T) {
	client := initClient()

	err := client.Mini.ApplyVersionAudit(context.Background(), &ApplyVersionAuditBiz{
		AppVersion:   "0.0.1",
		AppLogo:      "https://appstoreisvpic.alipayobjects.com/prod/aa9efcb9-143e-4ebd-bf9c-38c4900ed8a0.png",
		ServicePhone: "13260683758",
		VersionDesc:  "这是一个测试，第一次做这个， 这是一个测试， 第一次做这个 ，这是一个测试， 第一次做这个， 这是一个测试 ，第一次做这个",
		RegionType:   "CHINA",
		AutoOnline:   "false",
		SpeedUp:      "false",
	}, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCancelAudit(t *testing.T) {
	client := initClient()
	err := client.Mini.CancelVersionAudit(context.Background(), &CancelVersionAuditBiz{
		AppVersion: "0.0.1",
	}, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateExperience(t *testing.T) {
	client := initClient()
	err := client.Mini.CreateExperience(context.Background(), &CreateExperienceBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	}, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetExperienceStatus(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryExperience(context.Background(), &QueryExperienceBiz{
		AppVersion: "0.0.2",
		BundleID:   "com.alipay.alipaywallet",
	}, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestAliMiniMembersQuery(t *testing.T) {
	client := initClient()
	resp, err := client.App.QueryAppMembers(context.Background(), &QueryAppMembersBiz{
		Role: "EXPERIENCER",
	}, AppAuthToken(""))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}
