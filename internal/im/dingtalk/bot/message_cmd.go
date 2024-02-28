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
	"strings"

	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/dingtalk"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdSendParams struct {
	UserAgent   string
	AccessToken string
	Secret      string
	MsgType     string
	AtUser      string
	AtMobile    string
	IsAtAll     bool
	Data        string
}

func (t *CmdSendParams) Validate() error {
	if err := ValidateMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdSend 发送钉钉自定义机器人消息
func CmdSend(arg *CmdSendParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := Message{
		MsgType: arg.MsgType,
	}

	buf := new(bytes.Buffer)
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case MsgTypeText:
		var msgMeta TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta

		var at AtMeta
		if arg.IsAtAll || arg.AtUser != "" || arg.AtMobile != "" {
			if arg.AtUser != "" {
				at.AtUserIds = strings.Split(arg.AtUser, "|")
			}
			if arg.AtMobile != "" {
				at.AtMobiles = strings.Split(arg.AtMobile, "|")
			}
			at.IsAtAll = arg.IsAtAll
			msg.At = &at
		}
	case MsgTypeLink:
		var msgMeta LinkMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.Link = &msgMeta
	case MsgTypeMarkdown:
		var msgMeta MarkdownMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.Markdown = &msgMeta

		var at AtMeta
		if arg.IsAtAll || arg.AtUser != "" || arg.AtMobile != "" {
			if arg.AtUser != "" {
				at.AtUserIds = strings.Split(arg.AtUser, "|")
			}
			if arg.AtMobile != "" {
				at.AtMobiles = strings.Split(arg.AtMobile, "|")
			}
			at.IsAtAll = arg.IsAtAll
			msg.At = &at
		}
	case MsgTypeSingleActionCard:
		var msgMeta SingleActionCardMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.MsgType = MsgTypeActionCard
		msg.ActionCard = &msgMeta
	case MsgTypeActionCard:
		var msgMeta ActionCardMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.ActionCard = &msgMeta
	case MsgTypeFeedCard:
		var msgMeta FeedCardMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.FeedCard = &msgMeta
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if err := Send(arg.AccessToken, arg.Secret, &msg); err != nil {
		return err
	}
	fmt.Println(dingtalk.MessageOK)

	return nil
}
