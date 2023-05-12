package api_call_token

import (
	"encoding/json"
	"github.com/lemuzhi/open-api/global"
	"github.com/lemuzhi/open-api/utils"
	"net/http"
)

type GetAccessTokenResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		ExpiresAt   int    `json:"expires_at"`
	} `json:"data"`
}

// GetAccessToken 接口调用凭证相关接口
func (t *ApiCallToken) GetAccessToken() (*GetAccessTokenResp, error) {
	u := "https://developer.toutiao.com/api/apps/v2/token"

	header := map[string]string{
		"content-type": "application/json",
	}
	params := map[string]string{
		"appid":      global.AppID,
		"secret":     global.Secret,
		"grant_type": "client_credential",
	}

	body, err := utils.Request(http.MethodPost, u, header, params)
	if err != nil {
		return nil, err
	}

	var respData GetAccessTokenResp
	err = json.Unmarshal(*body, &respData)
	return &respData, err
}
