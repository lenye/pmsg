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

package message

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/internal/im/weixin/token"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdMpBizSendSubscribeParams struct {
	UserAgent   string
	AccessToken string
	AppID       string
	AppSecret   string
	ToUser      string
	TemplateID  string
	Page        string
	Mini        map[string]string
	Data        string
}

func (t *CmdMpBizSendSubscribeParams) Validate() error {
	if t.AccessToken == "" && t.AppID == "" {
		return flags.ErrWeixinAccessToken
	}

	// 跳小程序
	if t.Mini != nil {
		if miniAppID, ok := t.Mini[flags.MiniAppID]; !ok {
			return fmt.Errorf("mini flag %q not set", flags.MiniAppID)
		} else {
			if miniAppID == "" {
				return fmt.Errorf("mini flag %q not set", flags.MiniAppID)
			}
		}

		if miniPagePath, ok := t.Mini[flags.MiniPagePath]; !ok {
			return fmt.Errorf("mini flag %q not set", flags.MiniPagePath)
		} else {
			if miniPagePath == "" {
				return fmt.Errorf("mini flag %q not set", flags.MiniPagePath)
			}
		}
	}

	return nil
}

// CmdMpBizSendSubscribe 发送微信公众号订阅通知消息
func CmdMpBizSendSubscribe(arg *CmdMpBizSendSubscribeParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	var dataItem map[string]SubscribeDataItem
	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	if buf.String() != "" {
		if err := json.Unmarshal(buf.Bytes(), &dataItem); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		for k, v := range dataItem {
			if v.Value == "" {
				return fmt.Errorf("data %v.value not set", k)
			}
		}
	}

	msg := SubscribeMessage{
		ToUser:     arg.ToUser,
		TemplateID: arg.TemplateID,
		Page:       arg.Page,
		Data:       dataItem,
	}

	// 跳小程序
	if arg.Mini != nil {
		miniAppID, _ := arg.Mini[flags.MiniAppID]
		miniPagePath, _ := arg.Mini[flags.MiniPagePath]
		msg.MiniProgram = &MiniProgramMeta{
			AppID:    miniAppID,
			PagePath: miniPagePath,
		}
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if arg.AccessToken == "" {
		accessTokenResp, err := token.FetchAccessToken(arg.AppID, arg.AppSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if err := BizSendSubscribe(arg.AccessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
