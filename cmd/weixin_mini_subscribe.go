package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin/mini/message"
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

// weiXinMiniSubCmd 微信小程序订阅消息
var weiXinMiniSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin miniprogram subscribe message",
	Long:    `publish weixin miniprogram subscribe message`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinMpSendSubscribe(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinMiniCmd.AddCommand(weiXinMiniSubCmd)

	weiXinMiniSubCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinMiniSubCmd.MarkFlagRequired(nameOpenID)

	weiXinMiniSubCmd.Flags().StringVarP(&templateID, nameTemplateID, "p", "", "weixin template id (required)")
	weiXinMiniSubCmd.MarkFlagRequired(nameTemplateID)

	weiXinMiniSubCmd.Flags().StringVarP(&miniProgramState, nameMiniProgramState, "g", "", "miniprogram_state")
	weiXinMiniSubCmd.Flags().StringVar(&page, namePage, "", "page")
	weiXinMiniSubCmd.Flags().StringVar(&language, nameLanguage, "", "language")
}

// WeiXinMpSendSubscribe 发送微信小程序订阅消息
func WeiXinMpSendSubscribe(args []string) error {

	if userAgent != "" {
		client.UserAgent = userAgent
	}

	if accessToken == "" {
		if appID == "" {
			return ErrMultiRequiredOne
		}
	}

	var dataItem map[string]message.SubscribeDataItem
	data = strings.Join(args, "")
	if data != "" {
		if err := json.Unmarshal([]byte(data), &dataItem); err != nil {
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
		switch language {
		case message.LanguageZhCN, message.LanguageEnUS, message.LanguageZhHK, message.LanguageZhTW:
		default:
			return fmt.Errorf("%s not set to [%q %q %q %q]", nameLanguage, message.LanguageZhCN, message.LanguageEnUS, message.LanguageZhHK, message.LanguageZhTW)
		}
		msg.Language = language
	}

	if miniProgramState != "" {
		switch miniProgramState {
		case message.MiniProgramStateDeveloper, message.MiniProgramStateTrial, message.MiniProgramStateFormal:
		default:
			return fmt.Errorf("%s not set to [%q %q %q]", nameMiniProgramState, message.MiniProgramStateDeveloper, message.MiniProgramStateTrial, message.MiniProgramStateFormal)
		}
		msg.MiniProgramState = miniProgramState
	}

	if accessToken == "" {
		accessTokenResp, err := token.GetAccessToken(appID, appSecret)
		if err != nil {
			return err
		}
		accessToken = accessTokenResp.AccessToken.Token
	}

	if err := message.SendSubscribe(accessToken, &msg); err != nil {
		return err
	}

	return nil
}
