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
	MiniProgram *MiniProgram                `json:"miniprogram,omitempty"` // 可选, 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        map[string]TemplateDataItem `json:"data"`                  // 必须, 模板数据, JSON 格式的 []byte, 满足特定的模板需求
}

const subscribeTemplateSendURL = "https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=%s"

// SendTemplateSubscribe 发送微信公众号一次性订阅消息
func SendTemplateSubscribe(accessToken string, msg *TemplateSubscribeMessage) error {
	url := fmt.Sprintf(subscribeTemplateSendURL, accessToken)
	var resp weixin.ResponseCode
	_, err := client.PostJSON(url, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; uri: %q, response: %v", weixin.ErrWeiXinRequest, url, resp)
	}
	return nil
}
