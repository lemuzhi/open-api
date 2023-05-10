package test

import (
	"fmt"
	"github.com/lemuzhi/open-api"
	"testing"
)

const (
	appid  = "tta6a9d0143997e46301"
	secret = "088908088d57fcb413ff8d85eaf135d2c6557456"
	salt   = "N56SNeRoC5Y9LUWQbtVbrIwACBxi6slvfMzvZ5Pe"
	code   = ""
)

func TestRun(t *testing.T) {
	tk := open_api.NewTiktokOpenApi(appid, secret, salt)
	//res, err := tk.Jscode2session(code)
	//res, err := tk.ShortVideoMount("0801121846313073694773626f3359725a507377724c542f34413d3d")
	res, err := tk.OpenApi.EscrowPay.Pay.QueryOrder("7018002683625031257")
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println(res)
}
