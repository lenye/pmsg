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
	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
)

// weiXinCmd 微信
var weiXinCmd = &cobra.Command{
	Use:     "weixin",
	Aliases: []string{"wx"},
	Short:   "weixin message",
}

func init() {
	rootCmd.AddCommand(weiXinCmd)

	weiXinCmd.PersistentFlags().StringVarP(&userAgent, flags.UserAgent, "a", "", "http user agent")
}

// weiXinSetAccessTokenFlags 设置微信access_token或者app_id/app_secret命令行参数
func weiXinSetAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&accessToken, flags.AccessToken, "t", "", "weixin access token")

	cmd.Flags().StringVarP(&appID, flags.AppID, "i", "", "weixin app id (required if app secret is set)")
	cmd.Flags().StringVarP(&appSecret, flags.AppSecret, "s", "", "weixin app secret (required if app id is set)")

	cmd.MarkFlagsMutuallyExclusive(flags.AccessToken, flags.AppID)
	cmd.MarkFlagsRequiredTogether(flags.AppID, flags.AppSecret)
}
