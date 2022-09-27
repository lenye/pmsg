package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/offiaccount/message"
)

// weiXinOfficialAccountTplSubCmd 微信公众号一次性订阅消息
var weiXinOfficialAccountTplSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin offiaccount template subscribe message (onetime)",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendTemplateSubscribeParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			AppID:       appID,
			AppSecret:   appSecret,
			ToUser:      toUser,
			TemplateID:  templateID,
			Scene:       scene,
			Title:       title,
			Url:         url,
			Mini:        mini,
			Data:        args[0],
		}
		if err := message.CmdMpSendTemplateSubscribe(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinOfficialAccountTplCmd.AddCommand(weiXinOfficialAccountTplSubCmd)

	weiXinSetAccessTokenFlags(weiXinOfficialAccountTplSubCmd)

	weiXinOfficialAccountTplSubCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(flags.ToUser)

	weiXinOfficialAccountTplSubCmd.Flags().StringVarP(&templateID, flags.TemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(flags.TemplateID)

	weiXinOfficialAccountTplSubCmd.Flags().StringVar(&scene, flags.Scene, "", "weixin subscribe scene (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(flags.Scene)

	weiXinOfficialAccountTplSubCmd.Flags().StringVar(&title, flags.Title, "", "weixin message title (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(flags.Title)

	weiXinOfficialAccountTplSubCmd.Flags().StringVar(&url, flags.Url, "", "url")
	weiXinOfficialAccountTplSubCmd.Flags().StringToStringVar(&mini, flags.Mini, nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}
