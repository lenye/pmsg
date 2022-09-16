package cmd

import (
	"github.com/spf13/cobra"
)

// weiXinMpCmd 微信公众号
var weiXinMpCmd = &cobra.Command{
	Use:   "mp",
	Short: "publish weixin mp message",
	Long: `publish weixin mp message:
weixin mp template message,
weixin mp template subscribe message`,
}

func init() {
	weiXinCmd.AddCommand(weiXinMpCmd)
}
