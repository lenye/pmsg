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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/offiaccount/message"
)

// weiXinOfficialAccountTplCmd 微信公众号模板消息
var weiXinOfficialAccountTplCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "publish weixin official account template message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendTemplateParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			AppID:       appID,
			AppSecret:   appSecret,
			ToUser:      toUser,
			TemplateID:  templateID,
			Url:         url,
			Mini:        mini,
			Color:       color,
			ClientMsgID: clientMsgID,
			Data:        args[0],
		}
		if err := message.CmdMpSendTemplate(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg weixin offiaccount template -i app_id -s app_secret -p template_id -o open_id '{\"first\":{\"value\":\"test\"}}'",
}

func init() {
	weiXinOfficialAccountTplCmd.AddCommand(weiXinOfficialAccountTplSubCmd)

	weiXinSetAccessTokenFlags(weiXinOfficialAccountTplCmd)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountTplCmd.MarkFlagRequired(flags.ToUser)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&templateID, flags.TemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountTplCmd.MarkFlagRequired(flags.TemplateID)

	weiXinOfficialAccountTplCmd.Flags().StringVar(&url, flags.Url, "", "url")
	weiXinOfficialAccountTplCmd.Flags().StringToStringVar(&mini, flags.Mini, nil, "weixin mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")

	weiXinOfficialAccountTplCmd.Flags().StringVar(&color, flags.Color, "", "template color")
	weiXinOfficialAccountTplCmd.Flags().StringVarP(&clientMsgID, flags.ClientMsgID, "c", "", "client message id")
}
