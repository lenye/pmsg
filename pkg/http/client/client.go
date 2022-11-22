// Copyright 2022 The pmsg Authors. All rights reserved.
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
	"bufio"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"runtime"

	"github.com/lenye/pmsg/pkg/version"
)

var UserAgent string

var ErrHttpRequest = errors.New("http request error")

const (
	contentTypeJson = "application/json;charset=utf-8"
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
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent())

	return http.DefaultClient.Do(req)
}

// Post http post
func Post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	req.Header.Set("User-Agent", userAgent())

	return http.DefaultClient.Do(req)
}

// PostFile 上传文件
func PostFile(url string, fieldName, fileName string) (*http.Response, error) {
	var bodyBuf bufio.ReadWriter
	bodyWriter := multipart.NewWriter(&bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(fieldName, fileName)
	if err != nil {
		return nil, fmt.Errorf("multipart.Writer.CreateFormFile failed, %w", err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("open file failed, %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(fileWriter, f); err != nil {
		return nil, fmt.Errorf("file io.Copy failed, %w", err)
	}

	contentType := bodyWriter.FormDataContentType()
	if err := bodyWriter.Close(); err != nil {
		return nil, fmt.Errorf("multipart.Writer.Close failed, %w", err)
	}

	return Post(url, contentType, bodyBuf)
}
