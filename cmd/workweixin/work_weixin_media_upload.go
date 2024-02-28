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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/variable"
	"github.com/lenye/pmsg/internal/flags"
	"github.com/lenye/pmsg/internal/im/weixin/work/asset"
	"github.com/lenye/pmsg/pkg/helper"
)

// mediaUploadCmd 企业微信上传临时素材
var mediaUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "work weixin media upload",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := asset.CmdWorkMediaUploadParams{
			UserAgent:   variable.UserAgent,
			AccessToken: variable.AccessToken,
			CorpID:      variable.CorpID,
			CorpSecret:  variable.CorpSecret,
			MediaType:   variable.MediaType,
		}

		if variable.IsRaw {
			arg.File = args[0]
		} else {
			var err error
			arg.File, err = helper.StrRaw2Interpreted(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if err := asset.CmdWorkMediaUpload(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin upload -i corp_id -m image /img/app.png",
}

func init() {
	workWeiXinSetAccessTokenFlags(mediaUploadCmd)

	mediaUploadCmd.Flags().StringVarP(&variable.MediaType, flags.MediaType, "m", "", "media type (required)")
	_ = mediaUploadCmd.MarkFlagRequired(flags.MediaType)

	mediaUploadCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "strings without any escape processing")
}
