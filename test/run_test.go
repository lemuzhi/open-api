package test

import (
	"fmt"
	"github.com/lemuzhi/open-api"
	"log"
	"testing"
)

const (
	appid  = ""
	secret = ""
	salt   = ""
	code   = ""
)

func TestRun(t *testing.T) {
	tk := open_api.NewTiktokOpenApi(appid, secret, salt)
	//res, err := tk.Jscode2session(code)
	//res, err := tk.ShortVideoMount("0801121846313073694773626f3359725a507377724c542f34413d3d")
	//res, err := tk.EscrowPay.Pay.QueryOrder("7013158683625009646")
	//res, err := tk.EscrowPay.Pay.CreateOrder("7013158683611111111", "测试", "测试1", 0.01, 300)
	res, err := tk.ApiCallToken.GetAccessToken()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
