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

	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/internal/im/weixin/work/token"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdWorkSendExternalContactParams struct {
	UserAgent              string
	AccessToken            string
	CorpID                 string
	CorpSecret             string
	RecvScope              int
	ToParentUserID         []string
	ToStudentUserID        []string
	ToParty                []string
	ToAll                  int
	MsgType                string
	AgentID                int64
	EnableIDTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
	Data                   string
}

func (t *CmdWorkSendExternalContactParams) Validate() error {
	if t.AccessToken == "" && t.CorpID == "" {
		return flags.ErrWeixinWorkAccessToken
	}

	if t.ToParentUserID != nil && len(t.ToParentUserID) > 1000 {
		return fmt.Errorf("%v supports up to 1000", flags.ToParentUserID)
	}
	if t.ToStudentUserID != nil && len(t.ToStudentUserID) > 1000 {
		return fmt.Errorf("%v supports up to 1000", flags.ToStudentUserID)
	}
	if t.ToParty != nil && len(t.ToParty) > 100 {
		return fmt.Errorf("%v supports up to 100", flags.ToParty)
	}

	if t.RecvScope != 0 && t.RecvScope != 1 && t.RecvScope != 2 {
		return fmt.Errorf("invalid %v", flags.RecvScope)
	}
	if t.ToAll != 0 && t.ToAll != 1 {
		return fmt.Errorf("invalid %v", flags.ToAll)
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

	if err := ValidateExternalContactMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdWorkSendExternalContact 发送企业微信互联企业消息
func CmdWorkSendExternalContact(arg *CmdWorkSendExternalContactParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := ExternalContactMessage{
		RecvScope:              arg.RecvScope,
		ToParentUserID:         arg.ToParentUserID,
		ToStudentUserID:        arg.ToStudentUserID,
		ToParty:                arg.ToParty,
		ToAll:                  arg.ToAll,
		MsgType:                arg.MsgType,
		AgentID:                arg.AgentID,
		EnableIDTrans:          arg.EnableIDTrans,
		EnableDuplicateCheck:   arg.EnableDuplicateCheck,
		DuplicateCheckInterval: arg.DuplicateCheckInterval,
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case ExternalContactMsgTypeText:
		var msgMeta TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case ExternalContactMsgTypeImage:
		var msgMeta ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case ExternalContactMsgTypeVoice:
		var msgMeta VoiceMeta
		msgMeta.MediaID = buf.String()
		msg.Voice = &msgMeta
	case ExternalContactMsgTypeVideo:
		var msgMeta VideoMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MediaID == "" {
			return errors.New("media_id is empty")
		}
		msg.Video = &msgMeta
	case ExternalContactMsgTypeFile:
		var msgMeta FileMeta
		msgMeta.MediaID = buf.String()
		msg.File = &msgMeta
	case ExternalContactMsgTypeNews:
		var msgMeta NewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.News = &msgMeta
	case ExternalContactMsgTypeMpNews:
		var msgMeta MpNewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.MpNews = &msgMeta
	case ExternalContactMsgTypeMiniProgramNotice:
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
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if resp, err := SendExternalContact(arg.AccessToken, &msg); err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, resp))
	}

	return nil
}
