package message

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

// CustomerMessage 微信客服消息
type CustomerMessage struct {
	ToUser          string               `json:"touser"`
	MsgType         string               `json:"msgtype"`
	CustomService   *ServiceMeta         `json:"customservice,omitempty"`
	Text            *TextMeta            `json:"text,omitempty"`
	Image           *ImageMeta           `json:"image,omitempty"`
	Voice           *VoiceMeta           `json:"voice,omitempty"`
	Video           *VideoMeta           `json:"video,omitempty"`
	Music           *MusicMeta           `json:"music,omitempty"`
	News            *NewsMeta            `json:"news,omitempty"`
	MpNews          *MpNewsMeta          `json:"mpnews,omitempty"`
	MpNewsArticle   *MpNewsArticleMeta   `json:"mpNewsArticle,omitempty"`
	MsgMenu         *MsgMenuMeta         `json:"msgmenu,omitempty"`
	WxCard          *WxCardMeta          `json:"wxcard,omitempty"`
	MiniProgramPage *MiniProgramPageMeta `json:"miniprogrampage,omitempty"`
	Link            *LinkMeta            `json:"link,omitempty"`
}

const customerSendURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"

// SendCustomer 发送微信客服消息
func SendCustomer(accessToken string, msg *CustomerMessage) error {
	url := fmt.Sprintf(customerSendURL, accessToken)
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
