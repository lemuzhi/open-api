package escrow_pay

import (
	"encoding/json"
	"fmt"
	"github.com/lemuzhi/open-api/global"
	"github.com/lemuzhi/open-api/utils"
)

type Pay struct {
}

type QueryOrderResp struct {
	ErrNo       int    `json:"err_no"`
	ErrTips     string `json:"err_tips"`
	OutOrderNo  string `json:"out_order_no"`
	OrderID     string `json:"order_id"`
	PaymentInfo struct {
		TotalFee    int    `json:"total_fee"`
		OrderStatus string `json:"order_status"`
		PayTime     string `json:"pay_time"`
		Way         int    `json:"way"`
		ChannelNo   string `json:"channel_no"`
		SellerUID   string `json:"seller_uid"`
		ItemID      string `json:"item_id"`
		CpsInfo     string `json:"cps_info"`
	} `json:"payment_info"`
}

func (p *Pay) QueryOrder(orderNo string) (*QueryOrderResp, error) {
	//查询订单结果只有out_order_no加签
	paramsMap := map[string]interface{}{
		"app_id":        global.AppID, //小程序APPID
		"thirdparty_id": "",           //第三方平台服务商 id，非服务商模式留空
		"out_order_no":  orderNo,      //开发者侧的订单号。 只能是数字、大小写字母_-*且在同一个商户号下唯一
		"sign":          "",           //签名，详见签名DEMO:https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/open-capacity/guaranteed-payment/TE/#_%E4%B8%89%E7%AD%BE%E5%90%8D-demo
	}
	fmt.Println("salt=", global.AppID)
	sign := utils.RequestSign(paramsMap, global.Salt)
	fmt.Println("sign=", sign)
	body := map[string]interface{}{
		"app_id":       global.AppID,
		"out_order_no": orderNo, //开发者侧的订单号, 同一小程序下不可重复
		"sign":         sign,    //sign
	}
	url := "https://developer.toutiao.com/api/apps/ecpay/v1/query_order"

	header := map[string]string{
		"Content-Type": "application/json",
	}

	res, err := utils.Request("POST", url, header, body)
	if err != nil {
		return nil, err
	}
	var respData QueryOrderResp
	err = json.Unmarshal(*res, &respData)
	return &respData, err
}
