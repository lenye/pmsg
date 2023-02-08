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

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// workWeiXinLinkedCorpCmd 企业微信互联企业消息
var workWeiXinLinkedCorpCmd = &cobra.Command{
	Use:     "linkedcorp",
	Aliases: []string{"lc"},
	Short:   "publish work weixin linked corp message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendLinkedCorpParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			CorpID:      corpID,
			CorpSecret:  corpSecret,
			ToAll:       toAll,
			AgentID:     agentID,
			MsgType:     msgType,
			Safe:        safe,
			Data:        args[0],
		}

		if toUser != "" {
			arg.ToUser = strings.Split(toUser, "|")
		}
		if toParty != "" {
			arg.ToParty = strings.Split(toParty, "|")
		}
		if toTag != "" {
			arg.ToTag = strings.Split(toTag, "|")
		}

		if err := message.CmdWorkSendLinkedCorp(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg workweixin linkedcorp -i corp_id -s corp_secret -o 'userid1|userid2' -m text 'hello world'",
}

func init() {
	workWeiXinSetAccessTokenFlags(workWeiXinLinkedCorpCmd)

	workWeiXinLinkedCorpCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "work weixin user id list")
	workWeiXinLinkedCorpCmd.Flags().StringVarP(&toParty, flags.ToParty, "p", "", "work weixin party id list")
	workWeiXinLinkedCorpCmd.Flags().StringVarP(&toTag, flags.ToTag, "g", "", "work weixin tag id list")
	workWeiXinLinkedCorpCmd.Flags().IntVarP(&toAll, flags.ToAll, "l", 0, "send to all user")

	workWeiXinLinkedCorpCmd.Flags().Int64VarP(&agentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	workWeiXinLinkedCorpCmd.MarkFlagRequired(flags.AgentID)

	workWeiXinLinkedCorpCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	workWeiXinLinkedCorpCmd.MarkFlagRequired(flags.MsgType)

	workWeiXinLinkedCorpCmd.Flags().IntVar(&safe, flags.Safe, 0, "safe")

}
