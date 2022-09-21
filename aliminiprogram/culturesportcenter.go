package aliminiprogram

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"net/http"
	"strings"
)

const ( // 平台自己的定义
	BasketBall    = iota + 1 // 篮球
	FootBall                 // 足球
	Badminton                // 羽毛球
	Tennis                   // 网球
	Natatorium               // 游泳馆
	TableTennis              // 乒乓
	VolleyBall               // 排球
	Trampoline               // 蹦床
	Archery                  // 射箭
	MultiFunction            // 多功能
	Extension                // 扩展
	Ski                      // 滑雪
	Gym                      // 健身房
)

// CultureSportCenter 文体中心
type CultureSportCenter Service

func VenueTypeProcess(platformVenueType []int) (alipayVenueType []string, err error) {
	if len(platformVenueType) == 0 {
		err = fmt.Errorf("场馆类型未选择")
		return alipayVenueType, err
	}
	for _, value := range platformVenueType {
		switch value {
		case BasketBall:
			alipayVenueType = append(alipayVenueType, "02")
		case FootBall:
			alipayVenueType = append(alipayVenueType, "01")
		case Badminton:
			alipayVenueType = append(alipayVenueType, "04")
		case Tennis:
			alipayVenueType = append(alipayVenueType, "09")
		case Natatorium:
			alipayVenueType = append(alipayVenueType, "08")
		case TableTennis:
			alipayVenueType = append(alipayVenueType, "03")
		case VolleyBall:
			alipayVenueType = append(alipayVenueType, "24")
		case Trampoline:
			alipayVenueType = append(alipayVenueType, "00")
		case Archery:
			alipayVenueType = append(alipayVenueType, "06")
		case MultiFunction:
			alipayVenueType = append(alipayVenueType, "00")
		case Extension:
			alipayVenueType = append(alipayVenueType, "00")
		case Ski:
			alipayVenueType = append(alipayVenueType, "21")
		case Gym:
			alipayVenueType = append(alipayVenueType, "22")
		}
	}
	return alipayVenueType, nil
}

// ImageProcess 图片压缩处理
func ImageProcess(path string) (string, error) {
	imageFile, err := http.Get(path)
	if err != nil {
		return "", nil
	}
	defer imageFile.Body.Close()
	img, err := jpeg.Decode(imageFile.Body)
	if err != nil && !strings.Contains(err.Error(), io.ErrUnexpectedEOF.Error()) {
		return "", err
	}
	set := resize.Resize(uint(img.Bounds().Max.X-img.Bounds().Max.X/10), uint(img.Bounds().Max.Y-img.Bounds().Max.Y/10), img, resize.Lanczos3)
	var buffer bytes.Buffer
	if err := jpeg.Encode(&buffer, set, &jpeg.Options{Quality: 30}); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// SportStoreEntryNew 新场馆入驻
func (c *CultureSportCenter) SportStoreEntryNew(ctx context.Context, biz *SportStoreEntryNewReq, opts ...ValueOptions) (SportStoreEntryNewResp, error) {
	apiMethod := "alipay.commerce.sports.venue.simple.create"
	req, err := c.Client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return SportStoreEntryNewResp{}, err
	}
	sportStoreEntryNewResp := new(SportStoreEntryNewResp)
	_, err = c.Client.Do(ctx, req, sportStoreEntryNewResp)
	if err != nil {
		return SportStoreEntryNewResp{}, err
	}
	return *sportStoreEntryNewResp, nil
}

// SportStoreInfoQuery 场馆信息查询
func (c *CultureSportCenter) SportStoreInfoQuery(ctx context.Context, biz SportStoreInfoQueryReq, opts ...ValueOptions) (SportStoreInfoResp, error) {
	method := "alipay.commerce.sports.venue.query"
	req, err := c.Client.NewRequest(method, biz, opts...)
	if err != nil {
		return SportStoreInfoResp{}, err
	}
	sportStoreEntryNewResp := new(SportStoreInfoResp)
	_, err = c.Client.Do(ctx, req, sportStoreEntryNewResp)
	if err != nil {
		return SportStoreInfoResp{}, err
	}
	return *sportStoreEntryNewResp, nil
}

