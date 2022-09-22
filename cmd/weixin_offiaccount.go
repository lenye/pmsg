package cmd

import (
	"github.com/spf13/cobra"
)

const (
	nameMsgType   = "type"
	nameKfAccount = "kf_account"
)

var (
	msgType   string
	kfAccount string
)

// weiXinOfficialAccountCmd 微信公众号
var weiXinOfficialAccountCmd = &cobra.Command{
	Use:     "offiaccount",
	Aliases: []string{"mp"},
	Short:   "publish weixin offiaccount message",
	Long: `publish weixin offiaccount message:
 template message,
 template subscribe message (onetime),
 subscribe message,
 customer message`,
}

func init() {
	weiXinCmd.AddCommand(weiXinOfficialAccountCmd)
}
