package aliminiprogram

import "time"

// SportStoreEntryModifyCommonField 场馆入驻和场馆修改公共请求参数
type SportStoreEntryModifyCommonField struct {
	FacilityList         []int    `json:"facility_list"`                     // 场馆设施
	TagList              []string `json:"tag_list"`                          // 标签列表
	PictureList          []string `json:"picture_list"`                      // 场馆图片编码列表 要求参见主图要求
	ProductTypeList      []string `json:"product_type_list"`                 // 场馆售卖的产品集合 中心化必填 半中心化不用写
	Phone                []string `json:"phone" validate:"required"`         // 联系电话
	VenueType            []string `json:"venue_type" validate:"required"`    // 场馆类型
	OutVenueId           string   `json:"out_venue_id" validate:"required"`  // 主场馆在服务商的ID
	JoinType             string   `json:"join_type"`                         // 接入方式
	VenueId              string   `json:"venue_id"`                          // 支付宝主场馆ID，不可变更
	VenuePid             string   `json:"venue_pid"`                         // 场馆商户PId
	Name                 string   `json:"name" validate:"required"`          // 场馆名称
	Desc                 string   `json:"desc"`                              // 场馆描述
	Poster               string   `json:"poster" validate:"required"`        // 场馆主图海报图片的base64编码 图片格式必须是jpg 图片大小不超过125KB 非data uri格式
	OpeningHours         string   `json:"opening_hours" validate:"required"` // 营业开始时间-结束时间
	ProvinceCode         string   `json:"province_code" validate:"required"` // 省份code
	CityCode             string   `json:"city_code" validate:"required"`     // 城市code
	AreaCode             string   `json:"area_code" validate:"required"`     // 区域code
	Longitude            string   `json:"longitude" validate:"required"`     // 经度
	Latitude             string   `json:"latitude" validate:"required"`      // 纬度
	Address              string   `json:"address" validate:"required"`       // 地址
	Traffic              string   `json:"traffic" validate:"required"`       // 交通信息
	Poi                  string   `json:"poi"`                               // POI
	Bookable             string   `json:"bookable"`                          // 场馆是否可预定Y/N，不传默认可预定
	ExtraServiceUrl      string   `json:"extra_service_url"`                 // 场馆更多服务链接：可从文体场馆页跳转进此链接，进入服务商的该场馆页面
	PayeeAccount         string   `json:"payee_account"`                     // （半中心化场馆）不填;（中心化场馆）收款方支付宝账户，当payment_method为空或account时必传
	PaymentType          string   `json:"payment_type"`                      // 半中心化场馆）不填; （中心化场馆）收款类型 （indirect=间连/direct=直连） 直连：收款方为商户/场馆 间连：收款方为服务商
	PaymentMethod        string   `json:"payment_method"`                    // 收款方式 空值/account：通过支付宝账号收款； smid：通过smid收款
	SubVenuePid          string   `json:"sub_venue_pid"`                     // 子场馆pid（payment_method为smid时必传）
	SubVenueSmid         string   `json:"sub_venue_smid"`                    // 子场馆商户二级smid（payment_method为smid时必传）
	AdmissionRequirement string   `json:"admission_requirement"`             // 入场要求
	Announcement         string   `json:"announcement"`                      // 公告
	Promotion            string   `json:"promotion"`                         // 促销信息
	EquipmentRental      string   `json:"equipment_rental"`                  // 器材租赁信息
	Vip                  string   `json:"vip"`                               // 会员卡信息
	RecPrice             string   `json:"rec_price"`                         // 场馆维度的推荐价格
	Training             string   `json:"training"`                          // 培训信息
}

