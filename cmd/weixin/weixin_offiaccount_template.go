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

// officialAccountTplCmd 微信公众号模板消息
var officialAccountTplCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "发送微信公众号模板消息",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendTemplateParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			AppID:       variable.AppID,
			AppSecret:   variable.AppSecret,
			ToUser:      variable.ToUser,
			TemplateID:  variable.TemplateID,
			Url:         variable.Url,
			Mini:        variable.Mini,
			Color:       variable.Color,
			ClientMsgID: variable.ClientMsgID,
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

		if err := message.CmdMpSendTemplate(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg weixin offiaccount template -i app_id -s app_secret -p template_id -o open_id '{\"first\":{\"value\":\"test\"}}'",
}

func init() {
	officialAccountTplCmd.Flags().SortFlags = false
	officialAccountTplCmd.AddCommand(officialAccountTplSubCmd)

	weiXinSetAccessTokenFlags(officialAccountTplCmd)

	officialAccountTplCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "微信接收用户 openid (必填)")
	_ = officialAccountTplCmd.MarkFlagRequired(flags.ToUser)

	officialAccountTplCmd.Flags().StringVarP(&variable.TemplateID, flags.TemplateID, "p", "", "模板 id (必填)")
	_ = officialAccountTplCmd.MarkFlagRequired(flags.TemplateID)

	officialAccountTplCmd.Flags().StringVar(&variable.Url, flags.Url, "", "用户点击后跳转的url")
	officialAccountTplCmd.Flags().StringToStringVar(&variable.Mini, flags.Mini, nil, "跳小程序所需数据, 样例: app_id=XiaoChengXuAppId,page_path=index?foo=bar")

	officialAccountTplCmd.Flags().StringVar(&variable.Color, flags.Color, "", "模板内容字体颜色，不填默认为黑色")
	officialAccountTplCmd.Flags().StringVarP(&variable.ClientMsgID, flags.ClientMsgID, "c", "", "防重入 id")

	officialAccountTplCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "模板数据是原始字符串字面值（不转义处理）")
}
