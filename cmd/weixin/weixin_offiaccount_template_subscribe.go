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

// officialAccountTplSubCmd 微信公众号一次性订阅消息
var officialAccountTplSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "发送微信公众号一次性订阅消息",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendTemplateSubscribeParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			AppID:       variable.AppID,
			AppSecret:   variable.AppSecret,
			ToUser:      variable.ToUser,
			TemplateID:  variable.TemplateID,
			Scene:       variable.Scene,
			Title:       variable.Title,
			Url:         variable.Url,
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

		if err := message.CmdMpSendTemplateSubscribe(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg weixin offiaccount template subscribe -i app_id -s app_secret --Scene Scene --Title Title -p template_id -o open_id '{\"first\":{\"value\":\"test\"}}'",
}

func init() {
	officialAccountTplSubCmd.Flags().SortFlags = false
	weiXinSetAccessTokenFlags(officialAccountTplSubCmd)

	officialAccountTplSubCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "微信接收用户 openid (必填)")
	_ = officialAccountTplSubCmd.MarkFlagRequired(flags.ToUser)

	officialAccountTplSubCmd.Flags().StringVarP(&variable.TemplateID, flags.TemplateID, "p", "", "模板 id (必填)")
	_ = officialAccountTplSubCmd.MarkFlagRequired(flags.TemplateID)

	officialAccountTplSubCmd.Flags().StringVar(&variable.Scene, flags.Scene, "", "重定向后会带上scene参数，开发者可以填0-10000的整型值，用来标识订阅场景值 (必填)")
	_ = officialAccountTplSubCmd.MarkFlagRequired(flags.Scene)

	officialAccountTplSubCmd.Flags().StringVar(&variable.Title, flags.Title, "", "消息标题 (必填)")
	_ = officialAccountTplSubCmd.MarkFlagRequired(flags.Title)

	officialAccountTplSubCmd.Flags().StringVar(&variable.Url, flags.Url, "", "用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中")
	officialAccountTplSubCmd.Flags().StringToStringVar(&variable.Mini, flags.Mini, nil, "跳小程序所需数据, 样例: app_id=XiaoChengXuAppId,page_path=index?foo=bar")

	officialAccountTplSubCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "模板数据是原始字符串字面值（不转义处理）")
}
