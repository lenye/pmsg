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
			ToUser:      openID,
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

	weiXinOfficialAccountCustomerCmd.Flags().StringVarP(&openID, flags.NameOpenID, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountCustomerCmd.MarkFlagRequired(flags.NameOpenID)

	weiXinOfficialAccountCustomerCmd.Flags().StringVar(&msgType, flags.NameMsgType, "", "message type (required)")
	weiXinOfficialAccountCustomerCmd.MarkFlagRequired(flags.NameMsgType)

	weiXinOfficialAccountCustomerCmd.Flags().StringVarP(&kfAccount, flags.NameKfAccount, "k", "", "customer account")
}
