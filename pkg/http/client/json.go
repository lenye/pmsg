package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CheckHttpResponseStatusCode 检查HTTP响应状态码
func CheckHttpResponseStatusCode(url string, statusCode int) error {
	if statusCode/100 != 2 {
		return fmt.Errorf("%w; url: %q, http response status code: %v", ErrHttpRequest, url, statusCode)
	}
	return nil
}

func GetJSON(url string, respBody interface{}) (http.Header, error) {
	resp, err := Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w; get %q, %v", ErrHttpRequest, url, err)
	}
	defer resp.Body.Close()

	if err := CheckHttpResponseStatusCode(url, resp.StatusCode); err != nil {
		return nil, err
	}

	if respBody == nil {
		return resp.Header, nil
	}

	return resp.Header, json.NewDecoder(resp.Body).Decode(respBody)
}

// PostJSON http post json
func PostJSON(url string, reqBody, respBody interface{}) (http.Header, error) {
	jsonBuf := bytes.NewBuffer(make([]byte, 0, bufferSize))
	enc := json.NewEncoder(jsonBuf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := POST(url, contentTypeJson, jsonBuf)
	if err != nil {
		return nil, fmt.Errorf("%w; post %q, %v", ErrHttpRequest, url, err)
	}
	defer resp.Body.Close()

	if err := CheckHttpResponseStatusCode(url, resp.StatusCode); err != nil {
		return nil, err
	}

	if respBody == nil {
		return resp.Header, nil
	}

	return resp.Header, json.NewDecoder(resp.Body).Decode(respBody)
}
