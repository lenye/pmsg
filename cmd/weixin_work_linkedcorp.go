package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// weiXinWorkLinkedCorpCmd 企业微信互联企业消息
var weiXinWorkLinkedCorpCmd = &cobra.Command{
	Use:     "linkedcorp",
	Aliases: []string{"lc"},
	Short:   "publish work weixin linkedcorp message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if toUser != "" {
			toUsers = strings.Split(toUser, "|")
		}
		if toParty != "" {
			toPartys = strings.Split(toParty, "|")
		}
		if toTag != "" {
			toTags = strings.Split(toTag, "|")
		}

		arg := message.CmdWorkSendLinkedCorpParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			CorpID:      corpID,
			CorpSecret:  corpSecret,
			ToUser:      toUsers,
			ToParty:     toPartys,
			ToTag:       toTags,
			ToAll:       toAll,
			AgentID:     agentID,
			MsgType:     msgType,
			Safe:        safe,
			Data:        args[0],
		}
		if err := message.CmdWorkSendLinkedCorp(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkCmd.AddCommand(weiXinWorkLinkedCorpCmd)

	weiXinWorkSetAccessTokenFlags(weiXinWorkLinkedCorpCmd)

	weiXinWorkLinkedCorpCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "work weixin user id list")
	weiXinWorkLinkedCorpCmd.Flags().StringVarP(&toParty, flags.ToParty, "p", "", "work weixin party id list")
	weiXinWorkLinkedCorpCmd.Flags().StringVarP(&toTag, flags.ToTag, "g", "", "work weixin tag id list")
	weiXinWorkLinkedCorpCmd.Flags().IntVarP(&toAll, flags.ToAll, "l", 0, "send to all user")

	weiXinWorkLinkedCorpCmd.Flags().Int64VarP(&agentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	weiXinWorkLinkedCorpCmd.MarkFlagRequired(flags.AgentID)

	weiXinWorkLinkedCorpCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	weiXinWorkLinkedCorpCmd.MarkFlagRequired(flags.MsgType)

	weiXinWorkLinkedCorpCmd.Flags().IntVar(&safe, flags.Safe, 0, "safe")

}
