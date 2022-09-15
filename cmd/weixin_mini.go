package cmd

import (
	"github.com/spf13/cobra"
)

// weiXinMiniCmd 微信小程序
var weiXinMiniCmd = &cobra.Command{
	Use:   "mini",
	Short: "weixin miniprogram message",
	Long:  `weixin miniprogram subscribe message`,
}

func init() {
	weiXinCmd.AddCommand(weiXinMiniCmd)
}
