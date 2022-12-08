// Copyright 2022 The pmsg Authors. All rights reserved.
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
	"github.com/lenye/pmsg/pkg/weixin/work/webhook"
)

// weiXinWorkWebHookMessageCmd 推送企业微信群机器人消息
var weiXinWorkWebHookMessageCmd = &cobra.Command{
	Use:     "webhooksend",
	Aliases: []string{"whs"},
	Short:   "publish work weixin webhook message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := webhook.CmdSendParams{
			UserAgent: userAgent,
			Key:       key,
			MsgType:   msgType,
			ToUser:    toUser,
			ToMobile:  toMobile,
			Data:      args[0],
		}
		if err := webhook.CmdSend(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkSetKeyFlags(weiXinWorkWebHookMessageCmd)

	weiXinWorkWebHookMessageCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	weiXinWorkWebHookMessageCmd.MarkFlagRequired(flags.MsgType)

	weiXinWorkWebHookMessageCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "work weixin user id list")
	weiXinWorkWebHookMessageCmd.Flags().StringVarP(&toMobile, flags.ToMobile, "b", "", "mobile list")
}
