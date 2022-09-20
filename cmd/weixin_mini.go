package cmd

import (
	"github.com/spf13/cobra"
)

// weiXinMiniCmd 微信小程序
var weiXinMiniCmd = &cobra.Command{
	Use:   "mini",
	Short: "publish weixin miniprogram message",
	Long: `publish weixin miniprogram message:
 subscribe message`,
}

func init() {
	weiXinCmd.AddCommand(weiXinMiniCmd)
}
