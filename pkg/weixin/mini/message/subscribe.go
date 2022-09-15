package message

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

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

const (
	LanguageZhCN = "zh_CN" // 简体中文, 默认
	LanguageEnUS = "en_US" // 英文
	LanguageZhHK = "zh_HK" // 繁体中文
	LanguageZhTW = "zh_TW" // 繁体中文
)

const subscribeSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"

// SendSubscribe 发送微信小程序订阅消息
func SendSubscribe(accessToken string, msg *SubscribeMessage) (err error) {
	url := fmt.Sprintf(subscribeSendURL, accessToken)
	var resp weixin.ResponseCode
	_, err = client.PostJSON(url, msg, &resp)
	if err != nil {
		return
	}
	if !resp.Succeed() {
		err = fmt.Errorf("weixin request failed, uri=%q, resp=%+v", url, resp)
	}
	return
}
