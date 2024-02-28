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
	"errors"
	"fmt"
	"strings"

	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/internal/im/weixin/work/token"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdWorkSendAppParams struct {
	UserAgent              string
	AccessToken            string
	CorpID                 string
	CorpSecret             string
	ToUser                 string
	ToParty                string
	ToTag                  string
	AgentID                int64
	MsgType                string
	Safe                   int
	EnableIDTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
	Data                   string
}

func (t *CmdWorkSendAppParams) Validate() error {
	if t.AccessToken == "" && t.CorpID == "" {
		return flags.ErrWeixinWorkAccessToken
	}

	if t.ToUser == "" && t.ToParty == "" && t.ToTag == "" {
		return fmt.Errorf("%v、%v、%v cannot be empty at the same time", flags.ToUser, flags.ToParty, flags.ToTag)
	}

	if t.ToUser == flags.AllUser {
		t.ToParty = ""
		t.ToTag = ""
	} else {
		if toUsers := strings.Split(t.ToUser, "|"); len(toUsers) > 1000 {
			return fmt.Errorf("%v supports up to 1000", flags.ToUser)
		}
		if toPartys := strings.Split(t.ToParty, "|"); len(toPartys) > 100 {
			return fmt.Errorf("%v supports up to 100", flags.ToParty)
		}

		if toTags := strings.Split(t.ToTag, "|"); len(toTags) > 100 {
			return fmt.Errorf("%v supports up to 100", flags.ToTag)
		}
	}

	if t.Safe != 0 && t.Safe != 1 {
		return fmt.Errorf("invalid %v", flags.Safe)
	}

	if t.EnableIDTrans != 0 && t.EnableIDTrans != 1 {
		return fmt.Errorf("invalid %v", flags.EnableIDTrans)
	}

	if t.EnableDuplicateCheck != 0 && t.EnableDuplicateCheck != 1 {
		return fmt.Errorf("invalid %v", flags.EnableDuplicateCheck)
	}

	if t.DuplicateCheckInterval <= 0 || t.DuplicateCheckInterval > 3600*4 {
		return fmt.Errorf("invalid %v", flags.DuplicateCheckInterval)
	}

	if err := ValidateAppMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdWorkSendApp 发送企业微信应用消息
func CmdWorkSendApp(arg *CmdWorkSendAppParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := AppMessage{
		ToUser:                 arg.ToUser,
		ToParty:                arg.ToParty,
		ToTag:                  arg.ToTag,
		AgentID:                arg.AgentID,
		MsgType:                arg.MsgType,
		Safe:                   arg.Safe,
		EnableIDTrans:          arg.EnableIDTrans,
		EnableDuplicateCheck:   arg.EnableDuplicateCheck,
		DuplicateCheckInterval: arg.DuplicateCheckInterval,
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case AppMsgTypeText:
		var msgMeta TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case AppMsgTypeImage:
		var msgMeta ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case AppMsgTypeVoice:
		var msgMeta VoiceMeta
		msgMeta.MediaID = buf.String()
		msg.Voice = &msgMeta
	case AppMsgTypeVideo:
		var msgMeta VideoMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MediaID == "" {
			return errors.New("media_id is empty")
		}
		msg.Video = &msgMeta
	case AppMsgTypeFile:
		var msgMeta FileMeta
		msgMeta.MediaID = buf.String()
		msg.File = &msgMeta
	case AppMsgTypeTextCard:
		var msgMeta TextCardMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.TextCard = &msgMeta
	case AppMsgTypeNews:
		var msgMeta NewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.News = &msgMeta
	case AppMsgTypeMpNews:
		var msgMeta MpNewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.MpNews = &msgMeta
	case AppMsgTypeMarkdown:
		var msgMeta MarkdownMeta
		msgMeta.Content = buf.String()
		msg.Markdown = &msgMeta
	case AppMsgTypeMiniProgramNotice:
		var msgMeta MiniProgramNoticeMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.ContentItem)
		if lenArticles > 10 {
			return errors.New("content_item up to 10")
		}
		msg.MiniProgramNotice = &msgMeta
	case AppMsgTypeTemplateCard:
		var msgMeta TemplateCardMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if err := ValidateAppTemplateCardType(msgMeta.CardType); err != nil {
			return fmt.Errorf("invalid card_type: %v", err)
		}
		msg.TemplateCard = &msgMeta
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if arg.AccessToken == "" {
		accessTokenResp, err := token.FetchAccessToken(arg.CorpID, arg.CorpSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if resp, err := SendApp(arg.AccessToken, &msg); err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, resp))
	}

	return nil
}
