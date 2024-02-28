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
	"github.com/lenye/pmsg/internal/im/weixin/work"
)

// 企业微信群聊推送消息 类型
const (
	AppChatMsgTypeText     = "text"     // 文本消息
	AppChatMsgTypeImage    = "image"    // 图片消息
	AppChatMsgTypeVoice    = "voice"    // 语音消息
	AppChatMsgTypeVideo    = "video"    // 视频消息
	AppChatMsgTypeFile     = "file"     // 文件消息
	AppChatMsgTypeTextCard = "textcard" // 文本卡片消息
	AppChatMsgTypeNews     = "news"     // 图文消息
	AppChatMsgTypeMpNews   = "mpnews"   // 图文消息
	AppChatMsgTypeMarkdown = "markdown" // markdown消息
)

// ValidateAppChatMsgType 验证
func ValidateAppChatMsgType(v string) error {
	switch v {
	case AppChatMsgTypeText, AppChatMsgTypeImage, AppChatMsgTypeVoice, AppChatMsgTypeVideo,
		AppChatMsgTypeFile, AppChatMsgTypeTextCard, AppChatMsgTypeNews, AppChatMsgTypeMpNews,
		AppChatMsgTypeMarkdown:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q %q %q]", v,
			AppChatMsgTypeText, AppChatMsgTypeImage, AppChatMsgTypeVoice, AppChatMsgTypeVideo,
			AppChatMsgTypeFile, AppChatMsgTypeTextCard, AppChatMsgTypeNews, AppChatMsgTypeMpNews,
			AppChatMsgTypeMarkdown)
	}
	return nil
}

// AppChatMessage 企业微信群聊推送消息 touser、toparty、totag不能同时为空
type AppChatMessage struct {
	ChatID   string        `json:"chatid"`             // 群聊id
	MsgType  string        `json:"msgtype"`            // 消息类型
	Safe     int           `json:"safe,omitempty"`     // 表示是否是保密消息，0表示否，1表示是，默认0
	Text     *TextMeta     `json:"text,omitempty"`     // 文本消息
	Image    *ImageMeta    `json:"image,omitempty"`    // 图片消息
	Voice    *VoiceMeta    `json:"voice,omitempty"`    // 语音消息
	Video    *VideoMeta    `json:"video,omitempty"`    // 视频消息
	File     *FileMeta     `json:"file,omitempty"`     // 文件消息
	TextCard *TextCardMeta `json:"textcard,omitempty"` // 文本卡片消息
	News     *NewsMeta     `json:"news,omitempty"`     // 图文消息
	MpNews   *MpNewsMeta   `json:"mpnews,omitempty"`   // 图文消息
	Markdown *MarkdownMeta `json:"markdown,omitempty"` // markdown消息
}

const appChatSendURL = work.Host + "/cgi-bin/appchat/send?access_token="

// SendAppChat 发送企业微信群聊推送消息
func SendAppChat(accessToken string, msg *AppChatMessage) error {
	u := appChatSendURL + url.QueryEscape(accessToken)
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
