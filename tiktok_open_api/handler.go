package tiktok_open_api

type TiktokOpenApi struct {
	AppID  string
	Secret string
	Salt   string
}

func (*TiktokOpenApi) NewTiktokOpenApi(appid, secret, salt string) *TiktokOpenApi {
	return &TiktokOpenApi{
		AppID:  appid,
		Secret: secret,
		Salt:   salt,
	}
}
