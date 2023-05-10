package login

import (
	"encoding/json"
	"github.com/lemuzhi/open-api/global"
	"github.com/lemuzhi/open-api/utils"
	"net/http"
)

type code2sessionResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		SessionKey      string `json:"session_key"`
		Openid          string `json:"openid"`
		AnonymousOpenid string `json:"anonymous_openid"`
		Unionid         string `json:"unionid"`
	} `json:"data"`
}

func (t *Login) Jscode2session(code string) (*code2sessionResp, error) {
	u := "https://developer.toutiao.com/api/apps/v2/jscode2session"

	body := map[string]string{
		"appid":  global.AppID,
		"secret": global.Secret,
		"code":   code,
	}

	header := map[string]string{
		"content-type": "application/json",
	}

	res, err := utils.Request(http.MethodPost, u, header, body)

	var respData code2sessionResp
	err = json.Unmarshal(*res, &respData)
	return &respData, err
}
