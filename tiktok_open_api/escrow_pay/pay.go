package escrow_pay

import (
	"encoding/json"
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

type CreateOrderResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		OrderID    string `json:"order_id"`
		OrderToken string `json:"order_token"`
	} `json:"data"`
}

// CreateOrder 预下单接口
// orderNo 开发者侧的订单号
//
// subject 商品描述
// body 商品详情
// totalAmount 支付价格
// expiredAt 订单过期时间

func (p *Pay) CreateOrder(orderNo, subject, body string, totalAmount float64, expiredAt uint32) (*CreateOrderResp, error) {
	paramsMap := map[string]interface{}{
		"app_id":        global.AppID,      //小程序APPID
		"thirdparty_id": "",                //第三方平台服务商 id，非服务商模式留空
		"out_order_no":  orderNo,           //开发者侧的订单号。 只能是数字、大小写字母_-*且在同一个商户号下唯一
		"total_amount":  totalAmount * 100, //支付价格。 单位为[分]， 取值范围： [1,10000000000]
		"subject":       subject,           //商品描述。 长度限制不超过 128 字节且不超过 42 字符
		"body":          body,              //商品详情 长度限制不超过 128 字节且不超过 42 字符
		"valid_time":    expiredAt,         //订单过期时间(秒)。最小5分钟，最大2天，小于5分钟会被置为5分钟，大于2天会被置为2天，取值范围： [300,172800]，当前表示30分钟过期
		//"notify_url":    "https://www.xxx.com", //商户自定义回调地址，必须以 https 开头，支持 443 端口。 指定时，支付成功后抖音会请求该地址通知开发者
		//"disable_msg":   0, //是否屏蔽支付完成后推送用户抖音消息，1-屏蔽 0-非屏蔽，默认为0。 特别注意： 若接入POI, 请传1。因为POI订单体系会发消息，所以不用再接收一次担保支付推送消息，
		//"msg_page":      "pages/user/orderDetail/orderDetail?id=997979879879879879", //支付完成后推送给用户的抖音消息跳转页面
		"sign": "", //签名，详见签名DEMO:https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/open-capacity/guaranteed-payment/TE/#_%E4%B8%89%E7%AD%BE%E5%90%8D-demo
	}
	sign := utils.RequestSign(paramsMap, global.Salt)
	paramsMap["sign"] = sign

	u := "https://developer.toutiao.com/api/apps/ecpay/v1/create_order"

	header := map[string]string{
		"Content-Type": "application/json",
	}

	res, err := utils.Request("POST", u, header, paramsMap)
	if err != nil {
		return nil, err
	}

	var respData CreateOrderResp
	err = json.Unmarshal(*res, &respData)
	return &respData, err
}
func (p *Pay) QueryOrder(orderNo string) (*QueryOrderResp, error) {
	//查询订单结果只有out_order_no加签
	paramsMap := map[string]interface{}{
		"app_id":        global.AppID, //小程序APPID
		"thirdparty_id": "",           //第三方平台服务商 id，非服务商模式留空
		"out_order_no":  orderNo,      //开发者侧的订单号。 只能是数字、大小写字母_-*且在同一个商户号下唯一
		"sign":          "",           //签名，详见签名DEMO:https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/open-capacity/guaranteed-payment/TE/#_%E4%B8%89%E7%AD%BE%E5%90%8D-demo
	}
	sign := utils.RequestSign(paramsMap, global.Salt)
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
