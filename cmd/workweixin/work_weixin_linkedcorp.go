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
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/im"
	"github.com/lenye/pmsg/im/weixin/work/message"
)

// linkedCorpCmd 企业微信互联企业消息
var linkedCorpCmd = &cobra.Command{
	Use:     "linkedcorp",
	Aliases: []string{"lc"},
	Short:   "发送企业微信互联企业消息",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendLinkedCorpParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			CorpID:      variable.CorpID,
			CorpSecret:  variable.CorpSecret,
			ToAll:       variable.ToAll,
			AgentID:     variable.AgentID,
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

		if variable.ToUser != "" {
			arg.ToUser = strings.Split(variable.ToUser, "|")
		}
		if variable.ToParty != "" {
			arg.ToParty = strings.Split(variable.ToParty, "|")
		}
		if variable.ToTag != "" {
			arg.ToTag = strings.Split(variable.ToTag, "|")
		}

		if err := message.CmdWorkSendLinkedCorp(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin linkedcorp -i corp_id -s corp_secret -o 'userid1|userid2' -m text 'hello world'",
}

func init() {
	linkedCorpCmd.Flags().SortFlags = false
	workWeiXinSetAccessTokenFlags(linkedCorpCmd)

	linkedCorpCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "指定接收消息的成员，成员ID列表，最多支持1000个，多个接收者用‘|’分隔")
	linkedCorpCmd.Flags().StringVarP(&variable.ToParty, flags.ToParty, "p", "", "指定接收消息的部门，部门ID列表，最多支持100个，多个接收者用‘|’分隔")
	linkedCorpCmd.Flags().StringVarP(&variable.ToTag, flags.ToTag, "g", "", "指定接收消息的标签，标签ID列表，最多支持100个，多个接收者用‘|’分隔")
	linkedCorpCmd.Flags().IntVarP(&variable.ToAll, flags.ToAll, "l", 0, "1表示发送给应用可见范围内的所有人（包括互联企业的成员），默认为0")

	linkedCorpCmd.Flags().Int64VarP(&variable.AgentID, flags.AgentID, "e", 0, "企业应用的id (必填)")
	_ = linkedCorpCmd.MarkFlagRequired(flags.AgentID)

	linkedCorpCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "消息类型 (必填)，text(文本消息)、image(图片消息)、voice(语音消息)、video(视频消息)、file(文件消息)、textcard(文本卡片消息)、news(图文消息)、mpnews(图文消息)、markdown(markdown消息)、miniprogram_notice(小程序通知消息)")
	_ = linkedCorpCmd.MarkFlagRequired(flags.MsgType)

	linkedCorpCmd.Flags().IntVar(&variable.Safe, flags.Safe, 0, "表示是否是保密消息，0表示否，1表示是，默认0")

	linkedCorpCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")

}
