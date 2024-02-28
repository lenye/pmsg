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
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin/work/message"
	"github.com/lenye/pmsg/pkg/helper"
)

// externalContactCmd 企业微信家校消息
var externalContactCmd = &cobra.Command{
	Use:     "externalcontact",
	Aliases: []string{"ec"},
	Short:   "publish work weixin external contact message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendExternalContactParams{
			UserAgent:              variable.UserAgent,
			AccessToken:            variable.AccessToken,
			CorpID:                 variable.CorpID,
			CorpSecret:             variable.CorpSecret,
			RecvScope:              variable.RecvScope,
			ToAll:                  variable.ToAll,
			MsgType:                variable.MsgType,
			AgentID:                variable.AgentID,
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

		if variable.ToParentUserID != "" {
			arg.ToParentUserID = strings.Split(variable.ToParentUserID, "|")
		}
		if variable.ToStudentUserID != "" {
			arg.ToStudentUserID = strings.Split(variable.ToStudentUserID, "|")
		}
		if variable.ToParty != "" {
			arg.ToParty = strings.Split(variable.ToParty, "|")
		}

		if err := message.CmdWorkSendExternalContact(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin externalcontact -i corp_id -s corp_secret -e agent_id -n 'parentuserid1|parentuserid2' -m text 'hello world'",
}

func init() {
	workWeiXinSetAccessTokenFlags(externalContactCmd)

	externalContactCmd.Flags().IntVarP(&variable.RecvScope, flags.RecvScope, "o", 0, "receive scope")

	externalContactCmd.Flags().StringVarP(&variable.ToParentUserID, flags.ToParentUserID, "n", "", "work weixin parent user id list")
	externalContactCmd.Flags().StringVarP(&variable.ToStudentUserID, flags.ToStudentUserID, "u", "", "work weixin student user id list")
	externalContactCmd.Flags().StringVarP(&variable.ToParty, flags.ToParty, "p", "", "work weixin party id list")
	externalContactCmd.Flags().IntVarP(&variable.ToAll, flags.ToAll, "l", 0, "send to all user")

	externalContactCmd.Flags().Int64VarP(&variable.AgentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	_ = externalContactCmd.MarkFlagRequired(flags.AgentID)

	externalContactCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	_ = externalContactCmd.MarkFlagRequired(flags.MsgType)

	externalContactCmd.Flags().IntVarP(&variable.EnableIDTrans, flags.EnableIDTrans, "r", 0, "enable id translated")
	externalContactCmd.Flags().IntVarP(&variable.EnableDuplicateCheck, flags.EnableDuplicateCheck, "c", 0, "enable duplicate check")
	externalContactCmd.Flags().IntVarP(&variable.DuplicateCheckInterval, flags.DuplicateCheckInterval, "d", 1800, "duplicate check interval")

	externalContactCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")

}
