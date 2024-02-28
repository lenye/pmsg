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

type CmdMpSendCustomerParams struct {
	UserAgent   string
	AccessToken string
	AppID       string
	AppSecret   string
	ToUser      string
	MsgType     string
	KfAccount   string
	Data        string
}

func (t *CmdMpSendCustomerParams) Validate() error {
	if t.AccessToken == "" && t.AppID == "" {
		return flags.ErrWeixinAccessToken
	}
	if err := ValidateMpMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdMpSendCustomer 发送微信公众号客服消息
func CmdMpSendCustomer(arg *CmdMpSendCustomerParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := CustomerMessage{
		ToUser:  arg.ToUser,
		MsgType: arg.MsgType,
	}

	if arg.KfAccount != "" {
		msg.CustomService = &ServiceMeta{KfAccount: arg.KfAccount}
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case MpMsgTypeText:
		var msgMeta TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case MpMsgTypeImage:
		var msgMeta ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case MpMsgTypeVoice:
		var msgMeta VoiceMeta
		msgMeta.MediaID = buf.String()
		msg.Voice = &msgMeta
	case MpMsgTypeVideo:
		var msgMeta VideoMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MediaID == "" {
			return errors.New("media_id is empty")
		}
		if msgMeta.ThumbMediaID == "" {
			return errors.New("thumb_media_id is empty")
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.Description == "" {
			return errors.New("description is empty")
		}
		msg.Video = &msgMeta
	case MpMsgTypeMusic:
		var msgMeta MusicMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MusicUrl == "" {
			return errors.New("musicurl is empty")
		}
		if msgMeta.HqmusicUrl == "" {
			return errors.New("hqmusicurl is empty")
		}
		if msgMeta.ThumbMediaID == "" {
			return errors.New("thumb_media_id is empty")
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.Description == "" {
			return errors.New("description is empty")
		}
		msg.Music = &msgMeta
	case MpMsgTypeNews:
		var msgMeta NewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if len(msgMeta.Articles) != 1 {
			return errors.New("articles length != 1")
		}
		article := msgMeta.Articles[0]
		if article.Url == "" {
			return errors.New("url is empty")
		}
		if article.PicUrl == "" {
			return errors.New("picurl is empty")
		}
		if article.Title == "" {
			return errors.New("title is empty")
		}
		if article.Description == "" {
			return errors.New("description is empty")
		}
		msg.News = &msgMeta
	case MpMsgTypeMpNews:
		var msgMeta MpNewsMeta
		msgMeta.MediaID = buf.String()
		msg.MpNews = &msgMeta
	case MpMsgTypeMpNewsArticle:
		var msgMeta MpNewsArticleMeta
		msgMeta.ArticleID = buf.String()
		msg.MpNewsArticle = &msgMeta
	case MpMsgTypeMsgMenu:
		var msgMeta MsgMenuMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.HeadContent == "" {
			return errors.New("head_content is empty")
		}
		if msgMeta.TailContent == "" {
			return errors.New("tail_content is empty")
		}
		lenList := len(msgMeta.List)
		if lenList == 0 {
			return errors.New("list is empty")
		}
		for i := 0; i < lenList; i++ {
			if msgMeta.List[i].ID == "" {
				return fmt.Errorf("list[%v].id is empty", i)
			}
			if msgMeta.List[i].Content == "" {
				return fmt.Errorf("list[%v].content is empty", i)
			}
		}
		msg.MsgMenu = &msgMeta
	case MpMsgTypeWxCard:
		var msgMeta WxCardMeta
		msgMeta.CardID = buf.String()
		msg.WxCard = &msgMeta
	case MiniProgramMsgTypeMiniProgramPage:
		var msgMeta MiniProgramPageMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.AppID == "" {
			return errors.New("appid is empty")
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
