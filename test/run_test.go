package test

import (
	"fmt"
	"github.com/lemuzhi/open-api/tiktok_open_api"
	"testing"
)

const (
	appid  = ""
	secret = ""
	salt   = ""
	code   = ""
)

func TestRun(t *testing.T) {
	tk := tiktok_open_api.NewTiktokOpenApi(appid, secret, salt)
	//res, err := tk.Jscode2session(code)
	res, err := tk.ShortVideoMount("0801121846313073694773626f3359725a507377724c542f34413d3d")
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println(res)
}
