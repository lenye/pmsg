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
	"github.com/lenye/pmsg/im/weixin/work/bot"
)

// botCmd 企业微信群机器人
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "发送企业微信群机器人消息",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent: variable.UserAgent,
			Key:       variable.Secret,
			MsgType:   variable.MsgType,
			AtUser:    variable.AtUser,
			AtMobile:  variable.AtMobile,
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
	Example: "pmsg workweixin bot -k key -m text 'hello world'",
}

func init() {
	botCmd.Flags().SortFlags = false
	botCmd.AddCommand(botUploadCmd)

	workWeiXinBotSetKeyFlags(botCmd)

	botCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "消息类型 (必填)，text(文本消息)、markdown(markdown消息)、image(图片消息)、news(图文消息)、file(文件消息)、text_notice(文本通知模版卡片)、news_notice(图文展示模版卡片)")
	_ = botCmd.MarkFlagRequired(flags.MsgType)

	botCmd.Flags().StringVarP(&variable.AtUser, flags.AtUser, "o", "", "文本消息时，提醒群中的指定成员(@某个成员)，多个接收者用‘|’分隔，@all表示提醒所有人，如果开发者获取不到userid，可以使用at_mobile")
	botCmd.Flags().StringVarP(&variable.AtMobile, flags.AtMobile, "b", "", "文本消息时，提醒手机号对应的群成员(@某个成员)，多个接收者用‘|’分隔，@all表示提醒所有人")

	botCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")
}

// workWeiXinBotSetKeyFlags 设置企业微信群机器人key命令行参数
func workWeiXinBotSetKeyFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&variable.Secret, flags.Key, "k", "", "企业微信群机器人 key (必填)")
	_ = cmd.MarkFlagRequired(flags.Key)
}
