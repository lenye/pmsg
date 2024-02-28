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

package message

import "fmt"

// 小程序 msgtype 的合法值
const (
	MiniProgramMsgTypeText            = "text"            // 文本消息
	MiniProgramMsgTypeImage           = "image"           // 图片消息
	MiniProgramMsgTypeLink            = "link"            // 图文链接
	MiniProgramMsgTypeMiniProgramPage = "miniprogrampage" // 小程序卡片
)

// ValidateMiniProgramMsgType 验证
func ValidateMiniProgramMsgType(v string) error {
	switch v {
	case MiniProgramMsgTypeText, MiniProgramMsgTypeImage, MiniProgramMsgTypeLink, MiniProgramMsgTypeMiniProgramPage:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q]", v,
			MiniProgramMsgTypeText, MiniProgramMsgTypeImage, MiniProgramMsgTypeLink, MiniProgramMsgTypeMiniProgramPage)
	}
	return nil
}
