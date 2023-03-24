package tiktok_open_api

import (
	"encoding/json"
	"github.com/lemuzhi/open-api/utils"
	"net/http"
	"net/url"
)

type GetAccessToken struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	} `json:"data"`
}

// GetAccessToken 接口调用凭证相关接口
func (t *TiktokOpenApi) GetAccessToken() (*GetAccessToken, error) {
	u := "https://developer.toutiao.com/api/apps/v2/token"

	param := url.Values{}
	param.Set("appid", t.AppID)
	param.Set("secret", t.Secret)
	param.Set("grant_type", "client_credential")

	body, err := utils.Request(http.MethodPost, u, nil, param)
	if err != nil {
		return nil, err
	}

	var respData GetAccessToken
	err = json.Unmarshal(*body, &respData)
	return &respData, err
}
