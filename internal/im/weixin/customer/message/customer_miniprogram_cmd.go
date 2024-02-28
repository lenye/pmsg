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
	"github.com/lenye/pmsg/internal/im/weixin/token"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdMiniSendCustomerParams struct {
	UserAgent   string
	AccessToken string
	AppID       string
	AppSecret   string
	ToUser      string
	MsgType     string
	Data        string
}

func (t *CmdMiniSendCustomerParams) Validate() error {
	if t.AccessToken == "" && t.AppID == "" {
		return flags.ErrWeixinAccessToken
	}
	if err := ValidateMiniProgramMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdMiniSendCustomer 发送微信小程序客服消息
func CmdMiniSendCustomer(arg *CmdMiniSendCustomerParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := CustomerMessage{
		ToUser:  arg.ToUser,
		MsgType: arg.MsgType,
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case MiniProgramMsgTypeText:
		var msgMeta TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case MiniProgramMsgTypeImage:
		var msgMeta ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case MiniProgramMsgTypeLink:
		var msgMeta LinkMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.Description == "" {
			return errors.New("description is empty")
		}
		if msgMeta.Url == "" {
			return errors.New("url is empty")
		}
		if msgMeta.ThumbUrl == "" {
			return errors.New("thumb_url is empty")
		}
		msg.Link = &msgMeta
	case MiniProgramMsgTypeMiniProgramPage:
		var msgMeta MiniProgramPageMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.AppID != "" {
			return errors.New("no appid required")
		}
		if msgMeta.PagePath == "" {
			return errors.New("pagepath is empty")
		}
		if msgMeta.ThumbMediaID == "" {
			return errors.New("thumb_media_id is empty")
		}
		msg.MiniProgramPage = &msgMeta
	}

	httpclient.SetUserAgent(arg.UserAgent)

	if arg.AccessToken == "" {
		accessTokenResp, err := token.FetchAccessToken(arg.AppID, arg.AppSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if err := SendCustomer(arg.AccessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
