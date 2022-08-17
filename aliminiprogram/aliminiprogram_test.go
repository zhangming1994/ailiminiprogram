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
	aliMiniProgram, err := NewAliMiniProgram("2021003140601108",
		"MIIEogIBAAKCAQEAjSqrvIkA4nat+pCu+GYl2O6seFSxk58aIWyvVMOk6GSZJH54oNDCUF+hiXj3mWN0gTcrRAwIuJKzOQqEfQeSx21rH8sGe+kjmBZ7DzS+sIyBN6nL4fGf3p9A/7Jh/F+S/w0O9rXEoz5f/aejRuUPZnhudAg1DaYvgdD8GhBrQ2BPi0yG0HD/rHvsHcYE5rkvE6GO+ugReMD25jCsse4fsjokPJUXT//N5kiFQ6QRjeu3mp++SPEwuWw4uY3vJMRhORnd7rgpee3j8F7LcnZMJijtbUyYlI6WZIZrCzm0eaJFSIPKfy8KDSPiRl8IKY0nkgTuiYdCBiSSg1vPuGz+JQIDAQABAoIBADK4YSoFY576mzoK4AL52KFf5/H4JDFxGddmkHx5KvggpPF6y2akexybr5OQNg4TsNl4dAAn/r4qX2CgScgOf4tN4g5zSOoecx6C/UYSYm1WnlrclBejVBiNybHVauVdhZhEHdtApd0tXRnrMHwZcvYTUUQ5aq5zxl/zMZE/0vpKyLv/j04aXPYysL/TIBdtpzXf0KucrhXcMpejMZ2gRwk74Q5yz5x6Z/8ohpPuqxFsQkCUfwapFJXdTdp7hOsERruPM9k7iIvlZZ+DZOe+ThESxMbWtL4aE6v8rge3zXFvwiNw+6SRv4AaQ8G6TvR7oPf+g3BvGFX+F6s6X5GBFEUCgYEAxF/kkDDeh1YWDJ349MkmZsfN5E3NRZdSfhcd2jD5iCVKow9vPzXv1RGargHZfoKr8pCPEb5Bl/hP52Gj5CCUqs3YL4cwWUHr7uTXRDl+YuqiCHZWxUjLCklZzei468bGKs8Ur/XV0D0m2uL7RDczRwyqFZaFu4WuwZQb4CISjAsCgYEAuAd+PI4M4DOGVkebDaF2fKISt7fBW1cHjchTDPEngjbf7tOWNxYSfTYg/NuamMiGXUIXM47OIhPbjpvWb8r0aMNRxzADu74MLpdBKTCxT6nHx4Rdjvfjqk65kUTs3AmBpKdYN6rJQn1ltzVQYmJ/vSXfcm2UnQySf/UWZYxozI8CgYBqdv7I+eW7g/2iS4cs111dfdfvfgrbPuY3fDrwD3tJx8YGT5jlsU1cvBWbX19WSmwEr0ERwo93X+WaMYKUbOGNFNqvMeERkz7hnunikDUoVcMRSW9TuFp0Mj5g91pCYdsQXIL/vw1zv6OwKj1Fx5OiYohX+ep5N0n+o1jVlENv7wKBgEmnhhjKadDaDmw94IsGEpwzafZp71OVc4qFoOfPJRbyy5GU0cxeYywGHBHX+vonV8+/gfW6tDnjZL6hFKRNwKCle/eS+guMNuf21becq3rM4w6xpZhxEoe5VoAT7BDtJRw3dhFho4efFLTM+81EywdzPwlXXmG1BHV+LGacWd11AoGAcOj80vCWcLeBqz/krUQI4QmAxzlP0gcVScq2qdfYAdHFlJ67wvF/cC7anN5KeAoFa8rTIAbm+hSGCa7lxRo0v6nzGdB/Una3Yj+qLK6lh1ZggnYzs008pij8tRl3R5V6/yJyxKs5mF+FGO4xYm102GB57PyimGcTc8eWTE29Vws=",
		"MIIBCgKCAQEAoOCv952JEy+TopAUPMGYqZvZlCJV8k4mxiIhuoJf/QhhGFjXJYpjLMWXcAPprdUzXLxo4IBpbchkkh+1OMPthns37aTneThBUVwli1qliiqu9E4D8NjnvBLHbpDUiC61mqxWHWEJjEgGzx8lqP0mYQnYZorKtQZVRRIlu/PjyZeBgoLXTmHync9IJOepTpKp1vO6Y0oS72zJIAqgKIDe1y3oZ042iBMirfXb6SFFq0kdbZrSIwWVCY+Texq9CGrdgHlOER4mJ8vP38Hz0TAaH3EfTXERYMBa6yJPceLU0bxVc8FhArHgqTZe1ZSpn1nsAfnrAd/9LCExG5zXlvvF0wIDAQAB",
		"687b59193f3f462dd5336e5abf83c5d8_02941eef3187dddf3d3b83462e1dfcf6",
		"cb10062425d256fe704f7cfd8fc5f69b")
	if err != nil {
		log.Panic(err.Error())
	}
	return aliMiniProgram.NewClient()
}

func TestUploadVersion(t *testing.T) {
	client := initClient()
	var oauth UploadVersionBiz
	oauth.AppVersion = "0.0.3"
	oauth.TemplateID = "2021003141661773"
	oauth.TemplateVersion = "0.0.5"
	err := client.Mini.UploadVersion(context.Background(), &oauth, AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetAccessToken(t *testing.T) {
	client := initClient()
	var oauth OAuthToken
	oauth.GrantType = "authorization_code"
	oauth.Code = "f7508d8ec8244563b9368315e5d3SX13"
	resp, err := client.OAuth.UserOAuthToken(context.Background(), oauth)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestBaseInfoQuery(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryBaseInfo(context.Background(), AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestVersionGet(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryVersionList(context.Background(), QueryVersionListBiz{VersionStatus: "INIT"}, AppAuthToken("202208BBac076a8bdc5444238f836b901e15dX12"))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}

func TestBuildVersionStatus(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryVersionBuild(context.Background(), &QueryVersionBuildBiz{AppVersion: "0.0.2"}, AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
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
	}, AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
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
	}, AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCancelAudit(t *testing.T) {
	client := initClient()
	err := client.Mini.CancelVersionAudit(context.Background(), &CancelVersionAuditBiz{
		AppVersion: "0.0.1",
	}, AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateExperience(t *testing.T) {
	client := initClient()
	err := client.Mini.CreateExperience(context.Background(), &CreateExperienceBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	}, AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetExperienceStatus(t *testing.T) {
	client := initClient()
	resp, err := client.Mini.QueryExperience(context.Background(), &QueryExperienceBiz{
		AppVersion: "0.0.2",
		BundleID:   "com.alipay.alipaywallet",
	}, AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
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
	}, AppAuthToken("202208BB5ec488e862a8499ea130fe6cf3cb6D71"))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resp)
	}
}