// SubVenueList 子场馆列表
type SubVenueList struct {
	OutSubVenueId        string   `json:"out_sub_venue_id"`      // 服务商场馆id
	SubVenueId           string   `json:"sub_venue_id"`          // 支付宝子场馆id
	SubVenuePid          string   `json:"sub_venue_pid"`         // 子场馆Pid
	PayeeAccount         string   `json:"payee_account"`         // 收款方支付宝账户
	PaymentType          string   `json:"payment_type"`          // 收款方式
	PaymentMethod        string   `json:"payment_method"`        // 收款方式
	SubVenueSmid         string   `json:"sub_venue_smid"`        // 二级smid
	VenueType            string   `json:"venue_type"`            // 场馆类型
	Name                 string   `json:"name"`                  // 场馆名称
	Desc                 string   `json:"desc"`                  // 介绍
	Poster               string   `json:"poster"`                // 主图
	PictureList          []string `json:"picture_list"`          // 图片链接
	ProductTypeList      []string `json:"product_type_list"`     // 售卖铲平集合
	OpeningHours         string   `json:"opening_hours"`         // 营业时间
	Phone                []string `json:"phone"`                 // 联系电话
	AdmissionRequirement string   `json:"admission_requirement"` // 入场要求
	FacilityList         []int    `json:"facility_list"`         // 设施
	Announcement         string   `json:"announcement"`          // 公告
	TagList              []string `json:"tag_list"`              // 标签列表
	Promotion            string   `json:"promotion"`             // 促销信息
	EquipmentRental      string   `json:"equipment_rental"`      // 器材祖灵
	Vip                  string   `json:"vip"`                   // vip
	Training             string   `json:"training"`              // 培训信息
	Bookable             string   `json:"bookable"`              // 是否可以预定
	SubVenueStatus       string   `json:"sub_venue_status"`      // 状态
}

// SportStoreEntryNewReq 场馆入驻请求字段
type SportStoreEntryNewReq struct {
	SportStoreEntryModifyCommonField
	TestVenue string `json:"test_venue"` // 是否为“测试场馆”。如果上传的场馆为想要进行测试的非正式场馆，则填写“Y”。如上传正式场馆，则不传或填写为N
}

// SportStoreEntryNewResp 场馆入驻返回信息
type SportStoreEntryNewResp struct {
	AlipayCommerceSportsVenueSimpleCreateResponse struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubVenueList []struct {
			SubVenueId     string `json:"sub_venue_id"`     // 支付宝对应的子场馆id
			OutSubVenueId  string `json:"out_sub_venue_id"` // 服务商对应子场馆id
			SubVenueStatus string `json:"sub_venue_status"` // 状态
		} `json:"sub_venue_list"` // 子场馆信息
		VenueId     string `json:"venue_id"`     // 场馆在支付宝的唯一id
		OutVenueId  string `json:"out_venue_id"` // 服务商对应的场馆id
		VenueStatus string `json:"venue_status"` // 场馆当前状态 安全审核中：infosec-audit 安全审核不通过：infosec-unpass 云验收中： cloud-audit 云验收不通过： cloud-unpass 上架： online 下架： offline 人工下架： manual-offline
		Desc        string `json:"desc"`         // 描述
	} `json:"alipay_commerce_sports_venue_simple_create_response"`
	Sign string `json:"sign"`
}

// SportStoreInfoQueryReq 场馆信息查询请求参数
type SportStoreInfoQueryReq struct {
	OutVenueId string `json:"out_venue_id"`
	VenueId    string `json:"venue_id"`
}

// SportStoreModifyReq 场馆信息修改
type SportStoreModifyReq struct {
	SportStoreEntryModifyCommonField
	Status string `json:"status"`
}

