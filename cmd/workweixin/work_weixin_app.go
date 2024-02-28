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

// appCmd 企业微信应用消息
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "publish work weixin app message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendAppParams{
			UserAgent:              variable.UserAgent,
			AccessToken:            variable.AccessToken,
			CorpID:                 variable.CorpID,
			CorpSecret:             variable.CorpSecret,
			ToUser:                 variable.ToUser,
			ToParty:                variable.ToParty,
			ToTag:                  variable.ToTag,
			AgentID:                variable.AgentID,
			MsgType:                variable.MsgType,
			Safe:                   variable.Safe,
			EnableIDTrans:          variable.EnableIDTrans,
			EnableDuplicateCheck:   variable.EnableDuplicateCheck,
			DuplicateCheckInterval: variable.DuplicateCheckInterval,
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

		if err := message.CmdWorkSendApp(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin app -i corp_id -s corp_secret -e agent_id -o '@all' -m text 'hello world'",
}

func init() {
	appCmd.AddCommand(undoAppCmd)

	workWeiXinSetAccessTokenFlags(appCmd)

	appCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "work weixin user id list")
	appCmd.Flags().StringVarP(&variable.ToParty, flags.ToParty, "p", "", "work weixin party id list")
	appCmd.Flags().StringVarP(&variable.ToTag, flags.ToTag, "g", "", "work weixin tag id list")

	appCmd.Flags().Int64VarP(&variable.AgentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	_ = appCmd.MarkFlagRequired(flags.AgentID)

	appCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	_ = appCmd.MarkFlagRequired(flags.MsgType)

	appCmd.Flags().IntVar(&variable.Safe, flags.Safe, 0, "Safe")
	appCmd.Flags().IntVarP(&variable.EnableIDTrans, flags.EnableIDTrans, "r", 0, "enable id translated")
	appCmd.Flags().IntVarP(&variable.EnableDuplicateCheck, flags.EnableDuplicateCheck, "c", 0, "enable duplicate check")
	appCmd.Flags().IntVarP(&variable.DuplicateCheckInterval, flags.DuplicateCheckInterval, "d", 1800, "duplicate check interval")

	appCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")
}
