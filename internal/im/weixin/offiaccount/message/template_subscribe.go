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
  "template_id": "TEMPLATE_ID",
  "url": "URL",
  "miniprogram": {
    "appid": "xiaochengxuappid12345",
    "pagepath": "index?foo=bar"
  },
  "scene": "SCENE",
  "title": "TITLE",
  "data": {
    "content": {
      "value": "VALUE",
      "color": "COLOR"
    }
  }
}

{
  "errcode": 0,
  "errmsg": "ok"
}
*/

// TemplateSubscribeMessage 微信公众号一次性订阅消息
type TemplateSubscribeMessage struct {
	ToUser      string                      `json:"touser"`                // 必须, 接受者OpenID
	TemplateID  string                      `json:"template_id"`           // 必须, 模版ID
	Scene       string                      `json:"scene"`                 // 订阅场景值
	Title       string                      `json:"title"`                 // 消息标题，15字以内
	URL         string                      `json:"url,omitempty"`         // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	MiniProgram *MiniProgramMeta            `json:"miniprogram,omitempty"` // 可选, 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        map[string]TemplateDataItem `json:"data"`                  // 必须, 模板数据, JSON 格式的 []byte, 满足特定的模板需求
}

const subscribeTplURL = weixin.Host + "/cgi-bin/message/template/subscribe?access_token="

// SendTemplateSubscribe 发送微信公众号一次性订阅消息
func SendTemplateSubscribe(accessToken string, msg *TemplateSubscribeMessage) error {
	u := subscribeTplURL + url.QueryEscape(accessToken)
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
