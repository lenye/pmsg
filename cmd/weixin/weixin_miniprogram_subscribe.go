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

package weixin

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin/miniprogram/message"
	"github.com/lenye/pmsg/pkg/helper"
)

// miniProgramSubCmd 微信小程序订阅消息
var miniProgramSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin miniprogram subscribe message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMiniSendSubscribeParams{
			UserAgent:        variable.UserAgent,
			AccessToken:      variable.AccessToken,
			AppID:            variable.AppID,
			AppSecret:        variable.AppSecret,
			ToUser:           variable.ToUser,
			TemplateID:       variable.TemplateID,
			MiniProgramState: variable.MiniProgramState,
			Page:             variable.Page,
			Language:         variable.Language,
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

		if err := message.CmdMiniProgramSendSubscribe(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg weixin miniprogram subscribe -i app_id -s app_secret -p template_id -o open_id '{\"first\":{\"value\":\"test\"}}'",
}

func init() {
	weiXinSetAccessTokenFlags(miniProgramSubCmd)

	miniProgramSubCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "weixin user open id (required)")
	_ = miniProgramSubCmd.MarkFlagRequired(flags.ToUser)

	miniProgramSubCmd.Flags().StringVarP(&variable.TemplateID, flags.TemplateID, "p", "", "weixin template id (required)")
	_ = miniProgramSubCmd.MarkFlagRequired(flags.TemplateID)

	miniProgramSubCmd.Flags().StringVarP(&variable.MiniProgramState, flags.MiniProgramState, "g", "", "miniprogram_state")
	miniProgramSubCmd.Flags().StringVar(&variable.Page, flags.Page, "", "Page")
	miniProgramSubCmd.Flags().StringVar(&variable.Language, flags.Language, "", "Language")

	miniProgramSubCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")
}
