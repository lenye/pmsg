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
	"strconv"
	"time"

	"github.com/lenye/pmsg/internal/im/dingtalk"
	"github.com/lenye/pmsg/internal/im/dingtalk/client"
)

const (
	MsgTypeText             = "text"              // 文本
	MsgTypeLink             = "link"              // 链接
	MsgTypeMarkdown         = "markdown"          // markdown
	MsgTypeFeedCard         = "feedCard"          // FeedCard
	MsgTypeActionCard       = "actionCard"        // ActionCard
	MsgTypeSingleActionCard = "single_actionCard" // Single ActionCard
)

// ValidateMsgType 验证
func ValidateMsgType(v string) error {
	switch v {
	case MsgTypeText, MsgTypeLink, MsgTypeMarkdown, MsgTypeActionCard, MsgTypeSingleActionCard, MsgTypeFeedCard:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q]", v,
			MsgTypeText, MsgTypeLink, MsgTypeMarkdown, MsgTypeActionCard, MsgTypeSingleActionCard, MsgTypeFeedCard)
	}
	return nil
}

// Message 钉钉自定义机器人消息
type Message struct {
	MsgType    string        `json:"msgtype"`              // 消息类型
	Text       *TextMeta     `json:"text,omitempty"`       // 文本消息
	Markdown   *MarkdownMeta `json:"markdown,omitempty"`   // markdown消息
	At         *AtMeta       `json:"at,omitempty"`         // @
	Link       *LinkMeta     `json:"link,omitempty"`       // 链接
	ActionCard any           `json:"actionCard,omitempty"` // ActionCard
	FeedCard   *FeedCardMeta `json:"feedCard,omitempty"`   // FeedCard
}

const sendURL = "https://oapi.dingtalk.com/robot/send?access_token="

// Send 发送钉钉自定义机器人消息
//
// 消息发送频率限制
// 每个机器人每分钟最多发送20条消息到群里，如果超过20条，会限流10分钟
func Send(accessToken, secret string, msg *Message) error {
	u := sendURL + url.QueryEscape(accessToken)
	if secret != "" {
		timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
		sign, err := dingtalk.Sign(timestamp, secret)
		if err != nil {
			return fmt.Errorf("sign failed: %w", err)
		}
		u = u + "&timestamp=" + timestamp + "&sign=" + url.QueryEscape(sign)
	}
	var resp dingtalk.ResponseMeta
	_, err := client.PostJSON(u, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; %v", dingtalk.ErrRequest, resp)
	}
	return nil
}
