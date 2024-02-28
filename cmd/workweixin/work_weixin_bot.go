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
	"github.com/lenye/pmsg/internal/im/weixin/work/bot"
	"github.com/lenye/pmsg/pkg/helper"
)

// botCmd 企业微信群机器人
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "publish work weixin group bot message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent: variable.UserAgent,
			Key:       variable.Secret,
			MsgType:   variable.MsgType,
			AtUser:    variable.AtUser,
			AtMobile:  variable.AtMobile,
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
	Example: "pmsg workweixin bot -k key -m text 'hello world'",
}

func init() {
	botCmd.AddCommand(botUploadCmd)

	workWeiXinBotSetKeyFlags(botCmd)

	botCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "message type (required)")
	_ = botCmd.MarkFlagRequired(flags.MsgType)

	botCmd.Flags().StringVarP(&variable.AtUser, flags.AtUser, "o", "", "work weixin user id list")
	botCmd.Flags().StringVarP(&variable.AtMobile, flags.AtMobile, "b", "", "mobile list")

	botCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")
}

// workWeiXinBotSetKeyFlags 设置企业微信群机器人key命令行参数
func workWeiXinBotSetKeyFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&variable.Secret, flags.Key, "k", "", "work weixin bot key (required)")
	_ = cmd.MarkFlagRequired(flags.Key)
}
