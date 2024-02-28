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
	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/internal/flags"
)

// Cmd 微信
var Cmd = &cobra.Command{
	Use:     "weixin",
	Aliases: []string{"wx"},
	Short:   "weixin",
}

func init() {
	Cmd.PersistentFlags().StringVarP(&variable.UserAgent, flags.UserAgent, "a", "", "http user agent")

	Cmd.AddCommand(accessTokenCmd)
	Cmd.AddCommand(miniProgramCmd)
	Cmd.AddCommand(officialAccountCmd)
	Cmd.AddCommand(mediaUploadCmd)

}

// weiXinSetAccessTokenFlags 设置微信access_token或者app_id/app_secret命令行参数
func weiXinSetAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&variable.AccessToken, flags.AccessToken, "t", "", "weixin access token")

	cmd.Flags().StringVarP(&variable.AppID, flags.AppID, "i", "", "weixin app id (required if app Secret is set)")
	cmd.Flags().StringVarP(&variable.AppSecret, flags.AppSecret, "s", "", "weixin app Secret (required if app id is set)")

	cmd.MarkFlagsMutuallyExclusive(flags.AccessToken, flags.AppID)
	cmd.MarkFlagsRequiredTogether(flags.AppID, flags.AppSecret)
}
