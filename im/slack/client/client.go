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

	"github.com/lenye/pmsg/httpclient"
	"github.com/lenye/pmsg/im"
)

func PostJSON(url string, reqBody any) (http.Header, error) {
	body, err := im.JsonEncode(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := httpclient.Post(url, httpclient.HdrValApplicationJson, body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == http.StatusTooManyRequests {
		return resp.Header, fmt.Errorf("rate limit exceeded, retry after %s second", resp.Header.Get("Retry-After"))
	}

	// Slack seems to send an HTML body along with 5xx error codes. Don't parse it.
	if resp.StatusCode != http.StatusOK {
		return resp.Header, fmt.Errorf("invalid http.response.status: %s", resp.Status)
	}

	return resp.Header, nil
}
