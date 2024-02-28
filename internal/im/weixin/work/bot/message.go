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
	"fmt"
	"net/url"

	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/internal/im/weixin/client"
	"github.com/lenye/pmsg/internal/im/weixin/work"
)

const (
	MsgTypeText     = "text"          // 文本
	MsgTypeMarkdown = "markdown"      // markdown
	MsgTypeImage    = "image"         // 图片
	MsgTypeNews     = "news"          // 图文
	MsgTypeFile     = "file"          // 文件
	MsgTypeTplCard  = "template_card" // 模版卡片

	TplCardTypeText = "text_notice" // 文本通知模版卡片
	TplCardTypeNews = "news_notice" // 图文展示模版卡片
)

// ValidateMsgType 验证
func ValidateMsgType(v string) error {
	switch v {
	case MsgTypeText, MsgTypeMarkdown, MsgTypeImage, MsgTypeNews, MsgTypeFile, TplCardTypeText, TplCardTypeNews:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q]", v,
			MsgTypeText, MsgTypeMarkdown, MsgTypeImage, MsgTypeNews, MsgTypeFile, TplCardTypeText, TplCardTypeNews)
	}
	return nil
}

// Message 企业微信群机器人消息
type Message struct {
	MsgType      string        `json:"msgtype"`                 // 消息类型
	Text         *TextMeta     `json:"text,omitempty"`          // 文本消息
	Markdown     *MarkdownMeta `json:"markdown,omitempty"`      // markdown消息
	Image        *ImageMeta    `json:"image,omitempty"`         // 图片消息
	News         *NewsMeta     `json:"news,omitempty"`          // 图文消息
	File         *FileMeta     `json:"file,omitempty"`          // 文件消息
	TemplateCard any           `json:"template_card,omitempty"` // 模版卡片
}

const sendURL = work.Host + "/cgi-bin/webhook/send?key="

// Send 发送企业微信群机器人消息
//
// 消息发送频率限制
// 每个机器人发送的消息不能超过20条/分钟。
func Send(key string, msg *Message) error {
	u := sendURL + url.QueryEscape(key)
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
