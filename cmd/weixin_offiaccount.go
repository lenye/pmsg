package cmd

import (
	"github.com/spf13/cobra"
)

// weiXinOfficialAccountCmd 微信公众号
var weiXinOfficialAccountCmd = &cobra.Command{
	Use:     "offiaccount",
	Aliases: []string{"mp"},
	Short:   "publish weixin offiaccount message",
}

func init() {
	weiXinCmd.AddCommand(weiXinOfficialAccountCmd)
}
