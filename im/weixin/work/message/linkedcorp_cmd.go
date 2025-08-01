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

	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/httpclient"
	"github.com/lenye/pmsg/im/weixin"
	"github.com/lenye/pmsg/im/weixin/work"
	"github.com/lenye/pmsg/im/weixin/work/token"
)

type CmdWorkSendLinkedCorpParams struct {
	UserAgent   string
	AccessToken string
	CorpID      string
	CorpSecret  string
	ToUser      []string
	ToParty     []string
	ToTag       []string
	ToAll       int
	AgentID     int64
	MsgType     string
	Safe        int
	Data        string
}

func (t *CmdWorkSendLinkedCorpParams) Validate() error {
	if t.AccessToken == "" && t.CorpID == "" {
		return flags.ErrWeixinWorkAccessToken
	}

	if t.ToUser != nil && len(t.ToUser) > 1000 {
		return fmt.Errorf("%v supports up to 1000", flags.ToUser)
	}
	if t.ToParty != nil && len(t.ToParty) > 100 {
		return fmt.Errorf("%v supports up to 100", flags.ToParty)
	}
	if t.ToTag != nil && len(t.ToTag) > 100 {
		return fmt.Errorf("%v supports up to 100", flags.ToTag)
	}

	if t.ToAll != 0 && t.ToAll != 1 {
		return fmt.Errorf("invalid %v", flags.ToAll)
	}
	if t.Safe != 0 && t.Safe != 1 {
		return fmt.Errorf("invalid %v", flags.Safe)
	}

	if err := ValidateLinkedCorpMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdWorkSendLinkedCorp 发送企业微信互联企业消息
func CmdWorkSendLinkedCorp(arg *CmdWorkSendLinkedCorpParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := LinkedCorpMessage{
		ToUser:  arg.ToUser,
		ToParty: arg.ToParty,
		ToTag:   arg.ToTag,
		ToAll:   arg.ToAll,
		AgentID: arg.AgentID,
		MsgType: arg.MsgType,
		Safe:    arg.Safe,
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case LinkedCorpMsgTypeText:
		var msgMeta TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case LinkedCorpMsgTypeImage:
		var msgMeta ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case LinkedCorpMsgTypeVoice:
		var msgMeta VoiceMeta
		msgMeta.MediaID = buf.String()
		msg.Voice = &msgMeta
	case LinkedCorpMsgTypeVideo:
		var msgMeta VideoMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MediaID == "" {
			return errors.New("media_id is empty")
		}
		msg.Video = &msgMeta
	case LinkedCorpMsgTypeFile:
		var msgMeta FileMeta
		msgMeta.MediaID = buf.String()
		msg.File = &msgMeta
	case LinkedCorpMsgTypeTextCard:
		var msgMeta TextCardMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.TextCard = &msgMeta
	case LinkedCorpMsgTypeNews:
		var msgMeta NewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.News = &msgMeta
	case LinkedCorpMsgTypeMpNews:
		var msgMeta MpNewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.MpNews = &msgMeta
	case LinkedCorpMsgTypeMarkdown:
		var msgMeta MarkdownMeta
		msgMeta.Content = buf.String()
		msg.Markdown = &msgMeta
	case LinkedCorpMsgTypeMiniProgramNotice:
		var msgMeta MiniProgramNoticeMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.ContentItem)
		if lenArticles > 10 {
			return errors.New("content_item up to 10")
		}
		msg.MiniProgramNotice = &msgMeta
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if arg.AccessToken == "" {
		accessTokenResp, err := token.FetchAccessToken(arg.CorpID, arg.CorpSecret)
		if err != nil {
			return fmt.Errorf("%w, %w", work.ErrRequest, err)
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if resp, err := SendLinkedCorp(arg.AccessToken, &msg); err != nil {
		return fmt.Errorf("%w, %w", work.ErrRequest, err)
	} else {
		fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, resp))
	}

	return nil
}
