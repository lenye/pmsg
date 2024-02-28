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
	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin/offiaccount/message"
	"github.com/lenye/pmsg/pkg/helper"
)

// officialAccountTplCmd 微信公众号模板消息
var officialAccountTplCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "publish weixin official account template message",
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
			arg.Data, err = helper.StrRaw2Interpreted(args[0])
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
	officialAccountTplCmd.AddCommand(officialAccountTplSubCmd)

	weiXinSetAccessTokenFlags(officialAccountTplCmd)

	officialAccountTplCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "weixin user open id (required)")
	_ = officialAccountTplCmd.MarkFlagRequired(flags.ToUser)

	officialAccountTplCmd.Flags().StringVarP(&variable.TemplateID, flags.TemplateID, "p", "", "weixin template id (required)")
	_ = officialAccountTplCmd.MarkFlagRequired(flags.TemplateID)

	officialAccountTplCmd.Flags().StringVar(&variable.Url, flags.Url, "", "Url")
	officialAccountTplCmd.Flags().StringToStringVar(&variable.Mini, flags.Mini, nil, "weixin Mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")

	officialAccountTplCmd.Flags().StringVar(&variable.Color, flags.Color, "", "template Color")
	officialAccountTplCmd.Flags().StringVarP(&variable.ClientMsgID, flags.ClientMsgID, "c", "", "client message id")

	officialAccountTplCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")
}
