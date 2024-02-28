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
	"github.com/lenye/pmsg/internal/im/weixin/work/message"
)

// undoAppCmd 撤回企业微信应用消息
var undoAppCmd = &cobra.Command{
	Use:   "undo",
	Short: "undo work weixin app message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkUndoAppParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			CorpID:      variable.CorpID,
			CorpSecret:  variable.CorpSecret,
			MsgID:       args[0],
		}
		if err := message.CmdWorkUndoApp(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin app undo -i corp_id -s corp_secret msg_id",
}

func init() {
	workWeiXinSetAccessTokenFlags(undoAppCmd)
}
