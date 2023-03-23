package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Requests(url, method string, header map[string]string, params interface{}) ([]byte, error) {
	var req *http.Request
	var err error
	client := http.Client{}
	if params != nil {
		binary, err := json.Marshal(params)
		if err != nil {
			log.Printf("json.Marshal(params) error: %v, url: %v, params: %v ", err, url, params)
			return nil, errors.New("json.Marshal(params) error")
		}
		fmt.Println(string(binary))
		req, err = http.NewRequest(method, url, bytes.NewReader(binary))
		//req, err = http.NewRequest(method, url, strings.NewReader(string(binary)))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Printf("http.NewRequest error: %v, url: %v, method: %v, params: %v ", err, url, method, params)
		return nil, errors.New("http.NewRequest error ")
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do error: %v, url: %v, method: %v, params: %v ", err, url, method, params)
		return nil, errors.New("client.Do error ")
	}

	defer resp.Body.Close()

	binaryBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("io.ReadAll(resp.Body) error: %v, url: %v, data: %v ", err, url, binaryBody)
		return nil, errors.New("io.ReadAll(resp.Body) error ")
	}
	//fmt.Println(string(binaryBody))
	return binaryBody, err

}
