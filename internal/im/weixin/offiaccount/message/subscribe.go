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

/*
数据示例

{
  "touser": "OPENID",
  "template_id": "TEMPLATEID",
  "page": "mp.weixin.qq.com",
  "miniprogram": {
    "appid": "APPID",
    "pagepath": "index?foo=bar"
  },
  "data": {
    "name1": {
      "value": "广州腾讯科技有限公司"
    },
    "thing8": {
      "value": "广州腾讯科技有限公司"
    },
    "time7": {
      "value": "2019年8月8日"
    }
  }
}


{
  "errcode": 0,
  "errmsg": "ok"
}
*/

// SubscribeMessage 微信公众号订阅通知消息
type SubscribeMessage struct {
	ToUser      string                       `json:"touser"`
	TemplateID  string                       `json:"template_id"`
	Page        string                       `json:"page,omitempty"`
	MiniProgram *MiniProgramMeta             `json:"miniprogram,omitempty"` // 可选, 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        map[string]SubscribeDataItem `json:"data"`
}

// SubscribeDataItem 模板变量值
type SubscribeDataItem struct {
	Value string `json:"value"`
}

const subscribeURL = weixin.Host + "/cgi-bin/message/subscribe/bizsend?access_token="

// BizSendSubscribe 发送微信公众号订阅通知消息
func BizSendSubscribe(accessToken string, msg *SubscribeMessage) error {
	u := subscribeURL + url.QueryEscape(accessToken)
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
