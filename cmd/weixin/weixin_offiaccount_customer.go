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
	"github.com/lenye/pmsg/im/weixin/customer/message"
)

// officialAccountCustomerCmd 微信公众号客服
var officialAccountCustomerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "发送微信公众号客服消息",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendCustomerParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			AppID:       variable.AppID,
			AppSecret:   variable.AppSecret,
			ToUser:      variable.ToUser,
			MsgType:     variable.MsgType,
			KfAccount:   variable.KfAccount,
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

		if err := message.CmdMpSendCustomer(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg weixin offiaccount customer -i app_id -s app_secret -o open_id -m text 'hello world'",
}

func init() {
	officialAccountCustomerCmd.Flags().SortFlags = false
	weiXinSetAccessTokenFlags(officialAccountCustomerCmd)

	officialAccountCustomerCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "微信接收用户 openid (必填)")
	_ = officialAccountCustomerCmd.MarkFlagRequired(flags.ToUser)

	officialAccountCustomerCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "消息类型 (必填)，text(文本消息)、image(图片消息)、voice(语音消息)、video(视频消息)、music(音乐消息)、news(图文消息)、mpnews(图文消息)、mpnewsarticle(图文消息)、msgmenu(菜单消息)、wxcard(卡券)、miniprogrampage(小程序卡片)")
	_ = officialAccountCustomerCmd.MarkFlagRequired(flags.MsgType)

	officialAccountCustomerCmd.Flags().StringVarP(&variable.KfAccount, flags.KfAccount, "k", "", "客服帐号")

	officialAccountCustomerCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")
}
