package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/customer/message"
)

// weiXinMiniProgramCustomerCmd 发送微信小程序客服消息
var weiXinMiniProgramCustomerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "publish weixin miniprogram customer message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMiniSendCustomerParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			AppID:       appID,
			AppSecret:   appSecret,
			ToUser:      openID,
			MsgType:     msgType,
			Data:        args[0],
		}
		if err := message.CmdMiniSendCustomer(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinMiniProgramCmd.AddCommand(weiXinMiniProgramCustomerCmd)

	weiXinSetAccessTokenFlags(weiXinMiniProgramCustomerCmd)

	weiXinMiniProgramCustomerCmd.Flags().StringVarP(&openID, flags.NameOpenID, "o", "", "weixin user open id (required)")
	weiXinMiniProgramCustomerCmd.MarkFlagRequired(flags.NameOpenID)

	weiXinMiniProgramCustomerCmd.Flags().StringVar(&msgType, flags.NameMsgType, "", "message type (required)")
	weiXinMiniProgramCustomerCmd.MarkFlagRequired(flags.NameMsgType)
}
