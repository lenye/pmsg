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

// weiXinCmd represents the weixin command
var weiXinCmd = &cobra.Command{
	Use:   "weixin",
	Short: "weixin message",
	Long: `weixin template message,
weixin kefu message,
weixin miniprogram message,`,
}

func init() {
	rootCmd.AddCommand(weiXinCmd)

	weiXinCmd.PersistentFlags().StringVarP(&userAgent, nameUserAgent, "a", "", "http user agent")

	weiXinCmd.PersistentFlags().StringVarP(&accessToken, nameAccessToken, "t", "", "weixin access token")

	weiXinCmd.PersistentFlags().StringVarP(&appID, nameAppID, "i", "", "weixin app id (required if app secret is set)")
	weiXinCmd.PersistentFlags().StringVarP(&appSecret, nameAppSecret, "s", "", "weixin app secret (required if app id is set)")

	weiXinCmd.MarkFlagsMutuallyExclusive(nameAccessToken, nameAppID)
	weiXinCmd.MarkFlagsRequiredTogether(nameAppID, nameAppSecret)
}
