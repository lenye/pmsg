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

// appChatCmd 企业微信群聊推送消息
var appChatCmd = &cobra.Command{
	Use:     "appchat",
	Aliases: []string{"chat"},
	Short:   "发送企业微信群聊推送消息",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendAppChatParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			CorpID:      variable.CorpID,
			CorpSecret:  variable.CorpSecret,
			ChatID:      variable.ChatID,
			MsgType:     variable.MsgType,
			Safe:        variable.Safe,
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

		if err := message.CmdWorkSendAppChat(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin appchat -i corp_id -s corp_secret -c chat_id -m text 'hello world'",
}

func init() {
	appChatCmd.Flags().SortFlags = false
	workWeiXinSetAccessTokenFlags(appChatCmd)

	appChatCmd.Flags().StringVarP(&variable.ChatID, flags.ChatID, "c", "", "群聊id (必填) ")
	_ = appChatCmd.MarkFlagRequired(flags.ChatID)

	appChatCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "消息类型 (必填)，text(文本消息)、image(图片消息)、voice(语音消息)、video(视频消息)、file(文件消息)、textcard(文本卡片消息)、news(图文消息)、mpnews(图文消息)、markdown(markdown消息)")
	_ = appChatCmd.MarkFlagRequired(flags.MsgType)

	appChatCmd.Flags().IntVar(&variable.Safe, flags.Safe, 0, "表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0")

	appChatCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")
}
