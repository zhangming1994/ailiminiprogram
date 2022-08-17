package aliminiprogram

import (
	"bytes"
	"context"
	"encoding/base64"
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

func VenueTypeProcess(platformVenueType []int) (alipayVenueType []string) {
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
	return
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
	if err := jpeg.Encode(&buffer, set, &jpeg.Options{Quality: 40}); err != nil {
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
func (c *CultureSportCenter) SportStoreInfoQuery(ctx context.Context, biz SportStoreInfoQueryReq, opts ...ValueOptions) (SportStoreInfoQueryResp, error) {
	method := "alipay.commerce.sports.venue.query"
	req, err := c.Client.NewRequest(method, biz, opts...)
	if err != nil {
		return SportStoreInfoQueryResp{}, err
	}
	sportStoreEntryNewResp := new(SportStoreInfoQueryResp)
	_, err = c.Client.Do(ctx, req, sportStoreEntryNewResp)
	if err != nil {
		return SportStoreInfoQueryResp{}, err
	}
	return *sportStoreEntryNewResp, nil
}

// SportStoreModifyNew 场馆信息修改
func (c *CultureSportCenter) SportStoreModifyNew(ctx context.Context, biz SportStoreModifyReq, opts ...ValueOptions) (SportStoreModifyResp, error) {
	method := "alipay.commerce.sports.venue.simple.modify"
	req, err := c.Client.NewRequest(method, biz, opts...)
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
