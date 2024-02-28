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

	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdWorkTokenParams struct {
	UserAgent  string
	CorpID     string
	CorpSecret string
}

// CmdWorkGetAccessToken 获取企业微信接口调用凭证
func CmdWorkGetAccessToken(arg *CmdWorkTokenParams) error {

	httpclient.SetUserAgent(arg.UserAgent)

	accessTokenResp, err := FetchAccessToken(arg.CorpID, arg.CorpSecret)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, accessTokenResp))

	return nil
}
