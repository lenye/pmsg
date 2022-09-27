package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/offiaccount/message"
)

// weiXinOfficialAccountSubCmd 微信公众号订阅通知消息
var weiXinOfficialAccountSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin offiaccount subscribe message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpBizSendSubscribeParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			AppID:       appID,
			AppSecret:   appSecret,
			ToUser:      toUser,
			TemplateID:  templateID,
			Page:        page,
			Mini:        mini,
			Data:        args[0],
		}
		if err := message.CmdMpBizSendSubscribe(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinOfficialAccountCmd.AddCommand(weiXinOfficialAccountSubCmd)

	weiXinSetAccessTokenFlags(weiXinOfficialAccountSubCmd)

	weiXinOfficialAccountSubCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountSubCmd.MarkFlagRequired(flags.ToUser)

	weiXinOfficialAccountSubCmd.Flags().StringVarP(&templateID, flags.TemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountSubCmd.MarkFlagRequired(flags.TemplateID)

	weiXinOfficialAccountSubCmd.Flags().StringVar(&page, flags.Page, "", "page")
	weiXinOfficialAccountSubCmd.Flags().StringToStringVar(&mini, flags.Mini, nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}
