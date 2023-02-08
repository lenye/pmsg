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
	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
)

// workWeiXinCmd 企业微信
var workWeiXinCmd = &cobra.Command{
	Use:     "workweixin",
	Aliases: []string{"wwx"},
	Short:   "work weixin",
}

func init() {
	workWeiXinCmd.PersistentFlags().StringVarP(&userAgent, flags.UserAgent, "a", "", "http user agent")

	workWeiXinCmd.AddCommand(workWeiXinAccessTokenCmd)
	workWeiXinCmd.AddCommand(workWeiXinAppCmd)
	workWeiXinCmd.AddCommand(workWeiXinAppChatCmd)
	workWeiXinCmd.AddCommand(workWeiXinCustomerCmd)
	workWeiXinCmd.AddCommand(workWeiXinExternalContactCmd)
	workWeiXinCmd.AddCommand(workWeiXinLinkedCorpCmd)
	workWeiXinCmd.AddCommand(workWeiXinMediaUploadCmd)
	workWeiXinCmd.AddCommand(workWeiXinBotCmd)

}

// workWeiXinSetAccessTokenFlags 设置企业微信access_token或者corp_id/corp_secret命令行参数
func workWeiXinSetAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&accessToken, flags.AccessToken, "t", "", "work weixin access token")

	cmd.Flags().StringVarP(&corpID, flags.CorpID, "i", "", "work weixin corp id (required if corp secret is set)")
	cmd.Flags().StringVarP(&corpSecret, flags.CorpSecret, "s", "", "work weixin corp secret (required if corp id is set)")

	cmd.MarkFlagsMutuallyExclusive(flags.AccessToken, flags.CorpID)
	cmd.MarkFlagsRequiredTogether(flags.CorpID, flags.CorpSecret)
}
