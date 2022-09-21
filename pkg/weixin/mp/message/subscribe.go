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
	MiniProgram *MiniProgram                 `json:"miniprogram,omitempty"` // 可选, 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        map[string]SubscribeDataItem `json:"data"`
}

// SubscribeDataItem 模板变量值
type SubscribeDataItem struct {
	Value string `json:"value"`
}

const subscribeBizSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/bizsend?access_token=%s"

// BizSendSubscribe 发送微信公众号订阅通知消息
func BizSendSubscribe(accessToken string, msg *SubscribeMessage) error {
	url := fmt.Sprintf(subscribeBizSendURL, accessToken)
	var resp weixin.ResponseCode
	_, err := client.PostJSON(url, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; uri: %q, %v", weixin.ErrWeiXinRequest, url, resp)
	}
	return nil
}
