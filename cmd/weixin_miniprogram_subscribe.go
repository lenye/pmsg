package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/miniprogram/message"
)

// weiXinMiniProgramSubCmd 微信小程序订阅消息
var weiXinMiniProgramSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin miniprogram subscribe message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMiniSendSubscribeParams{
			UserAgent:        userAgent,
			AccessToken:      accessToken,
			AppID:            appID,
			AppSecret:        appSecret,
			ToUser:           openID,
			TemplateID:       templateID,
			MiniProgramState: miniProgramState,
			Page:             page,
			Language:         language,
			Data:             args[0],
		}
		if err := message.CmdMiniProgramSendSubscribe(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinMiniProgramCmd.AddCommand(weiXinMiniProgramSubCmd)

	weiXinSetAccessTokenFlags(weiXinMiniProgramSubCmd)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&openID, flags.NameOpenID, "o", "", "weixin user open id (required)")
	weiXinMiniProgramSubCmd.MarkFlagRequired(flags.NameOpenID)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&templateID, flags.NameTemplateID, "p", "", "weixin template id (required)")
	weiXinMiniProgramSubCmd.MarkFlagRequired(flags.NameTemplateID)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&miniProgramState, flags.NameMiniProgramState, "g", "", "miniprogram_state")
	weiXinMiniProgramSubCmd.Flags().StringVar(&page, flags.NamePage, "", "page")
	weiXinMiniProgramSubCmd.Flags().StringVar(&language, flags.NameLanguage, "", "language")
}