// SportStoreModifyNew 场馆信息修改
func (c *CultureSportCenter) SportStoreModifyNew(ctx context.Context, biz SportStoreModifyReq, opts ...ValueOptions) (SportStoreModifyResp, error) {
	method := "alipay.commerce.sports.venue.simple.modify"
	param := make(map[string]interface{}, 0)
	if biz.OutVenueId == "" || biz.VenueId == "" {
		return SportStoreModifyResp{}, fmt.Errorf("param error:venueId or outVenueId")
	}
	param["out_venue_id"] = biz.OutVenueId
	param["venue_id"] = biz.VenueId
	if biz.JoinType != "" {
		param["join_type"] = biz.JoinType
	}
	if biz.Status != "" {
		param["status"] = biz.Status
	}
	if biz.VenuePid != "" {
		param["venue_pid"] = biz.VenuePid
	}
	if len(biz.VenueType) != 0 {
		param["venue_type"] = biz.VenueType
	}
	if biz.Name != "" {
		param["name"] = biz.Name
	}
	if biz.Desc != "" {
		param["desc"] = biz.Desc
	}
	if biz.Poster != "" {
		param["poster"] = biz.Poster
	}
	if len(biz.PictureList) != 0 {
		param["picture_list"] = biz.PictureList
	}
	if len(biz.ProductTypeList) != 0 {
		param["product_type_list"] = biz.ProductTypeList
	}
	if biz.OpeningHours != "" {
		param["opening_hours"] = biz.OpeningHours
	}
	if len(biz.Phone) != 0 {
		param["phone"] = biz.Phone
	}
	if biz.ProvinceCode != "" {
		param["province_code"] = biz.ProvinceCode
	}
	if biz.CityCode != "" {
		param["city_code"] = biz.CityCode
	}
	if biz.AreaCode != "" {
		param["area_code"] = biz.AreaCode
	}
	if biz.Longitude != "" {
		param["longitude"] = biz.Longitude
	}
	if biz.Latitude != "" {
		param["latitude"] = biz.Latitude
	}
	if biz.Address != "" {
		param["address"] = biz.Address
	}
	if biz.Traffic != "" {
		param["traffic"] = biz.Traffic
	}
	if biz.Poi != "" {
		param["poi"] = biz.Poi
	}
	if len(biz.TagList) != 0 {
		param["tag_list"] = biz.TagList
	}
	if biz.Bookable != "" {
		param["bookable"] = biz.Bookable
	}
	if biz.ExtraServiceUrl != "" {
		param["extra_service_url"] = biz.ExtraServiceUrl
	}
	if biz.PayeeAccount != "" {
		param["payee_account"] = biz.PayeeAccount
	}
	if biz.PaymentType != "" {
		param["payment_type"] = biz.PaymentType
	}
	if biz.PaymentMethod != "" {
		param["payment_method"] = biz.PaymentMethod
	}
	if biz.SubVenuePid != "" {
		param["sub_venue_pid"] = biz.SubVenuePid
	}
	if biz.SubVenueSmid != "" {
		param["sub_venue_smid"] = biz.SubVenueSmid
	}
	if biz.AdmissionRequirement != "" {
		param["admission_requirement"] = biz.AdmissionRequirement
	}
	if biz.Announcement != "" {
		param["announcement"] = biz.Announcement
	}
	if biz.Promotion != "" {
		param["promotion"] = biz.Promotion
	}
	if biz.EquipmentRental != "" {
		param["equipment_rental"] = biz.EquipmentRental
	}
	if biz.Vip != "" {
		param["vip"] = biz.Vip
	}
	if biz.RecPrice != "" {
		param["rec_price"] = biz.RecPrice
	}
	if biz.Training != "" {
		param["training"] = biz.Training
	}
	if len(biz.FacilityList) != 0 {
		param["facility_list"] = biz.FacilityList
	}
	req, err := c.Client.NewRequest(method, param, opts...)
	if err != nil {
		return SportStoreModifyResp{}, err
	}
	sportStoreEntryNewResp := new(SportStoreModifyResp)
	_, err = c.Client.Do(ctx, req, sportStoreEntryNewResp)
	if err != nil {
		return SportStoreModifyResp{}, err
	}
	return *sportStoreEntryNewResp, nil
}

