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

type CmdMiniSendSubscribeParams struct {
	UserAgent        string
	AccessToken      string
	AppID            string
	AppSecret        string
	ToUser           string
	TemplateID       string
	MiniProgramState string
	Page             string
	Language         string
	Data             string
}

func (t *CmdMiniSendSubscribeParams) Validate() error {
	if t.AccessToken == "" && t.AppID == "" {
		return flags.ErrWeixinAccessToken
	}
	if t.Language != "" {
		if err := ValidateLanguage(t.Language); err != nil {
			return fmt.Errorf("invalid flags %s: %v", flags.Language, err)
		}
	}

	if t.MiniProgramState != "" {
		if err := ValidateMiniProgramState(t.MiniProgramState); err != nil {
			return fmt.Errorf("invalid flags %s: %v", flags.MiniProgramState, err)
		}
	}

	return nil
}

// CmdMiniProgramSendSubscribe 发送微信小程序订阅消息
func CmdMiniProgramSendSubscribe(arg *CmdMiniSendSubscribeParams) error {

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
		ToUser:           arg.ToUser,
		TemplateID:       arg.TemplateID,
		Data:             dataItem,
		Page:             arg.Page,
		MiniProgramState: MiniProgramStateFormal,
		Language:         LanguageZhCN,
	}

	if arg.Language != "" {
		msg.Language = arg.Language
	}

	if arg.MiniProgramState != "" {
		msg.MiniProgramState = arg.MiniProgramState
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if arg.AccessToken == "" {
		accessTokenResp, err := token.FetchAccessToken(arg.AppID, arg.AppSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if err := SendSubscribe(arg.AccessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
