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

package bot

import (
	"fmt"

	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/pkg/helper"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdUploadParams struct {
	UserAgent string
	Key       string
	File      string
}

func (t *CmdUploadParams) Validate() error {
	if !helper.FileExists(t.File) {
		return fmt.Errorf("file is not exist, %v", t.File)
	}

	return nil
}

// CmdUpload 企业微信群机器人上传文件
func CmdUpload(arg *CmdUploadParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if meta, err := Upload(arg.Key, arg.File); err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, meta))
	}

	return nil
}
