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

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/bot"
)

// weiXinWorkCmd 企业微信群机器人
var weiXinWorkBotCmd = &cobra.Command{
	Use:   "bot",
	Short: "publish work weixin group bot message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent: userAgent,
			Key:       secret,
			MsgType:   msgType,
			AtUser:    atUser,
			AtMobile:  atMobile,
			Data:      args[0],
		}
		if err := bot.CmdSend(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg workweixin bot -k key -m text 'hello world'",
}

func init() {
	weiXinWorkBotCmd.AddCommand(weiXinWorkBotUploadCmd)

	weiXinWorkSetKeyFlags(weiXinWorkBotCmd)

	weiXinWorkBotCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	weiXinWorkBotCmd.MarkFlagRequired(flags.MsgType)

	weiXinWorkBotCmd.Flags().StringVarP(&atUser, flags.AtUser, "o", "", "work weixin user id list")
	weiXinWorkBotCmd.Flags().StringVarP(&atMobile, flags.AtMobile, "b", "", "mobile list")
}

// weiXinWorkSetKeyFlags 设置企业微信key命令行参数
func weiXinWorkSetKeyFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&secret, flags.Key, "k", "", "work weixin bot key (required)")
	cmd.MarkFlagRequired(flags.Key)
}
