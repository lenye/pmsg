package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetJSON(url string, respBody interface{}) (http.Header, error) {
	resp, err := Get(url)
	if err != nil {
		return nil, fmt.Errorf("http failed to get %q: %w", url, err)
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
		return nil, fmt.Errorf("http failed to post %q: %w", url, err)
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
