package message

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
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

const undoAppSendURL = "https://qyapi.weixin.qq.com/cgi-bin/message/recall?access_token="

// UndoApp 撤回企业微信应用消息
func UndoApp(accessToken string, msg *UndoAppMessage) error {
	url := appSendURL + accessToken
	var resp UndoAppMessageResponse
	_, err := client.PostJSON(url, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; %v", weixin.ErrWeiXinRequest, resp)
	}
	return nil
}
