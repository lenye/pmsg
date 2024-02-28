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

package weixin

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin/customer/message"
	"github.com/lenye/pmsg/pkg/helper"
)

// officialAccountCustomerCmd 微信公众号客服
var officialAccountCustomerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "publish weixin official account customer message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendCustomerParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			AppID:       variable.AppID,
			AppSecret:   variable.AppSecret,
			ToUser:      variable.ToUser,
			MsgType:     variable.MsgType,
			KfAccount:   variable.KfAccount,
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

		if err := message.CmdMpSendCustomer(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg weixin offiaccount customer -i app_id -s app_secret -o open_id -m text 'hello world'",
}

func init() {
	weiXinSetAccessTokenFlags(officialAccountCustomerCmd)

	officialAccountCustomerCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "weixin user open id (required)")
	_ = officialAccountCustomerCmd.MarkFlagRequired(flags.ToUser)

	officialAccountCustomerCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	_ = officialAccountCustomerCmd.MarkFlagRequired(flags.MsgType)

	officialAccountCustomerCmd.Flags().StringVarP(&variable.KfAccount, flags.KfAccount, "k", "", "customer account")

	officialAccountCustomerCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")
}
