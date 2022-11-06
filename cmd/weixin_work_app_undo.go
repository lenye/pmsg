package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// weiXinWorkUndoAppCmd 撤回企业微信应用消息
var weiXinWorkUndoAppCmd = &cobra.Command{
	Use:   "undo",
	Short: "undo work weixin app message",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkUndoAppParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			CorpID:      corpID,
			CorpSecret:  corpSecret,
			MsgID:       msgID,
		}
		if err := message.CmdWorkUndoApp(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkAppCmd.AddCommand(weiXinWorkUndoAppCmd)

	weiXinWorkSetAccessTokenFlags(weiXinWorkUndoAppCmd)

	weiXinWorkUndoAppCmd.Flags().StringVarP(&msgID, flags.MsgID, "c", "", "message id")
}
