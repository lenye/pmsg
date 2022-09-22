package message

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
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

const subscribeSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"

// SendSubscribe 发送微信小程序订阅消息
func SendSubscribe(accessToken string, msg *SubscribeMessage) error {
	url := fmt.Sprintf(subscribeSendURL, accessToken)
	var resp weixin.ResponseMeta
	_, err := client.PostJSON(url, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; %v", weixin.ErrWeiXinRequest, resp)
	}
	return nil
}
