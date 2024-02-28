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
	"github.com/lenye/pmsg/internal/im/weixin/offiaccount/message"
	"github.com/lenye/pmsg/pkg/helper"
)

// officialAccountSubCmd 微信公众号订阅通知消息
var officialAccountSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin official account subscribe message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpBizSendSubscribeParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			AppID:       variable.AppID,
			AppSecret:   variable.AppSecret,
			ToUser:      variable.ToUser,
			TemplateID:  variable.TemplateID,
			Page:        variable.Page,
			Mini:        variable.Mini,
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

		if err := message.CmdMpBizSendSubscribe(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg weixin offiaccount subscribe -i app_id -s app_secret -p template_id -o open_id '{\"first\":{\"value\":\"test\"}}'",
}

func init() {
	weiXinSetAccessTokenFlags(officialAccountSubCmd)

	officialAccountSubCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "weixin user open id (required)")
	_ = officialAccountSubCmd.MarkFlagRequired(flags.ToUser)

	officialAccountSubCmd.Flags().StringVarP(&variable.TemplateID, flags.TemplateID, "p", "", "weixin template id (required)")
	_ = officialAccountSubCmd.MarkFlagRequired(flags.TemplateID)

	officialAccountSubCmd.Flags().StringVar(&variable.Page, flags.Page, "", "Page")
	officialAccountSubCmd.Flags().StringToStringVar(&variable.Mini, flags.Mini, nil, "weixin Mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")

	officialAccountSubCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")
}
