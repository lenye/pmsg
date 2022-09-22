package cmd

import (
	"github.com/spf13/cobra"
)

// weiXinMiniProgramCmd 微信小程序
var weiXinMiniProgramCmd = &cobra.Command{
	Use:     "miniprogram",
	Aliases: []string{"mini"},
	Short:   "publish weixin miniprogram message",
	Long: `publish weixin miniprogram message:
 subscribe message,
 customer message`,
}

func init() {
	weiXinCmd.AddCommand(weiXinMiniProgramCmd)
}
