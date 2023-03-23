package tiktok_open_api

import "fmt"

// 登录相关接口

type Jscode2session struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		SessionKey      string `json:"session_key"`
		Openid          string `json:"openid"`
		AnonymousOpenid string `json:"anonymous_openid"`
		Unionid         string `json:"unionid"`
	} `json:"data"`
}

func (t *TiktokOpenApi) Jscode2session() {
	fmt.Println("1======", t.AppID)
	fmt.Println("2======", t.Secret)
	fmt.Println("3======", t.Salt)
}
