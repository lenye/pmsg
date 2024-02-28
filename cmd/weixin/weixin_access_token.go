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
	"github.com/lenye/pmsg/internal/im/weixin/token"
)

// accessTokenCmd 获取微信接口调用凭证
var accessTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "get weixin access token",
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
	accessTokenCmd.Flags().StringVarP(&variable.AppID, flags.AppID, "i", "", "weixin app id (required)")
	_ = accessTokenCmd.MarkFlagRequired(flags.AppID)

	accessTokenCmd.Flags().StringVarP(&variable.AppSecret, flags.AppSecret, "s", "", "weixin app Secret (required)")
	_ = accessTokenCmd.MarkFlagRequired(flags.AppSecret)
}
