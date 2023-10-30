// Copyright 2022-2023 The pmsg Authors. All rights reserved.
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
	"github.com/lenye/pmsg/pkg/conv"
)

// officialAccountTplSubCmd 微信公众号一次性订阅消息
var officialAccountTplSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin official account template subscribe message (onetime)",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := conv.StrRaw2Interpreted(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		arg := message.CmdMpSendTemplateSubscribeParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			AppID:       variable.AppID,
			AppSecret:   variable.AppSecret,
			ToUser:      variable.ToUser,
			TemplateID:  variable.TemplateID,
			Scene:       variable.Scene,
			Title:       variable.Title,
			Url:         variable.Url,
			Mini:        variable.Mini,
			Data:        data,
		}
		if err := message.CmdMpSendTemplateSubscribe(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg weixin offiaccount template subscribe -i app_id -s app_secret --Scene Scene --Title Title -p template_id -o open_id '{\"first\":{\"value\":\"test\"}}'",
}

func init() {
	weiXinSetAccessTokenFlags(officialAccountTplSubCmd)

	officialAccountTplSubCmd.Flags().StringVarP(&variable.ToUser, flags.ToUser, "o", "", "weixin user open id (required)")
	officialAccountTplSubCmd.MarkFlagRequired(flags.ToUser)

	officialAccountTplSubCmd.Flags().StringVarP(&variable.TemplateID, flags.TemplateID, "p", "", "weixin template id (required)")
	officialAccountTplSubCmd.MarkFlagRequired(flags.TemplateID)

	officialAccountTplSubCmd.Flags().StringVar(&variable.Scene, flags.Scene, "", "weixin subscribe Scene (required)")
	officialAccountTplSubCmd.MarkFlagRequired(flags.Scene)

	officialAccountTplSubCmd.Flags().StringVar(&variable.Title, flags.Title, "", "message Title (required)")
	officialAccountTplSubCmd.MarkFlagRequired(flags.Title)

	officialAccountTplSubCmd.Flags().StringVar(&variable.Url, flags.Url, "", "Url")
	officialAccountTplSubCmd.Flags().StringToStringVar(&variable.Mini, flags.Mini, nil, "weixin Mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}
