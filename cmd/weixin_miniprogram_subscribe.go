package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/miniprogram/message"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

const (
	namePage             = "page"
	nameMiniProgramState = "miniprogram_state"
	nameLanguage         = "lang"
)

var (
	page             string
	miniProgramState string
	language         string
)

// weiXinMiniProgramSubCmd 微信小程序订阅消息
var weiXinMiniProgramSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin miniprogram subscribe message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinMiniProgramSendSubscribe(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinMiniProgramCmd.AddCommand(weiXinMiniProgramSubCmd)

	weiXinSetAccessTokenFlags(weiXinMiniProgramSubCmd)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinMiniProgramSubCmd.MarkFlagRequired(nameOpenID)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&templateID, nameTemplateID, "p", "", "weixin template id (required)")
	weiXinMiniProgramSubCmd.MarkFlagRequired(nameTemplateID)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&miniProgramState, nameMiniProgramState, "g", "", "miniprogram_state")
	weiXinMiniProgramSubCmd.Flags().StringVar(&page, namePage, "", "page")
	weiXinMiniProgramSubCmd.Flags().StringVar(&language, nameLanguage, "", "language")
}

// WeiXinMiniProgramSendSubscribe 发送微信小程序订阅消息
func WeiXinMiniProgramSendSubscribe(args []string) error {

	if accessToken == "" {
		if appID == "" {
			return ErrMultiRequiredOne
		}
	}

	var dataItem map[string]message.SubscribeDataItem
	buf := bytes.NewBufferString("")
	buf.WriteString(args[0])
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

	msg := message.SubscribeMessage{
		ToUser:           openID,
		TemplateID:       templateID,
		Data:             dataItem,
		MiniProgramState: message.MiniProgramStateFormal,
		Language:         message.LanguageZhCN,
	}

	if page != "" {
		msg.Page = page
	}

	if language != "" {
		if err := message.ValidateLanguage(language); err != nil {
			return fmt.Errorf("invalid flags %s: %v", nameLanguage, err)
		}
		msg.Language = language
	}

	if miniProgramState != "" {
		if err := message.ValidateMiniProgramState(miniProgramState); err != nil {
			return fmt.Errorf("invalid flags %s: %v", nameMiniProgramState, err)
		}
		msg.MiniProgramState = miniProgramState
	}

	if userAgent != "" {
		client.UserAgent = userAgent
	}

	if accessToken == "" {
		accessTokenResp, err := token.GetAccessToken(appID, appSecret)
		if err != nil {
			return err
		}
		accessToken = accessTokenResp.AccessToken
	}

	if err := message.SendSubscribe(accessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
