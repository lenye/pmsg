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
			ToUser:      openID,
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

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&openID, flags.NameOpenID, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountTplCmd.MarkFlagRequired(flags.NameOpenID)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&templateID, flags.NameTemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountTplCmd.MarkFlagRequired(flags.NameTemplateID)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&url, flags.NameUrl, "u", "", "url")
	weiXinOfficialAccountTplCmd.Flags().StringToStringVarP(&mini, flags.NameMini, "m", nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&color, flags.NameColor, "c", "", "weixin template color")
	weiXinOfficialAccountTplCmd.Flags().StringVar(&clientMsgID, flags.NameClientMsgID, "", "weixin template client msg id")
}
