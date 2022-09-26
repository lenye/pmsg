package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

// weiXinAccessTokenCmd 获取微信接口调用凭证
var weiXinAccessTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "get weixin access token (offiaccount, miniprogram)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arg := token.CmdTokenParams{
			UserAgent: userAgent,
			AppID:     appID,
			AppSecret: appSecret,
		}
		if err := token.CmdGetAccessToken(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinCmd.AddCommand(weiXinAccessTokenCmd)

	weiXinAccessTokenCmd.Flags().StringVarP(&appID, flags.NameAppID, "i", "", "weixin app id (required)")
	weiXinAccessTokenCmd.MarkFlagRequired(flags.NameAppID)

	weiXinAccessTokenCmd.Flags().StringVarP(&appSecret, flags.NameAppSecret, "s", "", "weixin app secret (required)")
	weiXinAccessTokenCmd.MarkFlagRequired(flags.NameAppSecret)
}
