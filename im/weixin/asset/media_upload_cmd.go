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

package asset

import (
	"fmt"

	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/httpclient"
	"github.com/lenye/pmsg/im"
	"github.com/lenye/pmsg/im/weixin"
	"github.com/lenye/pmsg/im/weixin/token"
)

type CmdMediaUploadParams struct {
	UserAgent   string
	AccessToken string
	AppID       string
	AppSecret   string
	MediaType   string
	File        string
}

func (t *CmdMediaUploadParams) Validate() error {
	if t.AccessToken == "" && t.AppID == "" {
		return flags.ErrWeixinAccessToken
	}
	if err := ValidateMediaType(t.MediaType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MediaType, err)
	}

	if !im.FileExists(t.File) {
		return fmt.Errorf("file is not exist, %v", t.File)
	}

	return nil
}

// CmdMediaUpload 新增临时素材 微信公众号/小程序
func CmdMediaUpload(arg *CmdMediaUploadParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if arg.AccessToken == "" {
		accessTokenResp, err := token.FetchAccessToken(arg.AppID, arg.AppSecret)
		if err != nil {
			return fmt.Errorf("%w, %w", weixin.ErrRequest, err)
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if meta, err := MediaUpload(arg.AccessToken, arg.MediaType, arg.File); err != nil {
		return fmt.Errorf("%w, %w", weixin.ErrRequest, err)
	} else {
		fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, meta))
	}

	return nil
}
