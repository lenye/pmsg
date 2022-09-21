package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	nameAccessToken = "access_token"
	nameAppID       = "app_id"
	nameAppSecret   = "app_secret"
)

var (
	accessToken string
	appID       string
	appSecret   string
)

var ErrMultiRequiredOne = errors.New("flags in the group [access_token app_id] required set one")

// weiXinCmd 微信
var weiXinCmd = &cobra.Command{
	Use:     "weixin",
	Aliases: []string{"wx"},
	Short:   "publish weixin message",
	Long: `get weixin access token,
publish weixin message:
 mp template message,
 mp template subscribe message (onetime),
 mp subscribe message,
 miniprogram subscribe message`,
}

func init() {
	rootCmd.AddCommand(weiXinCmd)

	weiXinCmd.PersistentFlags().StringVarP(&userAgent, nameUserAgent, "a", "", "http user agent")
}

// weiXinSetAccessTokenFlags 设置微信access_token或者app_id/app_secret命令行参数
func weiXinSetAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&accessToken, nameAccessToken, "t", "", "weixin access token")

	cmd.Flags().StringVarP(&appID, nameAppID, "i", "", "weixin app id (required if app secret is set)")
	cmd.Flags().StringVarP(&appSecret, nameAppSecret, "s", "", "weixin app secret (required if app id is set)")

	cmd.MarkFlagsMutuallyExclusive(nameAccessToken, nameAppID)
	cmd.MarkFlagsRequiredTogether(nameAppID, nameAppSecret)
}
