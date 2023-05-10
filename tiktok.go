package open_api

import (
	"fmt"
	"github.com/lemuzhi/open-api/global"
	"github.com/lemuzhi/open-api/tiktok_open_api/api_call_token"
	"github.com/lemuzhi/open-api/tiktok_open_api/escrow_pay"
	"github.com/lemuzhi/open-api/tiktok_open_api/login"
	"github.com/lemuzhi/open-api/tiktok_open_api/mount"
)

type Tiktok struct {
	api_call_token.ApiCallToken // 接口调用凭证
	login.Login                 // 登录
	escrow_pay.EscrowPay        // 担保支付
	mount.Mount                 //挂载
}

func NewTiktokOpenApi(appid, secret, salt string) *Tiktok {
	fmt.Println("参数", appid, secret)
	global.AppID = appid
	global.Secret = secret
	global.Salt = salt
	return &Tiktok{}
}
