package escrow_pay

import (
	"encoding/json"
	"fmt"
	open_api "github.com/lemuzhi/open-api"
	"log"
	"time"
)

const (
	appid  = ""
	secret = ""
	salt   = ""
	code   = ""
)

func ExampleOrderPush_OrderSync() {
	tk := open_api.NewTiktokOpenApi(appid, secret, salt)
	itemList := map[string]interface{}{
		"item_code": "2",
		"img":       "https://oss-article1.oss-cn-beijing.aliyuncs.com/media/tyh/6a62326f35b603702b105e80138989c.jpg",
		"title":     "子订单商品介绍标题",
		"sub_title": "子订单商品介绍副标题",
		"price":     0.01 * 100, //单类商品的总价，示例：1分钱
	}
	orderDetail := map[string]interface{}{
		"order_id":    "D931825839218563728",   //开发者侧业务单号(就是你自己设置的，随便你自己用什么生成，只要唯一不重复即可),长度 <= 64byte
		"create_time": 1683848340000 / 1e6,     //订单创建的时间，这是13位的毫秒级时间戳
		"status":      1,                       //订单状态，待支付，已支付，已取消，已超时，已核销，退款中，已退款，退款失败，填自己设置的枚举值
		"amount":      1,                       //订单商品总数
		"total_price": 0.01 * 100,              //订单总金额，示例：100元
		"detail_url":  "pages/my/order/order",  //小程序订单详情页 path，长度<=1024 byte，用户在抖音订单记录里点击订单，跳转的小程序页面
		"item_list":   []interface{}{itemList}, //子订单商品列表，不可为空
	}
	item, err := json.Marshal(orderDetail)
	if err != nil {
		log.Fatalln(err)
	}
	param := map[string]interface{}{
		"access_token": "",                                     //这里放通过服务端OpenAPI获取到的access_token，该接口地址：https://developer.toutiao.com/api/apps/v2/token
		"app_name":     "douyin",                               //默认，固定值douyin
		"open_id":      "_000Nk6QrtXaHrXIBTKnVA8lwKPp9O2-3G0y", //小程序用户的 open_id
		"order_detail": string(item),                           //json string，根据不同订单类型有不同的结构体，请参见 order_detail 字段说明（json string）
		"order_status": 0,                                      //0待支付 1已支付 2已取消（用户主动取消或者超时未支付导致的关单） 4：已核销（核销状态是整单核销,即一笔订单买了 3 个券，核销是指 3 个券核销的整单） 5：退款中 6：已退款8：退款失败，注意：普通小程序订单必传，担保支付分账依赖该状态
		"order_type":   0,                                      // 订单类型，0：普通小程序订单（非POI订单）9101：团购券订单（POI 订单）9001：景区门票订单（POI订单）
		"update_time":  time.Now().UnixNano() / 1e6,            // 订单信息变更时间，13 位毫秒级时间戳
	}
	res, err := tk.EscrowPay.OrderPush.OrderSync(itemList, orderDetail, param)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)

	// Output:
	//
}
