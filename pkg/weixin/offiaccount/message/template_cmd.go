package message

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

type CmdMpSendTemplateParams struct {
	UserAgent   string
	AccessToken string
	AppID       string
	AppSecret   string
	ToUser      string
	TemplateID  string
	Url         string
	Mini        map[string]string
	Color       string
	ClientMsgID string
	Data        string
}

func (t *CmdMpSendTemplateParams) Validate() error {
	if t.AccessToken == "" && t.AppID == "" {
		return flags.ErrMultiRequiredOne
	}

	// 跳小程序
	if t.Mini != nil {
		if miniAppID, ok := t.Mini[flags.NameMiniAppID]; !ok {
			return fmt.Errorf("mini flag %q not set", flags.NameMiniAppID)
		} else {
			if miniAppID == "" {
				return fmt.Errorf("mini flag %q not set", flags.NameMiniAppID)
			}
		}

		if miniPagePath, ok := t.Mini[flags.NameMiniPagePath]; !ok {
			return fmt.Errorf("mini flag %q not set", flags.NameMiniPagePath)
		} else {
			if miniPagePath == "" {
				return fmt.Errorf("mini flag %q not set", flags.NameMiniPagePath)
			}
		}
	}

	return nil
}

// CmdMpSendTemplate 发送微信公众号模板消息
func CmdMpSendTemplate(arg *CmdMpSendTemplateParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	var dataItem map[string]TemplateDataItem
	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	if buf.String() != "" {
		if err := json.Unmarshal(buf.Bytes(), &dataItem); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		for k, v := range dataItem {
			if v.Value == "" {
				return fmt.Errorf("data %v.value not set", k)
			}
		}
	}

	msg := TemplateMessage{
		ToUser:      arg.ToUser,
		TemplateID:  arg.TemplateID,
		Data:        dataItem,
		URL:         arg.Url,
		ClientMsgID: arg.ClientMsgID,
		Color:       arg.Color,
	}

	// 跳小程序
	if arg.Mini != nil {
		miniAppID, _ := arg.Mini[flags.NameMiniAppID]
		miniPagePath, _ := arg.Mini[flags.NameMiniPagePath]
		msg.MiniProgram = &MiniProgramMeta{
			AppID:    miniAppID,
			PagePath: miniPagePath,
		}
	}

	client.UserAgent = arg.UserAgent

	if arg.AccessToken == "" {
		accessTokenResp, err := token.GetAccessToken(arg.AppID, arg.AppSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if gotMsgID, err := SendTemplate(arg.AccessToken, &msg); err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("%v; msgid: %v", weixin.MessageOK, gotMsgID))
	}

	return nil
}
