package api_call_token

import (
	"encoding/json"
	"github.com/lemuzhi/open-api/global"
	"github.com/lemuzhi/open-api/utils"
	"net/http"
	"net/url"
)

type GetAccessTokenResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	} `json:"data"`
}

// GetAccessToken 接口调用凭证相关接口
func (t *ApiCallToken) GetAccessToken() (*GetAccessTokenResp, error) {
	u := "https://developer.toutiao.com/api/apps/v2/token"

	param := url.Values{}
	param.Set("appid", global.AppID)
	param.Set("secret", global.Secret)
	param.Set("grant_type", "client_credential")

	body, err := utils.Request(http.MethodPost, u, nil, param)
	if err != nil {
		return nil, err
	}

	var respData GetAccessTokenResp
	err = json.Unmarshal(*body, &respData)
	return &respData, err
}
