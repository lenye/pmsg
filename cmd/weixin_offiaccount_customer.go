package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/customer/message"
)

// weiXinOfficialAccountCustomerCmd 微信公众号客服
var weiXinOfficialAccountCustomerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "publish weixin offiaccount customer message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendCustomerParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			AppID:       appID,
			AppSecret:   appSecret,
			ToUser:      toUser,
			MsgType:     msgType,
			KfAccount:   kfAccount,
			Data:        args[0],
		}
		if err := message.CmdMpSendCustomer(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinOfficialAccountCmd.AddCommand(weiXinOfficialAccountCustomerCmd)

	weiXinSetAccessTokenFlags(weiXinOfficialAccountCustomerCmd)

	weiXinOfficialAccountCustomerCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountCustomerCmd.MarkFlagRequired(flags.ToUser)

	weiXinOfficialAccountCustomerCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	weiXinOfficialAccountCustomerCmd.MarkFlagRequired(flags.MsgType)

	weiXinOfficialAccountCustomerCmd.Flags().StringVarP(&kfAccount, flags.KfAccount, "k", "", "customer account")
}
