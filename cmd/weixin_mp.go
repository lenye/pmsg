package cmd

import (
	"github.com/spf13/cobra"
)

// weiXinMpCmd 微信公众号
var weiXinMpCmd = &cobra.Command{
	Use:   "mp",
	Short: "publish weixin mp message",
	Long: `publish weixin mp message:
 template message,
 template subscribe message (onetime),
 subscribe message`,
}

func init() {
	weiXinCmd.AddCommand(weiXinMpCmd)
}
