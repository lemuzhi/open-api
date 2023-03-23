package tiktok_open_api

import (
	"encoding/json"
	"io"
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
	u := "https://api.weixin.qq.com/sns/jscode2session"

	param := url.Values{}
	param.Set("appid", t.AppID)
	param.Set("secret", t.Secret)
	param.Set("code", code)
	p := param.Encode()

	cli := http.Client{}

	req, err := http.NewRequest(http.MethodGet, u+"?"+p, nil)
	if err != nil {
		return nil, err
	}

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var j2s Jscode2session
	err = json.Unmarshal(resb, &j2s)
	if err != nil {
		return nil, err
	}

	return &j2s, nil
}
