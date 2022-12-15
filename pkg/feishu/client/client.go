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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	httpClient "github.com/lenye/pmsg/pkg/http/client"
)

// PostJSON http post json
func PostJSON(url string, reqBody, respBody any) (http.Header, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Post(url, httpClient.HdrValContentTypeJsonCharset, buf)
	if err != nil {
		return nil, fmt.Errorf("%w; %s %s, %v", httpClient.ErrRequest, http.MethodPost, url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return resp.Header, json.NewDecoder(resp.Body).Decode(respBody)
	}

	return resp.Header, nil
}
