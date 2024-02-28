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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/cmd/dingtalk"
	"github.com/lenye/pmsg/cmd/feishu"
	"github.com/lenye/pmsg/cmd/slack"
	"github.com/lenye/pmsg/cmd/weixin"
	"github.com/lenye/pmsg/cmd/workweixin"
	"github.com/lenye/pmsg/pkg/version"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pmsg",
	Short: "publish message",
	Long: `发送消息的小工具：
    企业微信群机器人消息
    钉钉自定义机器人消息
    飞书自定义机器人消息
    Slack 机器人消息
    微信消息
    微信客服消息
    企业微信消息
    企业微信客服消息

    ` + version.OpenSource,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}`)
	rootCmd.Version = version.Print()

	rootCmd.AddCommand(weixin.Cmd)
	rootCmd.AddCommand(workweixin.Cmd)
	rootCmd.AddCommand(dingtalk.Cmd)
	rootCmd.AddCommand(feishu.Cmd)
	rootCmd.AddCommand(slack.Cmd)
}
