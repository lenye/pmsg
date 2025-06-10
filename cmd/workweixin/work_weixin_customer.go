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

package workweixin

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/im"
	"github.com/lenye/pmsg/im/weixin/work/message"
)

// customerCmd 企业微信客服消息
var customerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "发送企业微信客服消息",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendCustomerParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			CorpID:      variable.CorpID,
			CorpSecret:  variable.CorpSecret,
			ToUser:      variable.ToUser,
			OpenKfID:    variable.OpenKfID,
			MsgID:       variable.MsgID,
			MsgType:     variable.MsgType,
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

		if err := message.CmdWorkSendCustomer(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin customer -i corp_id -s corp_secret -o user_id -k kf_id -m text 'hello world'",
}

func init() {
	customerCmd.Flags().SortFlags = false
	workWeiXinSetAccessTokenFlags(customerCmd)

	customerCmd.Flags().StringVarP(&variable.OpenKfID, flags.OpenKfID, "k", "", "指定发送消息的客服帐号 id (必填)")
	_ = customerCmd.MarkFlagRequired(flags.OpenKfID)

	customerCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "指定接收消息的客户 userid (必填)")
	_ = customerCmd.MarkFlagRequired(flags.ToUser)

	customerCmd.Flags().StringVarP(&variable.MsgID, flags.MsgID, "c", "", "指定消息 id")

	customerCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "消息类型 (必填) text(文本消息)、image(图片消息)、voice(语音消息)、video(视频消息)、file(文件消息)、link(图文链接消息)、miniprogram(小程序消息)、msgmenu(菜单消息)、location(地理位置消息)")
	_ = customerCmd.MarkFlagRequired(flags.MsgType)

	customerCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")
}
