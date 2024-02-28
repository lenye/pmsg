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
  "template_id": "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
  "url": "http://weixin.qq.com/download",
  "miniprogram": {
    "appid": "xiaochengxuappid12345",
    "pagepath": "index?foo=bar"
  },
  "client_msg_id": "MSG_000001",
  "data": {
    "first": {
      "value": "恭喜你购买成功！",
      "color": "#173177"
    },
    "keyword1": {
      "value": "巧克力",
      "color": "#173177"
    },
    "keyword2": {
      "value": "39.8元",
      "color": "#173177"
    },
    "keyword3": {
      "value": "2014年9月22日",
      "color": "#173177"
    },
    "remark": {
      "value": "欢迎再次购买！",
      "color": "#173177"
    }
  }
}

{
  "errcode": 0,
  "errmsg": "ok",
  "msgid": 200228332
}
*/

// TemplateMessage 微信公众号模板消息
type TemplateMessage struct {
	ClientMsgID string                      `json:"client_msg_id,omitempty"` // 可选, 防重入id。对于同一个openid + client_msg_id, 只发送一条消息,10分钟有效,超过10分钟不保证效果。若无防重入需求，可不填
	ToUser      string                      `json:"touser"`                  // 必须, 接受者OpenID
	TemplateID  string                      `json:"template_id"`             // 必须, 模版ID
	URL         string                      `json:"url,omitempty"`           // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	MiniProgram *MiniProgramMeta            `json:"miniprogram,omitempty"`   // 可选, 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        map[string]TemplateDataItem `json:"data"`                    // 必须, 模板数据, JSON 格式的 []byte, 满足特定的模板需求
	Color       string                      `json:"color,omitempty"`         // 可选, 模板内容字体颜色，不填默认为黑色
}

/*
url 和 miniprogram 都是非必填字段，若都不传则模板无跳转；
若都传，会优先跳转至小程序。
开发者可根据实际需要选择其中一种跳转方式即可。
当用户的微信客户端版本不支持跳小程序时，将会跳转至url。
*/

// MiniProgramMeta 跳小程序所需数据
type MiniProgramMeta struct {
	AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序 appid 必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

// TemplateDataItem 模板变量值, 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`           // 必选, 模板内容
	Color string `json:"color,omitempty"` // 可选, 模板内容字体颜色，不填默认为黑色
}

// TemplateMessageResponse 模板消息的响应
type TemplateMessageResponse struct {
	weixin.ResponseMeta
	MsgID int64 `json:"msgid,omitempty"` // 消息id
}

func (t TemplateMessageResponse) String() string {
	return fmt.Sprintf("errcode: %v, errmsg: %q, msgid: %v", t.ErrorCode, t.ErrorMessage, t.MsgID)
}

const templateURL = weixin.Host + "/cgi-bin/message/template/send?access_token="

// SendTemplate 发送微信公众号模板消息
func SendTemplate(accessToken string, msg *TemplateMessage) (int64, error) {
	u := templateURL + url.QueryEscape(accessToken)
	var resp TemplateMessageResponse
	_, err := client.PostJSON(u, msg, &resp)
	if err != nil {
		return 0, err
	}
	if !resp.Succeed() {
		return 0, fmt.Errorf("%w; %v", weixin.ErrRequest, resp.ResponseMeta)
	}
	return resp.MsgID, nil
}
