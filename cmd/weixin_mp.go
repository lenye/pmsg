package cmd

import (
	"github.com/spf13/cobra"
)

// weiXinMpCmd 微信公众号
var weiXinMpCmd = &cobra.Command{
	Use:   "mp",
	Short: "weixin mp message",
	Long:  `weixin mp template message`,
}

func init() {
	weiXinCmd.AddCommand(weiXinMpCmd)
}
