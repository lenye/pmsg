// Copyright 2022-2023 The pmsg Authors. All rights reserved.
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

	"github.com/lenye/pmsg/pkg/file"
	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/work/token"
)

type CmdWorkMediaUploadParams struct {
	UserAgent   string
	AccessToken string
	CorpID      string
	CorpSecret  string
	MediaType   string
	File        string
}

func (t *CmdWorkMediaUploadParams) Validate() error {
	if t.AccessToken == "" && t.CorpID == "" {
		return flags.ErrWeixinWorkAccessToken
	}

	if err := ValidateMediaType(t.MediaType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MediaType, err)
	}

	if !file.Exists(t.File) {
		return fmt.Errorf("file is not exist, %v", t.File)
	}

	return nil
}

// CmdWorkMediaUpload 上传临时素材
func CmdWorkMediaUpload(arg *CmdWorkMediaUploadParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	client.SetUserAgent(arg.UserAgent)

	if arg.AccessToken == "" {
		accessTokenResp, err := token.FetchAccessToken(arg.CorpID, arg.CorpSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if meta, err := MediaUpload(arg.AccessToken, arg.MediaType, arg.File); err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, meta))
	}

	return nil
}
