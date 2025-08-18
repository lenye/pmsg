// Copyright 2022-2025 The pmsg Authors. All rights reserved.
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

package mattermost

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/im"
	"github.com/lenye/pmsg/im/mattermost/bot"
)

// botCmd mattermost bot
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "publish mattermost bot message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent: variable.UserAgent,
			URL:       variable.Url,
		}

		if variable.IsRaw {
			arg.Data = args[0]
		} else {
			var err error
			arg.Data, err = im.StrRaw2Interpreted(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if err := bot.CmdSend(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg mattermost bot --Url webhook_url 'Hello, World!'",
}

func init() {
	botCmd.Flags().SortFlags = false
	botCmd.Flags().StringVar(&variable.Url, flags.Url, "", "mattermost webhook Url")
	_ = botCmd.MarkFlagRequired(flags.Url)

	botCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")
}
