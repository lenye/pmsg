package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/lenye/pmsg/pkg/version"
)

var UserAgent string

const (
	contentTypeJson = "application/json;charset=utf-8"
	bufferSize      = 1024
)

func DefaultUserAgent() string {
	return fmt.Sprintf("%s/%s (%s; %s) %s/%s", version.CodeName, version.Version, runtime.GOOS, runtime.GOARCH, version.BuildGit, version.BuildTime)
}

func userAgent() string {
	if UserAgent != "" {
		return UserAgent
	}
	return DefaultUserAgent()
}

// Get http get
func Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent())

	return http.DefaultClient.Do(req)
}

// POST http post
func POST(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	req.Header.Set("User-Agent", userAgent())

	return http.DefaultClient.Do(req)
}

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

	// log.Printf("PostJSON body: %s", jsonBuf.String())

	resp, err := POST(url, contentTypeJson, jsonBuf)
	if err != nil {
		return nil, fmt.Errorf("http failed to post %q: %w", url, err)
	}
	defer resp.Body.Close()

	// dumpResp, _ := httputil.DumpResponse(resp, true)
	// log.Printf("resp: %s", string(dumpResp))

	if err := CheckHttpResponseStatusCode(url, resp.StatusCode); err != nil {
		return nil, err
	}

	if respBody == nil {
		return resp.Header, nil
	}

	return resp.Header, json.NewDecoder(resp.Body).Decode(respBody)
}

// CheckHttpResponseStatusCode 检查HTTP响应状态码
func CheckHttpResponseStatusCode(uri string, statusCode int) error {
	if statusCode/100 != 2 {
		return fmt.Errorf("http request failed, uri=%q, statusCode=%d", uri, statusCode)
	}
	return nil
}
