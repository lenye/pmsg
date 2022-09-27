package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
)

// weiXinCmd 微信
var weiXinCmd = &cobra.Command{
	Use:     "weixin",
	Aliases: []string{"wx"},
	Short:   "weixin message",
	Long: `get weixin access token,
publish weixin message:
 offiaccount template message,
 offiaccount template subscribe message (onetime),
 offiaccount subscribe message,
 offiaccount customer message,
 miniprogram subscribe message,
 miniprogram customer message`,
}

func init() {
	rootCmd.AddCommand(weiXinCmd)

	weiXinCmd.PersistentFlags().StringVarP(&userAgent, flags.UserAgent, "a", "", "http user agent")
}

// weiXinSetAccessTokenFlags 设置微信access_token或者app_id/app_secret命令行参数
func weiXinSetAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&accessToken, flags.AccessToken, "t", "", "weixin access token")

	cmd.Flags().StringVarP(&appID, flags.AppID, "i", "", "weixin app id (required if app secret is set)")
	cmd.Flags().StringVarP(&appSecret, flags.AppSecret, "s", "", "weixin app secret (required if app id is set)")

	cmd.MarkFlagsMutuallyExclusive(flags.AccessToken, flags.AppID)
	cmd.MarkFlagsRequiredTogether(flags.AppID, flags.AppSecret)
}
