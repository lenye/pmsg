// Copyright 2022 The pmsg Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	Example: "pmsg workweixin customer -i corp_id -s corp_secret -o user_id -k kf_id -m text 'hello world'",
}

func init() {
	weiXinWorkSetAccessTokenFlags(weiXinWorkCustomerCmd)

	weiXinWorkCustomerCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "work weixin user id (required)")
	weiXinWorkCustomerCmd.MarkFlagRequired(flags.ToUser)

	weiXinWorkCustomerCmd.Flags().StringVarP(&openKfID, flags.OpenKfID, "k", "", "work weixin customer account id (required)")
	weiXinWorkCustomerCmd.MarkFlagRequired(flags.OpenKfID)

	weiXinWorkCustomerCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	weiXinWorkCustomerCmd.MarkFlagRequired(flags.MsgType)

	weiXinWorkCustomerCmd.Flags().StringVarP(&msgID, flags.MsgID, "c", "", "message id")
}
