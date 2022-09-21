package aliminiprogram

type CreateProgramReq struct {
	AliPayAccount     string `json:"aliPayAccount"`     // 企业支付宝账号
	LegalPersonalName string `json:"legalPersonalName"` // 商家法人名称
	CertName          string `json:"certName"`          // 企业营业执照名称
	CertNO            string `json:"certNO"`            // 企业营业执照编码
	AppName           string `json:"appName"`           // 小程序名称
	ContactPhone      string `json:"contactPhone"`      // 商家联系人电话号码
	ContactName       string `json:"contactName"`       // 商家联系人名称
	Cid               string `json:"cid"`               // 运营商id
}

type MiniCategoryList struct {
	CategoryId         string `json:"categoryId"`
	CategoryName       string `json:"categoryName"`
	ParentCategoryId   string `json:"parentCategoryId"`
	HasChild           bool   `json:"hasChild"`
	NeedLicense        bool   `json:"needLicense"`
	NeedOutDoorPic     bool   `json:"needOutDoorPic"`
	NeedSpecialLicense bool   `json:"needSpecialLicense"`
}

type ModifyBaseInfo struct {
	Cid             string `json:"cid"`             // 运营商id
	AppName         string `json:"appName"`         // 小程序应用名称
	AppEnglishName  string `json:"appEnglishName"`  // 小程序应用英文名称
	AppSlogan       string `json:"appSlogan"`       // 小程序应用简介，一句话描述小程序功能
	AppDesc         string `json:"appDesc"`         // 小程序应用描述，20-200个字
	ServicePhone    string `json:"servicePhone"`    // 小程序客服电话
	ServiceEmail    string `json:"serviceEmail"`    // 小程序客服邮箱
	MiniCategoryIDs string `json:"miniCategoryIds"` // 新小程序前台类目，一级与二级、三级用下划线隔开，最多可以选四个类目，类目之间;隔开。使用后不再读取app_category_ids值，老前台类目将废弃
}

type FieldList struct {
	ApiName     string `json:"apiName"`
	FieldName   string `json:"fieldName"`
	PackageCode string `json:"packageCode"`
}

type AuthFieldScene struct {
	SceneCode string `json:"sceneCode"`
	SceneDesc string `json:"sceneDesc"`
}
