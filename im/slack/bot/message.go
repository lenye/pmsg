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
	"github.com/lenye/pmsg/im/slack/client"
)

// Send 发送消息
//
// 消息发送频率限制
// 1 per second，Short bursts >1 allowed. 每秒1次
func Send(webhookUrl, body string) error {
	_, err := client.PostJSON(webhookUrl, body)
	if err != nil {
		return err
	}
	return nil
}
