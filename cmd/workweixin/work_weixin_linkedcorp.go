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

// linkedCorpCmd 企业微信互联企业消息
var linkedCorpCmd = &cobra.Command{
	Use:     "linkedcorp",
	Aliases: []string{"lc"},
	Short:   "publish work weixin linked corp message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendLinkedCorpParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			CorpID:      variable.CorpID,
			CorpSecret:  variable.CorpSecret,
			ToAll:       variable.ToAll,
			AgentID:     variable.AgentID,
			MsgType:     variable.MsgType,
			Safe:        variable.Safe,
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

		if variable.ToUser != "" {
			arg.ToUser = strings.Split(variable.ToUser, "|")
		}
		if variable.ToParty != "" {
			arg.ToParty = strings.Split(variable.ToParty, "|")
		}
		if variable.ToTag != "" {
			arg.ToTag = strings.Split(variable.ToTag, "|")
		}

		if err := message.CmdWorkSendLinkedCorp(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin linkedcorp -i corp_id -s corp_secret -o 'userid1|userid2' -m text 'hello world'",
}

func init() {
	workWeiXinSetAccessTokenFlags(linkedCorpCmd)

	linkedCorpCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "work weixin user id list")
	linkedCorpCmd.Flags().StringVarP(&variable.ToParty, flags.ToParty, "p", "", "work weixin party id list")
	linkedCorpCmd.Flags().StringVarP(&variable.ToTag, flags.ToTag, "g", "", "work weixin tag id list")
	linkedCorpCmd.Flags().IntVarP(&variable.ToAll, flags.ToAll, "l", 0, "send to all user")

	linkedCorpCmd.Flags().Int64VarP(&variable.AgentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	_ = linkedCorpCmd.MarkFlagRequired(flags.AgentID)

	linkedCorpCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	_ = linkedCorpCmd.MarkFlagRequired(flags.MsgType)

	linkedCorpCmd.Flags().IntVar(&variable.Safe, flags.Safe, 0, "Safe")

	linkedCorpCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")

}
