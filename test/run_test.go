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
	res, err := tk.Jscode2session(code)
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println(res)
}
