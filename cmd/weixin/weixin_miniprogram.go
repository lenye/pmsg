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
)

// miniProgramCmd 微信小程序
var miniProgramCmd = &cobra.Command{
	Use:     "miniprogram",
	Aliases: []string{"mini"},
	Short:   "weixin miniprogram",
}

func init() {
	miniProgramCmd.AddCommand(miniProgramCustomerCmd)
	miniProgramCmd.AddCommand(miniProgramSubCmd)
}
