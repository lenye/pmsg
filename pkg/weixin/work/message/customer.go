package message

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

// 企业微信客服消息 类型
const (
	CustomerMsgTypeText        = "text"        // 文本消息
	CustomerMsgTypeImage       = "image"       // 图片消息
	CustomerMsgTypeVoice       = "voice"       // 语音消息
	CustomerMsgTypeVideo       = "video"       // 视频消息
	CustomerMsgTypeFile        = "file"        // 文件消息
	CustomerMsgTypeLink        = "link"        // 图文链接消息
	CustomerMsgTypeMiniProgram = "miniprogram" // 小程序消息
	CustomerMsgTypeMsgMenu     = "msgmenu"     // 菜单消息
	CustomerMsgTypeLocation    = "location"    // 地理位置消息
)

// ValidateCustomerMsgType 验证
func ValidateCustomerMsgType(v string) error {
	switch v {
	case CustomerMsgTypeText, CustomerMsgTypeImage, CustomerMsgTypeVoice, CustomerMsgTypeVideo,
		CustomerMsgTypeFile, CustomerMsgTypeLink, CustomerMsgTypeMiniProgram, CustomerMsgTypeMsgMenu,
		CustomerMsgTypeLocation:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q %q %q]", v,
			CustomerMsgTypeText, CustomerMsgTypeImage, CustomerMsgTypeVoice, CustomerMsgTypeVideo,
			CustomerMsgTypeFile, CustomerMsgTypeLink, CustomerMsgTypeMiniProgram, CustomerMsgTypeMsgMenu,
			CustomerMsgTypeLocation)
	}
	return nil
}

// CustomerMessage 企业微信客服消息
type CustomerMessage struct {
	ToUser      string           `json:"touser"`
	OpenKfID    string           `json:"open_kfid"`
	MsgID       string           `json:"msgid,omitempty"`
	MsgType     string           `json:"msgtype"`
	Text        *TextMeta        `json:"text,omitempty"`
	Image       *ImageMeta       `json:"image,omitempty"`
	Voice       *VoiceMeta       `json:"voice,omitempty"`
	Video       *VideoMeta       `json:"video,omitempty"`
	File        *FileMeta        `json:"file,omitempty"`
	Link        *LinkMeta        `json:"link,omitempty"`
	MiniProgram *MiniProgramMeta `json:"miniprogram,omitempty"`
	MsgMenu     *MsgMenuMeta     `json:"msgmenu,omitempty"`
	Location    *LocationMeta    `json:"location,omitempty"`
}

// CustomerMessageResponse 企业微信客服消息的响应
type CustomerMessageResponse struct {
	weixin.ResponseMeta
	// 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。
	// 若指定msgid，开发者需确保客服账号内唯一，否则接口返回错误。
	// 不多于32字节	字符串取值范围(正则表达式)：[0-9a-zA-Z_-]
	MsgID string `json:"msgid,omitempty"`
}

func (t CustomerMessageResponse) String() string {
	if t.Succeed() {
		return fmt.Sprintf("msgid: %q", t.MsgID)
	}
	return fmt.Sprintf("errcode: %v, errmsg: %q, msgid: %q", t.ErrorCode, t.ErrorMessage, t.MsgID)
}

const customerSendURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=%s"

// SendCustomer 发送微信客服消息
func SendCustomer(accessToken string, msg *CustomerMessage) (*CustomerMessageResponse, error) {
	url := fmt.Sprintf(customerSendURL, accessToken)
	var resp CustomerMessageResponse
	_, err := client.PostJSON(url, msg, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; %v", weixin.ErrWeiXinRequest, resp)
	}
	return &resp, nil
}
