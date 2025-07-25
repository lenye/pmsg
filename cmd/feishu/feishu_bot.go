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

package feishu

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/im"
	"github.com/lenye/pmsg/im/feishu/bot"
)

// botCmd 飞书自定义机器人
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "发送飞书自定义机器人消息",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			Secret:      variable.Secret,
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

		if err := bot.CmdSend(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg feishu bot -t access_token -m text 'hello world'",
}

func init() {
	botCmd.Flags().SortFlags = false
	botCmd.Flags().StringVarP(&variable.AccessToken, flags.AccessToken, "t", "", "飞书自定义机器人 token (必填), token 为 webhook 地址中 xxx 部分 https://open.feishu.cn/open-apis/bot/v2/hook/xxx")
	_ = botCmd.MarkFlagRequired(flags.AccessToken)

	botCmd.Flags().StringVarP(&variable.Secret, flags.Secret, "s", "", "签名密钥")

	botCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "消息类型 (必填)，text(文本消息)、post(富文本)、image(图片)、share_chat(分享群名片)、interactive(消息卡片)")
	_ = botCmd.MarkFlagRequired(flags.MsgType)

	botCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")

}
