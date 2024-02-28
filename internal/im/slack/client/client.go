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
	"net/http"
	"strings"

	"github.com/lenye/pmsg/internal/im/slack"
	"github.com/lenye/pmsg/pkg/httpclient"
)

// PostJSON http post json
func PostJSON(url, reqBody string) (http.Header, error) {
	resp, err := httpclient.Post(url, httpclient.HdrValApplicationJson, strings.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("%w; %s %s, %v", httpclient.ErrRequest, http.MethodPost, url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return resp.Header, fmt.Errorf("%w; rate limit exceeded, retry after %s second", slack.ErrRequest, resp.Header.Get("Retry-After"))
	}

	// Slack seems to send an HTML body along with 5xx error codes. Don't parse it.
	if resp.StatusCode != http.StatusOK {
		return resp.Header, fmt.Errorf("%w; server error: %s", slack.ErrRequest, resp.Status)
	}

	return resp.Header, nil
}
