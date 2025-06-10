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

// appCmd 企业微信应用消息
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "发送企业微信应用消息",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendAppParams{
			UserAgent:              variable.UserAgent,
			AccessToken:            variable.AccessToken,
			CorpID:                 variable.CorpID,
			CorpSecret:             variable.CorpSecret,
			ToUser:                 variable.ToUser,
			ToParty:                variable.ToParty,
			ToTag:                  variable.ToTag,
			AgentID:                variable.AgentID,
			MsgType:                variable.MsgType,
			Safe:                   variable.Safe,
			EnableIDTrans:          variable.EnableIDTrans,
			EnableDuplicateCheck:   variable.EnableDuplicateCheck,
			DuplicateCheckInterval: variable.DuplicateCheckInterval,
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

		if err := message.CmdWorkSendApp(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin app -i corp_id -s corp_secret -e agent_id -o '@all' -m text 'hello world'",
}

func init() {
	appCmd.Flags().SortFlags = false
	appCmd.AddCommand(undoAppCmd)

	workWeiXinSetAccessTokenFlags(appCmd)

	appCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "指定接收消息的成员，成员ID列表，最多支持1000个，多个接收者用‘|’分隔。指定为\"@all\"，则向该企业应用的全部成员发送")
	appCmd.Flags().StringVarP(&variable.ToParty, flags.ToParty, "p", "", "指定接收消息的部门，部门ID列表，最多支持100个，多个接收者用‘|’分隔。to_user\"@all\"时忽略本参数")
	appCmd.Flags().StringVarP(&variable.ToTag, flags.ToTag, "g", "", "指定接收消息的标签，标签ID列表，最多支持100个，多个接收者用‘|’分隔。to_user\"@all\"时忽略本参数")

	appCmd.Flags().Int64VarP(&variable.AgentID, flags.AgentID, "e", 0, "企业应用的id (必填)")
	_ = appCmd.MarkFlagRequired(flags.AgentID)

	appCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "消息类型 (必填)，text(文本消息)、image(图片消息)、voice(语音消息)、video(视频消息)、file(文件消息)、textcard(文本卡片消息)、news(图文消息)、mpnews(图文消息)、markdown(markdown消息)、miniprogram_notice(小程序通知消息)、template_card(模板卡片消息)")
	_ = appCmd.MarkFlagRequired(flags.MsgType)

	appCmd.Flags().IntVar(&variable.Safe, flags.Safe, 0, "表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0")
	appCmd.Flags().IntVarP(&variable.EnableIDTrans, flags.EnableIDTrans, "r", 0, "表示是否开启id转译，0表示否，1表示是，默认0。仅第三方应用需要用到，企业自建应用可以忽略")
	appCmd.Flags().IntVarP(&variable.EnableDuplicateCheck, flags.EnableDuplicateCheck, "c", 0, "表示是否开启重复消息检查，0表示否，1表示是，默认0")
	appCmd.Flags().IntVarP(&variable.DuplicateCheckInterval, flags.DuplicateCheckInterval, "d", 1800, "表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时(14400)")

	appCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")
}
