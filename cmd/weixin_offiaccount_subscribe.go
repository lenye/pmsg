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
			ToUser:      openID,
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

	weiXinOfficialAccountSubCmd.Flags().StringVarP(&openID, flags.NameOpenID, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountSubCmd.MarkFlagRequired(flags.NameOpenID)

	weiXinOfficialAccountSubCmd.Flags().StringVarP(&templateID, flags.NameTemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountSubCmd.MarkFlagRequired(flags.NameTemplateID)

	weiXinOfficialAccountSubCmd.Flags().StringVar(&page, flags.NamePage, "", "page")
	weiXinOfficialAccountSubCmd.Flags().StringToStringVarP(&mini, flags.NameMini, "m", nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}
