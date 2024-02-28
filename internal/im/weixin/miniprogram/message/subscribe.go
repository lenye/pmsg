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
  "page": "index",
  "miniprogram_state": "developer",
  "lang": "zh_CN",
  "data": {
    "number01": {
      "value": "339208499"
    },
    "date01": {
      "value": "2015年01月05日"
    },
    "site01": {
      "value": "TIT创意园"
    },
    "site02": {
      "value": "广州市新港中路397号"
    }
  }
}

{
  "errcode": 0,
  "errmsg": "ok"
}
*/

// SubscribeMessage 微信小程序订阅消息
type SubscribeMessage struct {
	ToUser           string                       `json:"touser"`
	TemplateID       string                       `json:"template_id"`
	Page             string                       `json:"page,omitempty"`
	MiniProgramState string                       `json:"miniprogram_state,omitempty"`
	Data             map[string]SubscribeDataItem `json:"data"`
	Language         string                       `json:"lang,omitempty"`
}

// SubscribeDataItem 模板变量值
type SubscribeDataItem struct {
	Value string `json:"value"`
}

const (
	MiniProgramStateDeveloper = "developer" // developer为开发版
	MiniProgramStateTrial     = "trial"     // trial为体验版
	MiniProgramStateFormal    = "formal"    // formal为正式版；默认
)

// ValidateMiniProgramState 验证
func ValidateMiniProgramState(v string) error {
	switch v {
	case MiniProgramStateDeveloper, MiniProgramStateTrial, MiniProgramStateFormal:
	default:
		return fmt.Errorf("%s not in [%q %q %q]", v, MiniProgramStateDeveloper, MiniProgramStateTrial, MiniProgramStateFormal)
	}
	return nil
}

const (
	LanguageZhCN = "zh_CN" // 简体中文, 默认
	LanguageEnUS = "en_US" // 英文
	LanguageZhHK = "zh_HK" // 繁体中文
	LanguageZhTW = "zh_TW" // 繁体中文
)

// ValidateLanguage 验证
func ValidateLanguage(v string) error {
	switch v {
	case LanguageZhCN, LanguageEnUS, LanguageZhHK, LanguageZhTW:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q]", v, LanguageZhCN, LanguageEnUS, LanguageZhHK, LanguageZhTW)
	}
	return nil
}

const reqURL = weixin.Host + "/cgi-bin/message/subscribe/send?access_token="

// SendSubscribe 发送微信小程序订阅消息
func SendSubscribe(accessToken string, msg *SubscribeMessage) error {
	u := reqURL + url.QueryEscape(accessToken)
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
