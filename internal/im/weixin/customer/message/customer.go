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
	"fmt"
	"net/url"

	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/internal/im/weixin/client"
)

// CustomerMessage 微信客服消息
type CustomerMessage struct {
	ToUser          string               `json:"touser"`
	MsgType         string               `json:"msgtype"`
	CustomService   *ServiceMeta         `json:"customservice,omitempty"`
	Text            *TextMeta            `json:"text,omitempty"`
	Image           *ImageMeta           `json:"image,omitempty"`
	Voice           *VoiceMeta           `json:"voice,omitempty"`
	Video           *VideoMeta           `json:"video,omitempty"`
	Music           *MusicMeta           `json:"music,omitempty"`
	News            *NewsMeta            `json:"news,omitempty"`
	MpNews          *MpNewsMeta          `json:"mpnews,omitempty"`
	MpNewsArticle   *MpNewsArticleMeta   `json:"mpNewsArticle,omitempty"`
	MsgMenu         *MsgMenuMeta         `json:"msgmenu,omitempty"`
	WxCard          *WxCardMeta          `json:"wxcard,omitempty"`
	MiniProgramPage *MiniProgramPageMeta `json:"miniprogrampage,omitempty"`
	Link            *LinkMeta            `json:"link,omitempty"`
}

const reqURL = weixin.Host + "/cgi-bin/message/custom/send?access_token="

// SendCustomer 发送微信客服消息
func SendCustomer(accessToken string, msg *CustomerMessage) error {
	u := reqURL + url.QueryEscape(accessToken)
	var resp weixin.ResponseMeta
	_, err := client.PostJSON(u, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; %v", weixin.ErrRequest, resp)
	}
	return nil
}
