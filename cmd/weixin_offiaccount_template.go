package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/offiaccount/message"
)

// weiXinOfficialAccountTplCmd 微信公众号模板消息
var weiXinOfficialAccountTplCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "publish weixin offiaccount template message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendTemplateParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			AppID:       appID,
			AppSecret:   appSecret,
			ToUser:      toUser,
			TemplateID:  templateID,
			Url:         url,
			Mini:        mini,
			Color:       color,
			ClientMsgID: clientMsgID,
			Data:        args[0],
		}
		if err := message.CmdMpSendTemplate(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinOfficialAccountCmd.AddCommand(weiXinOfficialAccountTplCmd)

	weiXinSetAccessTokenFlags(weiXinOfficialAccountTplCmd)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountTplCmd.MarkFlagRequired(flags.ToUser)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&templateID, flags.TemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountTplCmd.MarkFlagRequired(flags.TemplateID)

	weiXinOfficialAccountTplCmd.Flags().StringVar(&url, flags.Url, "", "url")
	weiXinOfficialAccountTplCmd.Flags().StringToStringVar(&mini, flags.Mini, nil, "weixin mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")

	weiXinOfficialAccountTplCmd.Flags().StringVar(&color, flags.Color, "", "template color")
	weiXinOfficialAccountTplCmd.Flags().StringVarP(&clientMsgID, flags.ClientMsgID, "c", "", "client message id")
}
