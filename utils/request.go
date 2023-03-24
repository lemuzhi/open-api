package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Request(method, url string, header map[string]string, params any) (*[]byte, error) {
	cli := http.Client{}

	var body io.Reader = nil
	if params != nil {
		binary, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(binary)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	binaryBody, err := io.ReadAll(res.Body)
	return &binaryBody, err
}
