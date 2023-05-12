package api_call_token

import (
	"fmt"
	open_api "github.com/lemuzhi/open-api"
	"log"
)

const (
	appid  = ""
	secret = ""
	salt   = ""
	code   = ""
)

func ExampleApiCallToken_GetAccessToken() {
	tk := open_api.NewTiktokOpenApi(appid, secret, salt)
	res, err := tk.ApiCallToken.GetAccessToken()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)

	// Output:
	//
}
