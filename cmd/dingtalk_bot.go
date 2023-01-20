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

	"github.com/lenye/pmsg/pkg/dingtalk/bot"
	"github.com/lenye/pmsg/pkg/flags"
)

// dingTalkBotCmd 钉钉自定义机器人
var dingTalkBotCmd = &cobra.Command{
	Use:   "bot",
	Short: "publish ding talk bot message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			Secret:      secret,
			MsgType:     msgType,
			AtUser:      atUser,
			AtMobile:    atMobile,
			IsAtAll:     isAtAll,
			Data:        args[0],
		}
		if err := bot.CmdSend(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg dingtalk bot -t access_token -m text 'hello world'",
}

func init() {
	dingTalkBotCmd.Flags().StringVarP(&accessToken, flags.AccessToken, "t", "", "dingtalk bot access token (required)")
	dingTalkBotCmd.MarkFlagRequired(flags.AccessToken)

	dingTalkBotCmd.Flags().StringVarP(&secret, flags.Secret, "s", "", "sign secret")

	dingTalkBotCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	dingTalkBotCmd.MarkFlagRequired(flags.MsgType)

	dingTalkBotCmd.Flags().StringVarP(&atUser, flags.AtUser, "o", "", "dingtalk user id list")
	dingTalkBotCmd.Flags().StringVarP(&atMobile, flags.AtMobile, "b", "", "mobile list")
	dingTalkBotCmd.Flags().BoolVarP(&isAtAll, flags.IsAtAll, "i", false, "is @all")

}
