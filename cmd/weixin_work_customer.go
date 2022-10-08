package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// weiXinWorkCustomerCmd 微信客服消息
var weiXinWorkCustomerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "publish work weixin customer message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendCustomerParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			CorpID:      corpID,
			CorpSecret:  corpSecret,
			ToUser:      toUser,
			OpenKfID:    openKfID,
			MsgID:       msgID,
			MsgType:     msgType,
			Data:        args[0],
		}
		if err := message.CmdWorkSendCustomer(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkCmd.AddCommand(weiXinWorkCustomerCmd)

	weiXinWorkSetAccessTokenFlags(weiXinWorkCustomerCmd)

	weiXinWorkCustomerCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "work weixin user id (required)")
	weiXinWorkCustomerCmd.MarkFlagRequired(flags.ToUser)

	weiXinWorkCustomerCmd.Flags().StringVarP(&openKfID, flags.OpenKfID, "k", "", "work weixin customer account id (required)")
	weiXinWorkCustomerCmd.MarkFlagRequired(flags.OpenKfID)

	weiXinWorkCustomerCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	weiXinWorkCustomerCmd.MarkFlagRequired(flags.MsgType)

	weiXinWorkCustomerCmd.Flags().StringVarP(&msgID, flags.MsgID, "c", "", "message id")
}
