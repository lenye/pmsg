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

	if resp.StatusCode/100 != 2 {
		return resp.Header, helper.JsonDecode(resp.Body, respBody)
	}

	return resp.Header, nil
}
