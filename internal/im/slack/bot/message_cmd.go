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

package bot

import (
	"fmt"

	"github.com/lenye/pmsg/internal/im/slack"
	"github.com/lenye/pmsg/pkg/httpclient"
)

type CmdSendParams struct {
	UserAgent string
	URL       string
	Data      string
}

// CmdSend 发送消息
func CmdSend(arg *CmdSendParams) error {

	httpclient.SetUserAgent(arg.UserAgent)

	if err := Send(arg.URL, arg.Data); err != nil {
		return err
	}
	fmt.Println(slack.MessageOK)

	return nil
}
