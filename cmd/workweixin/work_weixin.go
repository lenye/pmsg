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
	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/internal/flags"
)

// Cmd 企业微信
var Cmd = &cobra.Command{
	Use:     "workweixin",
	Aliases: []string{"wwx"},
	Short:   "work weixin",
}

func init() {
	Cmd.PersistentFlags().StringVarP(&variable.UserAgent, flags.UserAgent, "a", "", "http user agent")

	Cmd.AddCommand(accessTokenCmd)
	Cmd.AddCommand(appCmd)
	Cmd.AddCommand(appChatCmd)
	Cmd.AddCommand(customerCmd)
	Cmd.AddCommand(externalContactCmd)
	Cmd.AddCommand(linkedCorpCmd)
	Cmd.AddCommand(mediaUploadCmd)
	Cmd.AddCommand(botCmd)

}

// workWeiXinSetAccessTokenFlags 设置企业微信access_token或者corp_id/corp_secret命令行参数
func workWeiXinSetAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&variable.AccessToken, flags.AccessToken, "t", "", "work weixin access token")

	cmd.Flags().StringVarP(&variable.CorpID, flags.CorpID, "i", "", "work weixin corp id (required if corp Secret is set)")
	cmd.Flags().StringVarP(&variable.CorpSecret, flags.CorpSecret, "s", "", "work weixin corp Secret (required if corp id is set)")

	cmd.MarkFlagsMutuallyExclusive(flags.AccessToken, flags.CorpID)
	cmd.MarkFlagsRequiredTogether(flags.CorpID, flags.CorpSecret)
}
