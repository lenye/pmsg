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

package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/lenye/pmsg/pkg/file"
	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

type CmdSendParams struct {
	UserAgent string
	Key       string
	MsgType   string
	ToUser    string
	ToMobile  string
	Data      string
}

func (t *CmdSendParams) Validate() error {
	if err := ValidateMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	if t.MsgType == MsgTypeImage {
		if !file.Exists(t.Data) {
			return fmt.Errorf("file is not exist, %v", t.Data)
		}
	}

	return nil
}

// CmdSend 发送企业微信群机器人消息
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
		if arg.ToUser != "" {
			msgMeta.MentionedList = strings.Split(arg.ToUser, "|")
		}
		if arg.ToMobile != "" {
			msgMeta.MentionedMobileList = strings.Split(arg.ToMobile, "|")
		}
		msg.Text = &msgMeta
	case MsgTypeMarkdown:
		var msgMeta MarkdownMeta
		msgMeta.Content = buf.String()
		msg.Markdown = &msgMeta
	case MsgTypeImage:
		msgMeta, err := ImageFile2Meta(buf.String())
		if err != nil {
			return err
		}
		msg.Image = msgMeta
	case MsgTypeNews:
		var msgMeta NewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.News = &msgMeta
	case MsgTypeFile:
		var msgMeta FileMeta
		msgMeta.MediaID = buf.String()
		msg.File = &msgMeta
	case TplCardTypeText:
		var msgMeta TemplateCardText
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.CardType != TplCardTypeText {
			return errors.New("invalid card_type")
		}
		msg.MsgType = MsgTypeTplCard
		msg.TemplateCard = &msgMeta
	case TplCardTypeNews:
		var msgMeta TemplateCardNews
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.CardType != TplCardTypeNews {
			return errors.New("invalid card_type")
		}
		msg.MsgType = MsgTypeTplCard
		msg.TemplateCard = &msgMeta
	}

	client.UserAgent = arg.UserAgent

	if err := Send(arg.Key, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
