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

package httpclient

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/lenye/pmsg/pkg/version"
)

var ErrRequest = errors.New("http request error")

const (
	HdrKeyUserAgent       = "User-Agent"
	HdrKeyContentType     = "Content-Type"
	HdrValApplicationJson = "application/json"
)

const (
	Timeout = 5 * time.Second
)

var Default = &http.Client{
	Timeout: Timeout,
}

var userAgent string

func DefaultUserAgent() string {
	return fmt.Sprintf("%s/%s (%s; %s) %s/%s", version.AppName, version.Version, runtime.GOOS, runtime.GOARCH, version.BuildCommit, version.BuildTime)
}

func SetUserAgent(v string) {
	userAgent = v
}

func UserAgent() string {
	if userAgent != "" {
		return userAgent
	}
	return DefaultUserAgent()
}

func SetTransport(v *http.Transport) {
	Default.Transport = v
}

// Get http get
func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HdrKeyUserAgent, UserAgent())

	return Default.Do(req)
}

// Post http post
func Post(url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HdrKeyUserAgent, UserAgent())
	req.Header.Set(HdrKeyContentType, contentType)

	return Default.Do(req)
}

func fileToBody(bodyWriter *multipart.Writer, formName, fileName string) (err error) {
	var fileWriter io.Writer
	fileWriter, err = bodyWriter.CreateFormFile(formName, filepath.Base(fileName))
	if err != nil {
		return fmt.Errorf("multipart.Writer.CreateFormFile failed, %w", err)
	}

	var f *os.File
	f, err = os.Open(fileName)
	if err != nil {
		return fmt.Errorf("open file failed, %w", err)
	}
	defer func() {
		if tmpErr := f.Close(); tmpErr != nil {
			err = fmt.Errorf("close file failed, %w", tmpErr)
		}
	}()

	if _, err := io.Copy(fileWriter, f); err != nil {
		return fmt.Errorf("io.Copy failed, %w", err)
	}

	return nil
}

// MultipartForm 保存文件或其他字段信息
type MultipartForm struct {
	params map[string][]string
	files  map[string]string
}

func NewMultipartForm() *MultipartForm {
	return &MultipartForm{
		params: make(map[string][]string),
		files:  make(map[string]string),
	}
}

// AddFile 保存文件信息
func (t *MultipartForm) AddFile(name, fileName string) *MultipartForm {
	t.files[name] = fileName
	return t
}

// AddParam 保存参数信息
func (t *MultipartForm) AddParam(name, value string) *MultipartForm {
	if param, ok := t.params[name]; ok {
		t.params[name] = append(param, value)
	} else {
		t.params[name] = []string{value}
	}
	return t
}

// PostMultipartForm 上传文件或其他多个字段
func PostMultipartForm(url string, form *MultipartForm) (*http.Response, error) {
	bodyBuf := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuf)

	for formName, fileName := range form.files {
		if err := fileToBody(bodyWriter, formName, fileName); err != nil {
			return nil, err
		}
	}
	for k, v := range form.params {
		for _, vv := range v {
			if err := bodyWriter.WriteField(k, vv); err != nil {
				return nil, fmt.Errorf("multipart.Writer.WriteField failed, %w", err)
			}
		}
	}
	contentType := bodyWriter.FormDataContentType()
	if err := bodyWriter.Close(); err != nil {
		return nil, fmt.Errorf("multipart.Writer.Close failed, %w", err)
	}
	return Post(url, contentType, bodyBuf)
}

// PostFile 上传文件
func PostFile(url, formName, fileName string) (*http.Response, error) {
	form := NewMultipartForm().AddFile(formName, fileName)
	return PostMultipartForm(url, form)
}
