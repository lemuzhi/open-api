package escrow_pay

import (
	"encoding/json"
	"fmt"
	"github.com/lemuzhi/open-api/utils"
)

type OrderPush struct {
}

type OrderSyncResp struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Body    string `json:"body"`
}

func Pabc() {
	fmt.Println("test")
}

// OrderSync 订单同步
// 将小程序订单同步值抖音订单中心
func (i *OrderPush) OrderSync(itemList, orderDetail, param map[string]interface{}) (*OrderSyncResp, error) {
	url := "https://developer.toutiao.com/api/apps/order/v2/push"
	header := map[string]string{
		"Content-Type": "application/json",
	}

	orderDetail["item_list"] = []interface{}{itemList}

	item, err := json.Marshal(orderDetail)
	if err != nil {
		return nil, err
	}

	param["order_detail"] = string(item)

	res, err := utils.Request("POST", url, header, param)
	if err != nil {
		return nil, err
	}

	var respData OrderSyncResp
	err = json.Unmarshal(*res, &respData)
	return &respData, err
}
