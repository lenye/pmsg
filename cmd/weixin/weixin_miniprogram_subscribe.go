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
	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/im"
	"github.com/lenye/pmsg/im/weixin/miniprogram/message"
)

// miniProgramSubCmd 微信小程序订阅消息
var miniProgramSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "发送微信小程序订阅消息",
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
			arg.Data, err = im.StrRaw2Interpreted(args[0])
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
	miniProgramSubCmd.Flags().SortFlags = false
	weiXinSetAccessTokenFlags(miniProgramSubCmd)

	miniProgramSubCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "微信接收用户 openid (必填)")
	_ = miniProgramSubCmd.MarkFlagRequired(flags.ToUser)

	miniProgramSubCmd.Flags().StringVarP(&variable.TemplateID, flags.TemplateID, "p", "", "模板 id (必填)")
	_ = miniProgramSubCmd.MarkFlagRequired(flags.TemplateID)

	miniProgramSubCmd.Flags().StringVarP(&variable.MiniProgramState, flags.MiniProgramState, "g", "", "跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版")
	miniProgramSubCmd.Flags().StringVar(&variable.Page, flags.Page, "", "点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转")
	miniProgramSubCmd.Flags().StringVar(&variable.Language, flags.Language, "", "进入小程序查看的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN")

	miniProgramSubCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")
}
