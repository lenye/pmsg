package cmd

import (
	"github.com/spf13/cobra"
)

// weiXinWorkCmd 企业微信
var weiXinWorkCmd = &cobra.Command{
	Use:   "work",
	Short: "publish weixin work message",
}

func init() {
	weiXinCmd.AddCommand(weiXinWorkCmd)
}
