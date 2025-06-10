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
	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/im/weixin/token"
)

// accessTokenCmd 获取微信接口调用凭证
var accessTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "获取微信（公众号、小程序）接口调用凭证 access token",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arg := token.CmdTokenParams{
			UserAgent: variable.UserAgent,
			AppID:     variable.AppID,
			AppSecret: variable.AppSecret,
		}
		if err := token.CmdGetAccessToken(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg weixin token -i app_id -s app_secret",
}

func init() {
	accessTokenCmd.Flags().SortFlags = false

	accessTokenCmd.Flags().StringVarP(&variable.AppID, flags.AppID, "i", "", "微信 app id (必填)")
	_ = accessTokenCmd.MarkFlagRequired(flags.AppID)

	accessTokenCmd.Flags().StringVarP(&variable.AppSecret, flags.AppSecret, "s", "", "微信 app secret (必填)")
	_ = accessTokenCmd.MarkFlagRequired(flags.AppSecret)
}
