package errPkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	ew "github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const(
	JSON = "application/json"
)

func BasicSendPostWithBody(url string) error {

	reqBody := "{\"name\":\"dienvt\"}"
	resp, err := sendPost(url, reqBody)
	if err != nil {
		return ew.Wrap(err, "doSend")
	}

	if err := receiveData(resp); err != nil {
		return ew.Wrap(err, "receiveData")
	}

	return nil
}

func receiveData(resp *http.Response) error {
	if respBody, err := extractResponse(resp); err != nil {
		return ew.Wrap(err, "readBody")
	} else {
		fmt.Println("ok", respBody)
		if baseResp, uErr := unmarshal(respBody); uErr != nil {
			return ew.Wrap(uErr, "unmarshal")
		} else {
			fmt.Println("final result", baseResp)
		}
	}
	return nil
}

func unmarshal(resp string) (BaseResponse, error) {
	var r BaseResponse
	if err := json.Unmarshal([]byte(resp), &r); err != nil {
		return BaseResponse{}, ew.Wrap(err, "json.Unmarshal fail")
	}
	return r, nil
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func extractResponse(resp *http.Response) (string, error) {

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("read body fail %e", err))
	}

	return string(body), nil
}

func sendPost(url string, reqBody string) (*http.Response, error) {
	cli := http.Client{}
	reader := bytes.NewReader([]byte(reqBody))

	resp, err := cli.Post(url, JSON, reader)
	if err != nil {
		return nil, ew.Wrap(err, "send fail")
	}
	return resp, nil
}
