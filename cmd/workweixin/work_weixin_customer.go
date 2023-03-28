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

package workweixin

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
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
			Data:        args[0],
		}
		if err := message.CmdWorkSendCustomer(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg workweixin customer -i corp_id -s corp_secret -o user_id -k kf_id -m text 'hello world'",
}

func init() {
	workWeiXinSetAccessTokenFlags(customerCmd)

	customerCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "work weixin user id (required)")
	customerCmd.MarkFlagRequired(flags.ToUser)

	customerCmd.Flags().StringVarP(&variable.OpenKfID, flags.OpenKfID, "k", "", "work weixin customer account id (required)")
	customerCmd.MarkFlagRequired(flags.OpenKfID)

	customerCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	customerCmd.MarkFlagRequired(flags.MsgType)

	customerCmd.Flags().StringVarP(&variable.MsgID, flags.MsgID, "c", "", "message id")
}
