package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// weiXinWorkExternalContactCmd 企业微信家校消息
var weiXinWorkExternalContactCmd = &cobra.Command{
	Use:     "externalcontact",
	Aliases: []string{"ec"},
	Short:   "publish work weixin externalcontact message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if toParentUserID != "" {
			toParentUserIDs = strings.Split(toParentUserID, "|")
		}
		if toStudentUserID != "" {
			toStudentUserIDs = strings.Split(toStudentUserID, "|")
		}
		if toParty != "" {
			toPartys = strings.Split(toParty, "|")
		}

		arg := message.CmdWorkSendExternalContactParams{
			UserAgent:              userAgent,
			AccessToken:            accessToken,
			CorpID:                 corpID,
			CorpSecret:             corpSecret,
			RecvScope:              recvScope,
			ToParentUserID:         toParentUserIDs,
			ToStudentUserID:        toStudentUserIDs,
			ToParty:                toPartys,
			ToAll:                  toAll,
			MsgType:                msgType,
			AgentID:                agentID,
			EnableIDTrans:          enableIDTrans,
			EnableDuplicateCheck:   enableDuplicateCheck,
			DuplicateCheckInterval: duplicateCheckInterval,
			Data:                   args[0],
		}
		if err := message.CmdWorkSendExternalContact(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkCmd.AddCommand(weiXinWorkExternalContactCmd)

	weiXinWorkSetAccessTokenFlags(weiXinWorkExternalContactCmd)

	weiXinWorkExternalContactCmd.Flags().IntVarP(&recvScope, flags.RecvScope, "o", 0, "receive scope")

	weiXinWorkExternalContactCmd.Flags().StringVarP(&toParentUserID, flags.ToParentUserID, "n", "", "work weixin parent user id list")
	weiXinWorkExternalContactCmd.Flags().StringVarP(&toStudentUserID, flags.ToStudentUserID, "u", "", "work weixin student user id list")
	weiXinWorkExternalContactCmd.Flags().StringVarP(&toParty, flags.ToParty, "p", "", "work weixin party id list")
	weiXinWorkExternalContactCmd.Flags().IntVarP(&toAll, flags.ToAll, "l", 0, "send to all user")

	weiXinWorkExternalContactCmd.Flags().Int64VarP(&agentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	weiXinWorkExternalContactCmd.MarkFlagRequired(flags.AgentID)

	weiXinWorkExternalContactCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	weiXinWorkExternalContactCmd.MarkFlagRequired(flags.MsgType)

	weiXinWorkExternalContactCmd.Flags().IntVarP(&enableIDTrans, flags.EnableIDTrans, "r", 0, "enable id translated")
	weiXinWorkExternalContactCmd.Flags().IntVarP(&enableDuplicateCheck, flags.EnableDuplicateCheck, "c", 0, "enable duplicate check")
	weiXinWorkExternalContactCmd.Flags().IntVarP(&duplicateCheckInterval, flags.DuplicateCheckInterval, "d", 1800, "duplicate check interval")

}
