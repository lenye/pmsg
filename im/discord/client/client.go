// Copyright 2022-2025 The pmsg Authors. All rights reserved.
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
	"strings"

	"github.com/lenye/pmsg/httpclient"
	"github.com/lenye/pmsg/im"
	"github.com/lenye/pmsg/im/discord"
)

func PostJSON(url string, reqBody, respBody any) (http.Header, error) {
	body, err := im.JsonEncode(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := httpclient.Post(url, httpclient.HdrValApplicationJson, body)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", httpclient.ErrRequest, err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == http.StatusNoContent {
		return resp.Header, nil
	}

	if strings.EqualFold(resp.Header.Get("content-type"), httpclient.HdrValApplicationJson) {
		if err := im.JsonDecode(resp.Body, respBody); err != nil {
			return nil, fmt.Errorf("json decode failed, %w", err)
		}
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return resp.Header, fmt.Errorf("%w, rate limit exceeded, retry after %s second", discord.ErrRequest, resp.Header.Get("Retry-After"))
	}

	// Slack seems to send an HTML body along with 5xx error codes. Don't parse it.
	if resp.StatusCode != http.StatusOK {
		return resp.Header, fmt.Errorf("%w, http response status: %s", discord.ErrRequest, resp.Status)
	}

	return resp.Header, nil
}
