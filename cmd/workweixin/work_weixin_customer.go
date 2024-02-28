// Copyright 2022-2024 The pmsg Authors. All rights reserved.
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

package workweixin

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin/work/message"
	"github.com/lenye/pmsg/pkg/helper"
)

// customerCmd 微信客服消息
var customerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "publish work weixin customer message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendCustomerParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			CorpID:      variable.CorpID,
			CorpSecret:  variable.CorpSecret,
			ToUser:      variable.ToUser,
			OpenKfID:    variable.OpenKfID,
			MsgID:       variable.MsgID,
			MsgType:     variable.MsgType,
		}

		if variable.IsRaw {
			arg.Data = args[0]
		} else {
			var err error
			arg.Data, err = helper.StrRaw2Interpreted(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if err := message.CmdWorkSendCustomer(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin customer -i corp_id -s corp_secret -o user_id -k kf_id -m text 'hello world'",
}

func init() {
	workWeiXinSetAccessTokenFlags(customerCmd)

	customerCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "work weixin user id (required)")
	_ = customerCmd.MarkFlagRequired(flags.ToUser)

	customerCmd.Flags().StringVarP(&variable.OpenKfID, flags.OpenKfID, "k", "", "work weixin customer account id (required)")
	_ = customerCmd.MarkFlagRequired(flags.OpenKfID)

	customerCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	_ = customerCmd.MarkFlagRequired(flags.MsgType)

	customerCmd.Flags().StringVarP(&variable.MsgID, flags.MsgID, "c", "", "message id")

	customerCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")
}
