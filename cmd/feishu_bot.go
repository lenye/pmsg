// Copyright 2022 The pmsg Authors. All rights reserved.
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

	"github.com/lenye/pmsg/pkg/feishu/bot"
	"github.com/lenye/pmsg/pkg/flags"
)

// feiShuBotCmd 飞书自定义机器人
var feiShuBotCmd = &cobra.Command{
	Use:   "bot",
	Short: "publish fei shu bot message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			Secret:      secret,
			MsgType:     msgType,
			Data:        args[0],
		}
		if err := bot.CmdSend(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	feiShuBotCmd.Flags().StringVarP(&accessToken, flags.AccessToken, "t", "", "feishu bot access token (required)")
	feiShuBotCmd.MarkFlagRequired(flags.AccessToken)

	feiShuBotCmd.Flags().StringVarP(&secret, flags.Secret, "s", "", "sign secret")

	feiShuBotCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	feiShuBotCmd.MarkFlagRequired(flags.MsgType)

}
