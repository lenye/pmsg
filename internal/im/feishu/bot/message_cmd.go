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
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/feishu"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdSendParams struct {
	UserAgent   string
	AccessToken string
	Secret      string
	MsgType     string
	Data        string
}

func (t *CmdSendParams) Validate() error {
	if err := ValidateMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdSend 发送飞书自定义机器人消息
func CmdSend(arg *CmdSendParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := Message{
		MsgType: arg.MsgType,
	}

	if arg.Secret != "" {
		msg.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
		sign := feishu.Sign(msg.TimeStamp, arg.Secret)
		msg.Sign = sign
	}

	buf := new(bytes.Buffer)
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case MsgTypeText:
		var msgMeta ContentMeta
		msgMeta.Text = buf.String()
		msg.Content = &msgMeta
	case MsgTypeImage:
		var msgMeta ContentMeta
		msgMeta.ImageKey = buf.String()
		msg.Content = &msgMeta
	case MsgTypeShareChat:
		var msgMeta ContentMeta
		msgMeta.ShareChatID = buf.String()
		msg.Content = &msgMeta
	case MsgTypePost:
		var post PostMeta
		if err := json.Unmarshal(buf.Bytes(), &post); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msgMeta := &ContentMeta{
			Post: &post,
		}
		msg.Content = msgMeta
	case MsgTypeInteractive:
		var msgMeta CardMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.Card = &msgMeta
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if err := Send(arg.AccessToken, &msg); err != nil {
		return err
	}
	fmt.Println(feishu.MessageOK)

	return nil
}
