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

package token

import (
	"fmt"
	"net/url"
	"time"

	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/internal/im/weixin/client"
)

type AccessTokenMeta struct {
	AccessToken string    `json:"access_token"`        // 微信接口调用凭证
	ExpireIn    int64     `json:"expires_in"`          // 微信接口调用凭证有效时间，单位：秒
	ExpireAt    time.Time `json:"expire_at,omitempty"` // 微信接口调用凭证到期时间
}

func (t AccessTokenMeta) String() string {
	if t.ExpireAt.IsZero() {
		return fmt.Sprintf("access_token: %q, expires_in: %v", t.AccessToken, t.ExpireIn)
	}
	return fmt.Sprintf("access_token: %q, expires_in: %v, expire_at: %q", t.AccessToken, t.ExpireIn, t.ExpireAt.Format(time.RFC3339))
}

// AccessTokenResponse 响应
type AccessTokenResponse struct {
	weixin.ResponseMeta
	AccessTokenMeta
}

const reqURL = weixin.Host + "/cgi-bin/token?grant_type=client_credential&appid="

// FetchAccessToken 获取微信接口调用凭证
// 正常情况下，微信会返回下述 JSON
// {"access_token":"ACCESS_TOKEN","expires_in":7200}
// 错误时微信会返回错误码等信息，JSON数据包示例如下:
// {"errcode":40013,"errmsg":"invalid appid"}
func FetchAccessToken(appID, appSecret string) (*AccessTokenMeta, error) {
	u := reqURL + url.QueryEscape(appID) + "&secret=" + url.QueryEscape(appSecret)
	var resp AccessTokenResponse
	_, err := client.GetJSON(u, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; %v", weixin.ErrRequest, resp.ResponseMeta)
	}

	resp.AccessTokenMeta.ExpireAt = time.Now().Add(time.Second * time.Duration(resp.AccessTokenMeta.ExpireIn))

	return &resp.AccessTokenMeta, nil
}
