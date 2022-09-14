package client

import (
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
	return fmt.Sprintf("%s/%s (%s; %s) %s/%s", version.AppName, version.Version, runtime.GOOS, runtime.GOARCH, version.BuildGit, version.BuildTime)
}

func userAgent() string {
	if UserAgent != "" {
		return UserAgent
	}
	return DefaultUserAgent()
}

// Get http get
func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent())

	return http.DefaultClient.Do(req)
}

// POST http post
func POST(url string, bodyType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	req.Header.Set("User-Agent", userAgent())

	return http.DefaultClient.Do(req)
}

// CheckHttpResponseStatusCode 检查HTTP响应状态码
func CheckHttpResponseStatusCode(uri string, statusCode int) error {
	if statusCode/100 != 2 {
		return fmt.Errorf("http request failed, uri=%q, statusCode=%d", uri, statusCode)
	}
	return nil
}
