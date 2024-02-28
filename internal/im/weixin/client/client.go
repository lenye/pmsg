// Copyright 2022-2024 The pmsg Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lenye/pmsg/pkg/helper"
	"github.com/lenye/pmsg/pkg/httpclient"
)

// CheckHttpResponseStatusCode 检查HTTP响应状态码
func CheckHttpResponseStatusCode(method, url string, statusCode int) error {
	if statusCode/100 != 2 {
		return fmt.Errorf("%w; http response status code: %v, %s %s", httpclient.ErrRequest, statusCode, method, url)
	}
	return nil
}

func GetJSON(url string, respBody any) (http.Header, error) {
	resp, err := httpclient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w; %s %s, %v", httpclient.ErrRequest, http.MethodGet, url, err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if err := CheckHttpResponseStatusCode(http.MethodGet, url, resp.StatusCode); err != nil {
		return nil, err
	}

	return resp.Header, helper.JsonDecode(resp.Body, respBody)
}

// PostJSON http post json
func PostJSON(url string, reqBody, respBody any) (http.Header, error) {
	body, err := helper.JsonEncode(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := httpclient.Post(url, httpclient.HdrValApplicationJson, body)
	if err != nil {
		return nil, fmt.Errorf("%w; %s %s, %v", httpclient.ErrRequest, http.MethodPost, url, err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if err := CheckHttpResponseStatusCode(http.MethodPost, url, resp.StatusCode); err != nil {
		return nil, err
	}

	return resp.Header, helper.JsonDecode(resp.Body, respBody)
}

func PostFileJSON(url, fieldName, fileName string, respBody any) (http.Header, error) {
	resp, err := httpclient.PostFile(url, fieldName, fileName)
	if err != nil {
		return nil, fmt.Errorf("%w; %s %s, %v", httpclient.ErrRequest, http.MethodPost, url, err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if err := CheckHttpResponseStatusCode(http.MethodPost, url, resp.StatusCode); err != nil {
		return nil, err
	}

	return resp.Header, helper.JsonDecode(resp.Body, respBody)
}
