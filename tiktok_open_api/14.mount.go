package tiktok_open_api

import (
	"encoding/json"
	"github.com/lemuzhi/open-api/utils"
	"net/http"
)

// 挂载相关接口

type ShortVideoMount struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Data   struct {
		QrcodeURL          string `json:"qrcode_url"`
		QrcodeParseContent string `json:"qrcode_parse_content"`
	} `json:"data"`
}

func (t *TiktokOpenApi) ShortVideoMount(accessToken string) (*ShortVideoMount, error) {
	u := "https://developer.toutiao.com/api/apps/v1/capacity/get_self_mount_bind_qrcode"

	header := map[string]string{
		"access-token": accessToken,
	}

	body, err := utils.Request(http.MethodGet, u+"?capacity_key=video_self_mount", header, nil)
	if err != nil {
		return nil, err
	}

	var respData ShortVideoMount
	err = json.Unmarshal(*body, &respData)
	return &respData, err
}
