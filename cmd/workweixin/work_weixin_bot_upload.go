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
	"github.com/lenye/pmsg/flags"
	"github.com/lenye/pmsg/im"
	"github.com/lenye/pmsg/im/weixin/work/bot"
)

// botUploadCmd 企业微信群机器人上传文件
var botUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "企业微信群机器人上传文件",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdUploadParams{
			UserAgent: variable.UserAgent,
			Key:       variable.Secret,
		}

		if variable.IsRaw {
			arg.File = args[0]
		} else {
			var err error
			arg.File, err = im.StrRaw2Interpreted(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if err := bot.CmdUpload(&arg); err != nil {
			fmt.Println(err)
		}
	},
	Example: "pmsg workweixin bot upload -k key /img/app.png",
}

func init() {
	botUploadCmd.Flags().SortFlags = false
	workWeiXinBotSetKeyFlags(botUploadCmd)

	botUploadCmd.Flags().BoolVar(&variable.IsRaw, flags.IsRaw, false, "文件名称含路径是原始字符串字面值（不转义处理）")
}
