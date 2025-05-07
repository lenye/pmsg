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

	"github.com/lenye/pmsg/im/feishu"
	"github.com/lenye/pmsg/im/feishu/client"
)

const (
	MsgTypeText        = "text"        // 文本
	MsgTypePost        = "post"        // 富文本
	MsgTypeImage       = "image"       // 图片
	MsgTypeShareChat   = "share_chat"  // 分享群名片
	MsgTypeInteractive = "interactive" // 消息卡片
)

// ValidateMsgType 验证
func ValidateMsgType(v string) error {
	switch v {
	case MsgTypeText, MsgTypePost, MsgTypeImage, MsgTypeShareChat, MsgTypeInteractive:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q]", v,
			MsgTypeText, MsgTypePost, MsgTypeImage, MsgTypeShareChat, MsgTypeInteractive)
	}
	return nil
}

// Message 飞书自定义机器人消息
type Message struct {
	MsgType   string       `json:"msg_type"`            // 消息类型
	TimeStamp string       `json:"timestamp,omitempty"` // 为距当前时间不超过 1 小时(3600)的时间戳，时间单位s
	Sign      string       `json:"sign,omitempty"`      // 签名
	Content   *ContentMeta `json:"content,omitempty"`   // 消息内容
	Card      *CardMeta    `json:"card,omitempty"`      // 消息卡片
}

const sendURL = "https://open.feishu.cn/open-apis/bot/v2/hook/"

// Send 发送飞书自定义机器人消息
//
// 消息发送频率限制
// 自定义机器人的频率控制和普通应用不同，为单租户单机器人 100 次/分钟，5 次/秒。
// 建议发送消息尽量避开诸如 10:00、17:30 等整点及半点时间，否则可能出现因系统压力导致的 11232 限流错误，导致消息发送失败。
// 发送消息时，请求体的数据大小不能超过 20 KB。
func Send(accessToken string, msg *Message) error {
	u := sendURL + accessToken
	var resp feishu.ResponseMeta
	_, err := client.PostJSON(u, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; %v", feishu.ErrRequest, resp)
	}
	return nil
}
