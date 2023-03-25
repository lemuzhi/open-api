package tiktok_open_api

import (
	"encoding/json"
	"github.com/lemuzhi/open-api/utils"
	"net/http"
	"net/url"
)

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

func (t *TiktokOpenApi) Jscode2session(code string) (*Jscode2session, error) {
	u := "https://developer.toutiao.com/api/apps/v2/jscode2session"

	param := url.Values{}
	param.Set("appid", t.AppID)
	param.Set("secret", t.Secret)
	param.Set("code", code)
	//p := param.Encode()

	res, err := utils.Request(http.MethodPost, u, nil, param)

	var respData Jscode2session
	err = json.Unmarshal(*res, &respData)
	return &respData, err
}
