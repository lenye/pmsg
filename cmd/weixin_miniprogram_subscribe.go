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
			ToUser:           toUser,
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

	weiXinMiniProgramSubCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "weixin user open id (required)")
	weiXinMiniProgramSubCmd.MarkFlagRequired(flags.ToUser)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&templateID, flags.TemplateID, "p", "", "weixin template id (required)")
	weiXinMiniProgramSubCmd.MarkFlagRequired(flags.TemplateID)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&miniProgramState, flags.MiniProgramState, "g", "", "miniprogram_state")
	weiXinMiniProgramSubCmd.Flags().StringVar(&page, flags.Page, "", "page")
	weiXinMiniProgramSubCmd.Flags().StringVar(&language, flags.Language, "", "language")
}
