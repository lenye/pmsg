package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// weiXinWorkAppChatCmd 企业微信群聊推送消息
var weiXinWorkAppChatCmd = &cobra.Command{
	Use:     "appchat",
	Aliases: []string{"chat"},
	Short:   "publish work weixin appchat message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendAppChatParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			CorpID:      corpID,
			CorpSecret:  corpSecret,
			ChatID:      chatID,
			MsgType:     msgType,
			Safe:        safe,
			Data:        args[0],
		}
		if err := message.CmdWorkSendAppChat(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkCmd.AddCommand(weiXinWorkAppChatCmd)

	weiXinWorkSetAccessTokenFlags(weiXinWorkAppChatCmd)

	weiXinWorkAppChatCmd.Flags().StringVarP(&chatID, flags.ChatID, "c", "", "work weixin chat id (required)")
	weiXinWorkAppChatCmd.MarkFlagRequired(flags.ChatID)

	weiXinWorkAppChatCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	weiXinWorkAppChatCmd.MarkFlagRequired(flags.MsgType)

	weiXinWorkAppChatCmd.Flags().IntVar(&safe, flags.Safe, 0, "safe")
}