// SportStoreInfoResp 场馆信息查询
type SportStoreInfoResp struct {
	AlipayCommerceSportsVenueQueryResponse struct {
		Code            string         `json:"code"`
		Msg             string         `json:"msg"`
		VenueId         string         `json:"venue_id"`                          // 上官商户PID
		OutVenueId      string         `json:"out_venue_id" validate:"required"`  // 主场馆在服务商的ID
		JoinType        string         `json:"join_type"`                         // 接入方式
		VenueType       []string       `json:"venue_type" validate:"required"`    // 场馆类型
		Name            string         `json:"name" validate:"required"`          // 场馆名称
		Desc            string         `json:"desc"`                              // 场馆描述
		Poster          string         `json:"poster" validate:"required"`        // 场馆主图海报图片的base64编码 图片格式必须是jpg 图片大小不超过125KB 非data uri格式
		PictureList     []string       `json:"picture_list"`                      // 场馆图片编码列表 要求参见主图要求
		ProductTypeList []string       `json:"product_type_list"`                 // 场馆售卖的产品集合 中心化必填 半中心化不用写
		OpeningHours    string         `json:"opening_hours" validate:"required"` // 营业开始时间-结束时间
		Phone           []string       `json:"phone" validate:"required"`         // 联系电话
		ProvinceCode    string         `json:"province_code" validate:"required"` // 省份code
		CityCode        string         `json:"city_code" validate:"required"`     // 城市code
		AreaCode        string         `json:"area_code" validate:"required"`     // 区域code
		Longitude       string         `json:"longitude" validate:"required"`     // 经度
		Latitude        string         `json:"latitude" validate:"required"`      // 纬度
		Address         string         `json:"address" validate:"required"`       // 地址
		Traffic         string         `json:"traffic" validate:"required"`       // 交通信息
		Poi             string         `json:"poi"`                               // POI
		TagList         []string       `json:"tag_list"`                          // 标签列表
		VenueStatus     string         `json:"venue_status"`
		Bookable        string         `json:"bookable"`          // 场馆是否可预定Y/N，不传默认可预定
		ExtraServiceUrl string         `json:"extra_service_url"` // 场馆更多服务链接：可从文体场馆页跳转进此链接，进入服务商的该场馆页面
		SubVenueList    []SubVenueList `json:"sub_venue_list"`    // 子场馆列表
	} `json:"alipay_commerce_sports_venue_query_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// SportStoreModifyResp 场馆信息修改返回参数
type SportStoreModifyResp struct {
	AlipayCommerceSportsVenueSimpleModifyResponse struct {
		Code            string         `json:"code"`
		Msg             string         `json:"msg"`
		VenueId         string         `json:"venue_id"`          // 上官商户PID
		OutVenueId      string         `json:"out_venue_id"`      // 主场馆在服务商的ID
		JoinType        string         `json:"join_type"`         // 接入方式
		VenuePid        string         `json:"venue_pid"`         //
		VenueType       []string       `json:"venue_type"`        // 场馆类型
		Name            string         `json:"name"`              // 场馆名称
		Desc            string         `json:"desc"`              // 场馆描述
		Poster          string         `json:"poster"`            // 场馆主图海报图片的base64编码 图片格式必须是jpg 图片大小不超过125KB 非data uri格式
		PictureList     []string       `json:"picture_list"`      // 场馆图片编码列表 要求参见主图要求
		ProductTypeList string         `json:"product_type_list"` // 场馆售卖产品集合
		OpeningHours    string         `json:"opening_hours"`     // 营业开始时间-结束时间
		Phone           []string       `json:"phone"`             // 联系电话
		ProvinceCode    string         `json:"province_code"`     // 省份code
		CityCode        string         `json:"city_code"`         // 城市code
		AreaCode        string         `json:"area_code"`         // 区域code
		Longitude       string         `json:"longitude"`         // 经度
		Latitude        string         `json:"latitude"`          // 纬度
		Address         string         `json:"address"`           // 地址
		Traffic         string         `json:"traffic"`           // 交通信息
		Poi             string         `json:"poi"`               // POI
		TagList         []string       `json:"tag_list"`          // 标签列表
		VenueStatus     string         `json:"venue_status"`      // zhuangtai
		Bookable        string         `json:"bookable"`          // 场馆是否可预定Y/N，不传默认可预定
		ExtraServiceUrl string         `json:"extra_service_url"` // 场馆更多服务链接：可从文体场馆页跳转进此链接，进入服务商的该场馆页面
		SubVenueList    []SubVenueList `json:"sub_venue_list"`    // 子场馆列表
	} `json:"alipay_commerce_sports_venue_simple_modify_response"`
	Sign string `json:"sign"`
}

// SportStoreOrderQueryReq 订单信息查询请求参数
type SportStoreOrderQueryReq struct {
	OrderId    string `json:"order_id" validate:"required"`     // 支付宝业务订单唯一ID
	OutOrderId string `json:"out_order_id" validate:"required"` // 服务商内部唯一id
}

// SportStoreOrderQueryResp 订单信息查询返回参数
type SportStoreOrderQueryResp struct {
	ProductGroupList Products `json:"product_group_list"`
	OrderId          string   `json:"order_id"`
	CreateTime       string   `json:"create_time"`
	PaymentTime      string   `json:"payment_time"`
	RefundTime       string   `json:"refund_time"`
	RefundEndTime    string   `json:"refund_end_time"`
	OrderStatus      string   `json:"order_status"`
	Desc             string   `json:"desc"`
	PaymentAmount    float64  `json:"payment_amount"`
	TotalAmount      float64  `json:"total_amount"`
}

// Products 订单商品信息
type Products struct {
	Product    ProductInfo `json:"product"`
	TotalPrice string      `json:"total_price"`
	Count      string      `json:"count"`
}

// ProductInfo 铲平对象
type ProductInfo struct {
	CanBuyRule struct {
		CanBuyMinCount         int    `json:"can_buy_min_count"`
		CanBuyMaxCount         int    `json:"can_buy_max_count"`
		CategoryChooseMinCount int    `json:"category_choose_min_count"`
		CategoryChooseMaxCount int    `json:"category_choose_max_count"`
		CanBuyLimitType        string `json:"can_buy_limit_type"`
	} `json:"can_buy_rule"` // 购买限制规则信息
	CategoryList      []CategoryList `json:"category_list"`
	CanRefundMinute   int            `json:"can_refund_minute"`
	StockCount        int64          `json:"stock_count"`
	OutProductId      string         `json:"out_product_id"`
	Name              string         `json:"name"`
	Desc              string         `json:"desc"`
	Notice            string         `json:"notice"`
	Poster            string         `json:"poster"`
	ProductType       string         `json:"product_type"`
	StartTime         string         `json:"start_time"`
	EndTime           string         `json:"end_time"`
	SalePrice         string         `json:"sale_price"`
	OriginPrice       string         `json:"origin_price"`
	RefundDesc        string         `json:"refund_desc"`
	VoucherType       string         `json:"voucher_type"`
	VoucherVerifyType string         `json:"voucher_verify_type"`
}

// CategoryList 产品分类列表
type CategoryList struct {
	ZoneList []struct {
		StockCount  int    `json:"stock_count"`
		OutZoneId   string `json:"out_zone_id"`
		StartTime   string `json:"start_time"`
		EndTime     string `json:"end_time"`
		SalePrice   string `json:"sale_price"`
		OriginPrice string `json:"origin_price"`
	} `json:"zone_list"`
	OutCategoryId string `json:"out_category_id"`
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Date          string `json:"date"`
}

// RefundReq 服务商发起退款
type RefundReq struct {
	OrderId      string `json:"order_id"`      // 支付宝业务订单唯一id
	OutOrderId   string `json:"out_order_id"`  // 服务商内部唯一订单号
	RefundAmount string `json:"refund_amount"` // 退款金额
	Desc         string `json:"desc"`          // 退款原因
}

// RefundResp 退款返回参数
type RefundResp struct {
	OrderId      string `json:"order_id"`      // 支付宝业务唯一订单号
	OutOrderId   string `json:"out_order_id"`  // 服务商内部唯一订单号
	TradeNo      string `json:"trade_no"`      // 支付宝退款订单号
	RefundStatus string `json:"refund_status"` // 退款状态
	Desc         string `json:"desc"`          // 操作描述
}

// OrderVerifyReq 订单核销请求参数
type OrderVerifyReq struct {
	OrderId      string `json:"order_id"`
	OutOrderId   string `json:"out_order_id"`
	OutVoucherId string `json:"out_voucher_id"`
	VerifyStatus string `json:"verify_status"`
	VerifyCount  string `json:"verify_count"`
	Desc         string `json:"desc"`
}

//OrderVerifyVoucherQueryReq 核销凭证查询
type OrderVerifyVoucherQueryReq struct {
	OrderId      string `json:"order_id"`       // 支付宝业务订单唯一id
	OutOrderId   string `json:"out_order_id"`   // 服务商内部唯一id
	OutVoucherId string `json:"out_voucher_id"` // 外部凭证id
}

// OrderVerifyVoucherQueryResp 核销凭证查询返回参数
type OrderVerifyVoucherQueryResp struct {
	OutVoucherId        string `json:"out_voucher_id"`        // 外部凭证id
	VoucherStatus       string `json:"voucher_status"`        // 核销状态
	VerifyCount         string `json:"verify_count"`          // 凭证剩余核销次数
	VoucherType         string `json:"voucher_type"`          // 凭证类型
	VouchContent        string `json:"vouch_content"`         // 凭证内容
	VoucherDesc         string `json:"voucher_desc"`          // 凭证文案描述
	VoucherValidityTime string `json:"voucher_validity_time"` // 凭证生效时间
	VoucherExpireTime   string `json:"voucher_expire_time"`   // 凭证过期时间
	VoucherDisposeRule  string `json:"voucher_dispose_rule"`  // 凭证核销规则
}

// SportStoreProductNumQueryReq 查询场馆的产品和库存请求参数
type SportStoreProductNumQueryReq struct {
	VenueId       string `json:"venue_id"`         // 支付宝主场馆id
	SubVenueId    string `json:"sub_venue_id"`     // 支付宝子场馆id
	OutVenueId    string `json:"out_venue_id"`     // 服务商主场馆id
	OutSubVenueId string `json:"out_sub_venue_id"` // 服务商子场馆id
	ProductType   string `json:"product_type"`     // 产品类型
	StartDate     string `json:"start_date"`       // 产品剋用范围开始日期
	EndDate       string `json:"end_date"`         // 产品可用的范围的结束日期
	DataCount     string `json:"data_count"`       // 数量要求
}

//SportStoreProductNumQueryResp 查询场馆的产品和库存返回参数
type SportStoreProductNumQueryResp struct {
	CanBuyRule struct {
		CanBuyMinCount         int      `json:"can_buy_min_count"`         // 一笔订单下产品最小购买数量，不传默认1
		CanBuyMaxCount         int      `json:"can_buy_max_count"`         // 一笔订单下产品最大购买数量，不传默认无限制
		CategoryChooseMinCount int      `json:"category_choose_min_count"` // 场馆最小选择数，默认1，价格日历产品下代表一笔订单至少选择可跨几个分类(场地)，其他产品类型暂无意义
		CategoryChooseMaxCount int      `json:"category_choose_max_count"` // 最大选择数，默认无，价格日历产品下代表一笔订单最大可多选择N个分类(场地)，其他产品类型暂无意义
		CanBuyLimitType        string   `json:"can_buy_limit_type"`        // 购买限制的类型，0-无需关联，1-需要关联，默认无需关联。 名词解释：票券课程类型商品的无需在意本字段，在价格日历预订下，代表是否需要时间段或场地的连场限制。
		DailyStartTime         string   `json:"daily_start_time"`          // 已废弃，无效）可购买开始时间，当前时间小于这个时间用户不能下单。默认00:00:00
		DailyEndTime           string   `json:"daily_end_time"`            // （已废弃，无效）可购买截止时间，当前时间大于这个时间用户不能下单。 如果全天可下单，可购买开始和截止时间可以为空或者为00:00:00 - 23:59:59；如果全天不可下单，可购买开始时间(daily_start_time)和可购买截止时间(daily_end_time)均为00:00:00。
		BuyTimeLimit           []string `json:"buy_time_limit"`            // 购买时间限制，商品可下单的时间范围，为“可购买开始时间-可购买截止时间”时间段列表
	} `json:"can_buy_rule"` // 购买限制规则信息
	CategoryList []struct {
		ZoneList []struct {
			StockCount  int    `json:"stock_count"`  // 库存
			OutZoneId   string `json:"out_zone_id"`  // 区间唯一id
			ZoneName    string `json:"zone_name"`    // 上屏规格
			StartTime   string `json:"start_time"`   // 开始时间
			EndTime     string `json:"end_time"`     // 结束时间
			SalePrice   string `json:"sale_price"`   // 售卖价格
			OriginPrice string `json:"origin_price"` // 原价
		} `json:"zone_list"`
		OutCateGoryId string `json:"out_cate_gory_id"` // 类别唯一id
		Name          string `json:"name"`             // 类别名称
		Desc          string `json:"desc"`             // 描述
		Date          string `json:"date"`             // 使用日期
	} `json:"category_list"`
	UserNameRequired  bool   `json:"user_name_required"`  // 核销的时候是否需要用户姓名
	CanRefund         bool   `json:"can_refund"`          // 是否支持退款
	StockCount        int    `json:"stock_count"`         // 产品库存
	CanRefundMinute   int    `json:"can_refund_minute"`   // 使用前多长时间可以退款
	OutProductId      string `json:"out_product_id"`      // 服务商的产品id
	Name              string `json:"name"`                // 产品名称
	Desc              string `json:"desc"`                // 产品描述
	Notice            string `json:"notice"`              // 须知
	Poster            string `json:"poster"`              // 图片地址
	ProductType       string `json:"product_type"`        // 产品类型
	StartTime         string `json:"start_time"`          // 可用开始时间
	EndTime           string `json:"end_time"`            // 可用结束时间
	SalePrice         string `json:"sale_price"`          // 售卖价格
	OriginPrice       string `json:"origin_price"`        // 原价
	RefundDesc        string `json:"refund_desc"`         // 退款规则描述
	VoucherType       string `json:"voucher_type"`        // 凭证类型
	VoucherVerifyType string `json:"voucher_verify_type"` // 凭证核销方式
}

// SportStoreOrderConfirmReq 订单确认请求参数
type SportStoreOrderConfirmReq struct {
	ProductGroupList []struct {
		Product    SportStoreProductNumQueryResp `json:"product"`
		Count      string                        `json:"count"`
		TotalPrice string                        `json:"total_price"`
	} `json:"product_group_list"`
	TotalAmount   float64 `json:"total_amount"`
	OrderId       string  `json:"order_id"`
	OutOrderId    string  `json:"out_order_id"`
	ConfirmStatus string  `json:"confirm_status"`
	ConfirmDesc   string  `json:"confirm_desc"`
}

// SportStoreOrderConfirmResp 订单确认返回参数
type SportStoreOrderConfirmResp struct {
	OrderId     string `json:"orderId"`
	OrderStatus string `json:"order_status"`
}

// AliMiniNotifySportStoreAuditProperty  场馆审核蚂蚁消息属性
type AliMiniNotifySportStoreAuditProperty struct {
	VenueId       string `json:"venue_id"`
	OutVenueId    string `json:"out_venue_id"`
	OutSubVenueId string `json:"out_sub_venue_id"`
	/*
		场馆当前状态
		安全审核中：infosec-audit
		安全审核不通过：infosec-unpass
		云验收中： cloud-audit
		云验收不通过： cloud-unpass
		上架： online
		下架： offline
		人工下架： manual-offline
	*/
	VenueStatus string   `json:"venue_status"` // 状态
	Desc        string   `json:"desc"`         // 描述
	SubVenueId  []string `json:"sub_venue_id"` // 子场馆Id列表
}

// AliMiniNotifySportStoreOrderStatusProperty 订单状态消息属性
type AliMiniNotifySportStoreOrderStatusProperty struct {
	OrderId     string `json:"order_id"`     // 支付宝业务订单唯一id
	OutOrderId  string `json:"out_order_id"` // 服务商内部唯一id
	TotalAmount string `json:"total_amount"` // 订单总金额
	TradeNo     string `json:"trade_no"`     // 支付宝交易号
	OrderStatus string `json:"order_status"` // 支付结果 pay_succ- 已支付 refund_succ - 已退款 closed - 已关闭
	Desc        string `json:"desc"`         // 附加描述
}

// OrderCreateCommonField 订单创建公共请求参数
type OrderCreateCommonField struct {
	Method       string `json:"method"`
	Charset      string `json:"charset"`
	Version      string `json:"version"`
	BizAppId     string `json:"biz_app_id"`
	InvokeAppId  string `json:"invoke_app_id"`
	UtcTimestamp string `json:"utc_timestamp"`
	SignType     string `json:"sign_type"`
	Sign         string `json:"sign"`
}

// OrderCreateBodyField 订单创建body参数
type OrderCreateBodyField struct {
	ReceiveUserInfo struct {
		UserPhone string `json:"user_phone"` // 手机号
	} `json:"receive_user_info"` // 收货人信息
	ProductGroupList []struct {
		Product    SportStoreProductNumQueryResp `json:"product"`     // 产品对象
		Count      string                        `json:"count"`       // 数量
		TotalPrice string                        `json:"total_price"` // 总价
	} `json:"product_group_list"` // 订单商品信息列表
	TotalAmount   float64 `json:"total_amount"`     // 订单总金额
	OrderDesc     string  `json:"order_desc"`       // 订单描述
	OrderId       string  `json:"order_id"`         // 支付宝业务订单id
	TradeNo       string  `json:"trade_no"`         // 支付宝交易订单号
	OrderType     string  `json:"order_type"`       // 订单类型
	VenueId       string  `json:"venue_id"`         // 场馆id
	SubVenueId    string  `json:"sub_venue_id"`     // 子场馆id
	OutVenueId    string  `json:"out_venue_id"`     // 场馆外部id
	OutSubVenueId string  `json:"out_sub_venue_id"` // 子场馆外部id
	CreateTime    string  `json:"create_time"`      // 创建时间
	ProductType   string  `json:"product_type"`     // 产品类型
}

// OrderCreateResp  订单创建返回参数
type OrderCreateResp struct {
	OrderId        string `json:"order_id"`         // 支付宝业务订单唯一ID
	OutOrderId     string `json:"out_order_id"`     // 服务商内部唯一订单号
	OutOrderStatus string `json:"out_order_status"` // 创单操作结果，reverse_proc-预定中，reverse_succ-预定成功
	Desc           string `json:"desc"`             // 描述
}

// VenueOrder 订单回流请求参数
type VenueOrder struct {
	UserId           string                `json:"user_id"`            // 买家支付宝userid
	OutOrderId       string                `json:"out_order_id"`       // 服务商内部唯一id
	OrderType        string                `json:"order_type"`         // 订单类型
	OrderStatus      string                `json:"order_status"`       // 订单交易状态 pay_succ-已支付（若支持多次核销则在全部核销之前都是已支付状态）,refund_succ-已退款,verify_proc-使用中（已入场但是还未结束）,verify_succ-已使用,overdue-已过期（超过使用时间未使用且未退款）
	TotalAmount      float64               `json:"total_amount"`       // 订单总金额(元)
	VenueId          *string               `json:"venue_id"`           // 特殊可选字段 支付宝主场馆ID，场馆入驻时支付宝返回的主场馆ID。和out_venue_id之间至少存在一个 特殊可选字段
	SubVenueId       *string               `json:"sub_venue_id"`       // 可选字段 支付宝子场馆ID，场馆入驻时支付宝返回的子场馆ID。如果在场馆入驻时有子场馆则传入入驻时返回的sub_venue_id；如果场馆入驻时不存在子场馆，则无须传入。
	OutVenueId       string                `json:"out_venue_id"`       // 特殊可选字段 isv场馆id，与场馆入驻时一致，须保证系统内唯一。和venue_id之间至少存在一个
	OutSubVenueId    *string               `json:"out_sub_venue_id"`   // 可选字段 isv子场馆id，与场馆入驻时一致，须保证系统内唯一。如果在场馆入驻时有子场馆则传入入住时的out_sub_venue_id；如果场馆入驻时不存在子场馆，则无须传入
	CreateTime       string                `json:"create_time"`        // 订单创建时间 2022-04-05 12:12:12
	ProductGroupList []ProductSimpleInfo   `json:"product_group_list"` // 订单商品信息列表，目前仅支持1笔订单1条商品信息数据，即一笔订单只能包含1种商品。
	TradeInfoList    []VenueOrderTradeInfo `json:"trade_info_list"`    // 可选 订单的交易信息列表，传入支付、退款等操作的信息。第一次同步必传；第一次同步之后如果没有交易变化则调用时可以不传本参数。每一条交易数据同步后不支持修改。
}

type VenueOrderTradeInfo struct {
	Id              string  `json:"id"`                // 交易记录号 单笔订单内唯一
	TradeType       string  `json:"trade_type"`        // 交易类型，包括'pay'-支付、'refund'-退款
	TradeNo         string  `json:"trade_no"`          // 支付宝交易单号，本条记录对应的交易信息。如果是支付行为，则直接为交易单号；如果是基于原支付交易单原路返还退款，则为原支付交易单号；
	PartnerId       string  `json:"partner_id"`        // 交易所属pid，一般为发起交易的应用配置的pid。
	UserId          string  `json:"user_id"`           // 支付宝用户id
	Amount          float64 `json:"amount"`            // 金额(单位：元)，保留两位小数。支付时为订单金额、退款时为交易退款金额
	RefundRequestNo string  `json:"refund_request_no"` // 特殊可选 支付宝退款交易的请求号，标识一次退款请求，对应发起退款时的out_request_no。交易类型为退款时必传。
	OperationTime   string  `json:"operation_time"`    // 可选 Date类型, 2006-01-02 15:04:03 交易发起时间
	Desc            string  `json:"desc"`              // 可选 操作描述
}

type ProductSimpleInfo struct {
	ProductName  string     `json:"product_name"`  // 商品名称
	CategoryName string     `json:"category_name"` // 可选字段 商品规格名称。可以是场地名称（比如羽毛球场地1）。当产品类型为日历型或者日历型票券时必填。
	ZoneName     *string    `json:"zone_name"`     // 可选字段 产品规格信息 仅日历型票券商品需要配置该字段
	StartTime    *time.Time `json:"start_time"`    // 可选字段 商品使用开始时间
	EndTime      *time.Time `json:"end_time"`      // 特殊可选 商品使用结束时间（Date类型），结束时间非空时必须有开始时间传值，且开始时间必须早于结束时间
	ProductType  string     `json:"product_type"`  // 商品类型，'calendar'-日历型，'ticket'-票券，'course'-课程，'calendar_ticket'-日历型票券
	SalePrice    string     `json:"sale_price"`    // 售卖价格 元 保留两位小数
	Count        int64      `json:"count"`         // 可选 产品购买数量
}

type VenueOrderResp struct {
	AlipayCommerceSportsVenueOrderSyncResponse struct {
		OrderId string `json:"order_id"` // 支付宝业务订单唯一id
		Code    string `json:"code"`
		Msg     string `json:"msg"`
	} `json:"alipay_commerce_sports_venue_order_sync_response"`
	Sign string `json:"sign"`
}
