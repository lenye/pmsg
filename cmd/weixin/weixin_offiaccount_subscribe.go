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
	"github.com/lenye/pmsg/im/weixin/offiaccount/message"
)

// officialAccountSubCmd 微信公众号订阅通知消息
var officialAccountSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "发送微信公众号订阅通知消息",
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
			arg.Data, err = im.StrRaw2Interpreted(args[0])
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
	officialAccountSubCmd.Flags().SortFlags = false
	weiXinSetAccessTokenFlags(officialAccountSubCmd)

	officialAccountSubCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "微信接收用户 openid (必填)")
	_ = officialAccountSubCmd.MarkFlagRequired(flags.ToUser)

	officialAccountSubCmd.Flags().StringVarP(&variable.TemplateID, flags.TemplateID, "p", "", "模板 id (必填)")
	_ = officialAccountSubCmd.MarkFlagRequired(flags.TemplateID)

	officialAccountSubCmd.Flags().StringVar(&variable.Page, flags.Page, "", "跳转网页时填写")
	officialAccountSubCmd.Flags().StringToStringVar(&variable.Mini, flags.Mini, nil, "跳小程序所需数据, 样例: app_id=XiaoChengXuAppId,page_path=index?foo=bar")

	officialAccountSubCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "模板数据是原始字符串字面值（不转义处理）")
}
