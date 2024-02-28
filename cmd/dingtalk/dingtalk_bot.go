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

package dingtalk

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/dingtalk/bot"
	"github.com/lenye/pmsg/pkg/helper"
)

// botCmd 钉钉自定义机器人
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "publish ding talk bot message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			Secret:      variable.Secret,
			MsgType:     variable.MsgType,
			AtUser:      variable.AtUser,
			AtMobile:    variable.AtMobile,
			IsAtAll:     variable.IsAtAll,
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

		if err := bot.CmdSend(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg dingtalk bot -t access_token -m text 'hello world'",
}

func init() {
	botCmd.Flags().StringVarP(&variable.AccessToken, flags.AccessToken, "t", "", "dingtalk bot access token (required)")
	_ = botCmd.MarkFlagRequired(flags.AccessToken)

	botCmd.Flags().StringVarP(&variable.Secret, flags.Secret, "s", "", "sign Secret")

	botCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	_ = botCmd.MarkFlagRequired(flags.MsgType)

	botCmd.Flags().StringVarP(&variable.AtUser, flags.AtUser, "o", "", "dingtalk user id list")
	botCmd.Flags().StringVarP(&variable.AtMobile, flags.AtMobile, "b", "", "mobile list")
	botCmd.Flags().BoolVarP(&variable.IsAtAll, flags.IsAtAll, "i", false, "is @all")

	botCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")

}
