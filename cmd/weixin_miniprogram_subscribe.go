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
	"github.com/lenye/pmsg/pkg/weixin/miniprogram/message"
)

// weiXinMiniProgramSubCmd 微信小程序订阅消息
var weiXinMiniProgramSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin miniprogram subscribe message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMiniSendSubscribeParams{
			UserAgent:        userAgent,
			AccessToken:      accessToken,
			AppID:            appID,
			AppSecret:        appSecret,
			ToUser:           toUser,
			TemplateID:       templateID,
			MiniProgramState: miniProgramState,
			Page:             page,
			Language:         language,
			Data:             args[0],
		}
		if err := message.CmdMiniProgramSendSubscribe(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg weixin miniprogram subscribe -i app_id -s app_secret -p template_id -o open_id '{\"first\":{\"value\":\"test\"}}'",
}

func init() {
	weiXinSetAccessTokenFlags(weiXinMiniProgramSubCmd)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "weixin user open id (required)")
	weiXinMiniProgramSubCmd.MarkFlagRequired(flags.ToUser)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&templateID, flags.TemplateID, "p", "", "weixin template id (required)")
	weiXinMiniProgramSubCmd.MarkFlagRequired(flags.TemplateID)

	weiXinMiniProgramSubCmd.Flags().StringVarP(&miniProgramState, flags.MiniProgramState, "g", "", "miniprogram_state")
	weiXinMiniProgramSubCmd.Flags().StringVar(&page, flags.Page, "", "page")
	weiXinMiniProgramSubCmd.Flags().StringVar(&language, flags.Language, "", "language")
}
