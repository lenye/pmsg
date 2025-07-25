// Copyright 2022-2025 The pmsg Authors. All rights reserved.
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
	"strings"

	"github.com/lenye/pmsg/httpclient"
	"github.com/lenye/pmsg/im/discord"
	"github.com/lenye/pmsg/im/discord/client"
)

// Message Discord机器人消息
type Message struct {
	Content string `json:"content"` // 文本消息
}

// Send 发送消息
//
// 消息发送频率限制
// - 每个 Webhook 最多可以每 2 秒发送 5 个请求
// - 每 2 秒最多创建或删除 5 次 Webhook
// - 每个频道每分钟最多发送 30 次请求
func Send(webhookUrl string, msg *Message) error {
	var resp discord.ResponseMeta
	headers, err := client.PostJSON(webhookUrl, msg, &resp)
	if err != nil {
		return err
	}
	if strings.EqualFold(headers.Get("content-type"), httpclient.HdrValApplicationJson) {
		if !resp.Succeed() {
			return fmt.Errorf("%w, %s", discord.ErrRequest, resp.String())
		}
	}
	return nil
}