// SportVenueOrderSync 文体中心订单数据同步 由服务商将用户场馆订单数据同步到支付宝域内，经用户授权查询和展示用户订单数据。
func (c *CultureSportCenter) SportVenueOrderSync(ctx context.Context, biz VenueOrder, opts ...ValueOptions) (string, error) {
	method := "alipay.commerce.sports.venue.order.sync"
	req, err := c.Client.NewRequest(method, biz, opts...)
	if err != nil {
		return "", err
	}
	venueOrder := new(VenueOrderResp)
	_, err = c.Client.Do(ctx, req, venueOrder)
	if err != nil {
		return "", err
	}
	return venueOrder.AlipayCommerceSportsVenueOrderSyncResponse.OrderId, nil
}

// SportStoreOrderQuery 订单信息查询
func (c *CultureSportCenter) SportStoreOrderQuery(ctx context.Context, biz SportStoreOrderQueryReq, opts ...ValueOptions) (SportStoreOrderQueryResp, error) {
	method := "alipay.commerce.sports.venue.order.query"
	req, err := c.Client.NewRequest(method, biz, opts...)
	if err != nil {
		return SportStoreOrderQueryResp{}, err
	}
	sportStoreEntryNewResp := new(SportStoreOrderQueryResp)
	_, err = c.Client.Do(ctx, req, sportStoreEntryNewResp)
	if err != nil {
		return SportStoreOrderQueryResp{}, err
	}
	return *sportStoreEntryNewResp, nil
}

// SportStoreRefund 服务商发起退款
func (c *CultureSportCenter) SportStoreRefund(refund RefundReq) (RefundResp, error) {
	return RefundResp{}, nil
}

// SportStoreOrderVerify 订单核销
func (c *CultureSportCenter) SportStoreOrderVerify(verify OrderVerifyReq) error {
	return nil
}

// SportStoreOrderVerifyCertificate 订单核销凭证查询
func (c *CultureSportCenter) SportStoreOrderVerifyCertificate(verifyQuery OrderVerifyVoucherQueryReq) (OrderVerifyVoucherQueryResp, error) {
	return OrderVerifyVoucherQueryResp{}, nil
}

// SportStoreProductAndQuantityQuery 场馆产品和库存查询
func (c *CultureSportCenter) SportStoreProductAndQuantityQuery(productQuery SportStoreProductNumQueryReq) (SportStoreProductNumQueryResp, error) {
	return SportStoreProductNumQueryResp{}, nil
}

// SportStoreOrderConfirm 订单确认 异步接口
func (c *CultureSportCenter) SportStoreOrderConfirm(orderConfirm SportStoreOrderConfirmReq) (SportStoreOrderConfirmResp, error) {
	return SportStoreOrderConfirmResp{}, nil
}

// SportStoreAuditNotify 场馆审核通知
func (c *CultureSportCenter) SportStoreAuditNotify(storeAudit AliMiniNotifySportStoreAuditProperty) error {
	return nil
}

// SportStoreOrderCreate 订单创建
func (c *CultureSportCenter) SportStoreOrderCreate(orderParam OrderCreateBodyField) (OrderCreateResp, error) {
	return OrderCreateResp{}, nil
}

// SportStoreOrderStatus 订单状态同步
func (c *CultureSportCenter) SportStoreOrderStatus(orderStatus AliMiniNotifySportStoreOrderStatusProperty) error {
	return nil
}
