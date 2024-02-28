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

// UndoAppMessage 撤回企业微信应用消息
// 可以撤回24小时内通过发送应用消息接口推送的消息，仅可撤回企业微信端的数据，微信插件端的数据不支持撤回。
type UndoAppMessage struct {
	MsgID string `json:"msgid"` // 消息ID。从应用发送消息接口处获得
}

// UndoAppMessageResponse 企业微信撤回应用消息响应
type UndoAppMessageResponse struct {
	weixin.ResponseMeta
}

const undoAppSendURL = work.Host + "/cgi-bin/message/recall?access_token="

// UndoApp 撤回企业微信应用消息
func UndoApp(accessToken string, msg *UndoAppMessage) error {
	u := appSendURL + url.QueryEscape(accessToken)
	var resp UndoAppMessageResponse
	_, err := client.PostJSON(u, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; %v", weixin.ErrRequest, resp)
	}
	return nil
}
