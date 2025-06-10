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

// externalContactCmd 企业微信家校消息
var externalContactCmd = &cobra.Command{
	Use:     "externalcontact",
	Aliases: []string{"ec"},
	Short:   "发送企业微信学校通知消息",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendExternalContactParams{
			UserAgent:              variable.UserAgent,
			AccessToken:            variable.AccessToken,
			CorpID:                 variable.CorpID,
			CorpSecret:             variable.CorpSecret,
			RecvScope:              variable.RecvScope,
			ToAll:                  variable.ToAll,
			MsgType:                variable.MsgType,
			AgentID:                variable.AgentID,
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

		if variable.ToParentUserID != "" {
			arg.ToParentUserID = strings.Split(variable.ToParentUserID, "|")
		}
		if variable.ToStudentUserID != "" {
			arg.ToStudentUserID = strings.Split(variable.ToStudentUserID, "|")
		}
		if variable.ToParty != "" {
			arg.ToParty = strings.Split(variable.ToParty, "|")
		}

		if err := message.CmdWorkSendExternalContact(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin externalcontact -i corp_id -s corp_secret -e agent_id -n 'parentuserid1|parentuserid2' -m text 'hello world'",
}

func init() {
	externalContactCmd.Flags().SortFlags = false
	workWeiXinSetAccessTokenFlags(externalContactCmd)

	externalContactCmd.Flags().IntVarP(&variable.RecvScope, flags.RecvScope, "o", 0, "指定发送对象，0表示发送给家长，1表示发送给学生，2表示发送给家长和学生，默认为0")

	externalContactCmd.Flags().StringVarP(&variable.ToParentUserID, flags.ToParentUserID, "n", "", "recv_scope为0或2表示发送给对应的家长，recv_scope为1忽略，最多支持1000个，多个接收者用‘|’分隔")
	externalContactCmd.Flags().StringVarP(&variable.ToStudentUserID, flags.ToStudentUserID, "u", "", "recv_scope为0表示发送给学生的所有家长，recv_scope为1表示发送给学生，recv_scope为2表示发送给学生和学生的所有家长，最多支持1000个，多个接收者用‘|’分隔")
	externalContactCmd.Flags().StringVarP(&variable.ToParty, flags.ToParty, "p", "", "recv_scope为0表示发送给班级的所有家长，recv_scope为1表示发送给班级的所有学生，recv_scope为2表示发送给班级的所有学生和家长，最多支持100个，多个接收者用‘|’分隔")
	externalContactCmd.Flags().IntVarP(&variable.ToAll, flags.ToAll, "l", 0, "1表示字段生效，0表示字段无效。recv_scope为0表示发送给学校的所有家长，recv_scope为1表示发送给学校的所有学生，recv_scope为2表示发送给学校的所有学生和家长，默认为0")

	externalContactCmd.Flags().Int64VarP(&variable.AgentID, flags.AgentID, "e", 0, "企业应用的 id (必填")
	_ = externalContactCmd.MarkFlagRequired(flags.AgentID)

	externalContactCmd.Flags().StringVarP(&variable.MsgType, flags.MsgType, "m", "", "消息类型 (必填)，text(文本消息)、image(图片消息)、voice(语音消息)、video(视频消息)、file(文件消息)、news(图文消息)、mpnews(图文消息)、miniprogram_notice(小程序通知消息)")
	_ = externalContactCmd.MarkFlagRequired(flags.MsgType)

	externalContactCmd.Flags().IntVarP(&variable.EnableIDTrans, flags.EnableIDTrans, "r", 0, "表示是否开启id转译，0表示否，1表示是，默认0。仅第三方应用需要用到，企业自建应用可以忽略")
	externalContactCmd.Flags().IntVarP(&variable.EnableDuplicateCheck, flags.EnableDuplicateCheck, "c", 0, "表示是否开启重复消息检查，0表示否，1表示是，默认0")
	externalContactCmd.Flags().IntVarP(&variable.DuplicateCheckInterval, flags.DuplicateCheckInterval, "d", 1800, "表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时(14400)")

	externalContactCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "消息内容是原始字符串字面值（不转义处理）")

}
