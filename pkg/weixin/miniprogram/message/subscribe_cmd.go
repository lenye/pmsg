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

type CmdMiniSendSubscribeParams struct {
	UserAgent        string
	AccessToken      string
	AppID            string
	AppSecret        string
	ToUser           string
	TemplateID       string
	MiniProgramState string
	Page             string
	Language         string
	Data             string
}

func (t *CmdMiniSendSubscribeParams) Validate() error {
	if t.AccessToken == "" && t.AppID == "" {
		return flags.ErrWeixinAccessToken
	}
	if t.Language != "" {
		if err := ValidateLanguage(t.Language); err != nil {
			return fmt.Errorf("invalid flags %s: %v", flags.Language, err)
		}
	}

	if t.MiniProgramState != "" {
		if err := ValidateMiniProgramState(t.MiniProgramState); err != nil {
			return fmt.Errorf("invalid flags %s: %v", flags.MiniProgramState, err)
		}
	}

	return nil
}

// CmdMiniProgramSendSubscribe 发送微信小程序订阅消息
func CmdMiniProgramSendSubscribe(arg *CmdMiniSendSubscribeParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	var dataItem map[string]SubscribeDataItem
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

	msg := SubscribeMessage{
		ToUser:           arg.ToUser,
		TemplateID:       arg.TemplateID,
		Data:             dataItem,
		Page:             arg.Page,
		MiniProgramState: MiniProgramStateFormal,
		Language:         LanguageZhCN,
	}

	if arg.Language != "" {
		msg.Language = arg.Language
	}

	if arg.MiniProgramState != "" {
		msg.MiniProgramState = arg.MiniProgramState
	}

	client.UserAgent = arg.UserAgent

	if arg.AccessToken == "" {
		accessTokenResp, err := token.GetAccessToken(arg.AppID, arg.AppSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if err := SendSubscribe(arg.AccessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
