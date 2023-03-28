// Copyright 2022-2023 The pmsg Authors. All rights reserved.
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
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/customer/message"
)

// miniProgramCustomerCmd 发送微信小程序客服消息
var miniProgramCustomerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "publish weixin miniprogram customer message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMiniSendCustomerParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			AppID:       variable.AppID,
			AppSecret:   variable.AppSecret,
			ToUser:      variable.ToUser,
			MsgType:     variable.MsgType,
			Data:        args[0],
		}
		if err := message.CmdMiniSendCustomer(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg weixin miniprogram customer -i app_id -s app_secret -o open_id -m text 'hello world'",
}

func init() {
	weiXinSetAccessTokenFlags(miniProgramCustomerCmd)

	miniProgramCustomerCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "weixin user open id (required)")
	miniProgramCustomerCmd.MarkFlagRequired(flags.ToUser)

	miniProgramCustomerCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	miniProgramCustomerCmd.MarkFlagRequired(flags.MsgType)
}
