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

// 公众号 msgtype 的合法值
const (
	MpMsgTypeText            = "text"            // 文本消息
	MpMsgTypeImage           = "image"           // 图片消息
	MpMsgTypeVoice           = "voice"           // 语音消息
	MpMsgTypeVideo           = "video"           // 视频消息
	MpMsgTypeMusic           = "music"           // 音乐消息
	MpMsgTypeNews            = "news"            // 图文消息（点击跳转到外链）
	MpMsgTypeMpNews          = "mpnews"          // 图文消息（点击跳转到图文消息页面）
	MpMsgTypeMpNewsArticle   = "mpnewsarticle"   // 图文消息（点击跳转到图文消息页面）使用通过 “发布” 系列接口得到的 article_id
	MpMsgTypeMsgMenu         = "msgmenu"         // 菜单消息
	MpMsgTypeWxCard          = "wxcard"          // 卡券
	MpMsgTypeMiniProgramPage = "miniprogrampage" // 小程序卡片（要求小程序与公众号已关联）
)

// ValidateMpMsgType 验证
func ValidateMpMsgType(v string) error {
	switch v {
	case MpMsgTypeText, MpMsgTypeImage, MpMsgTypeVoice, MpMsgTypeVideo,
		MpMsgTypeMusic, MpMsgTypeNews, MpMsgTypeMpNews, MpMsgTypeMpNewsArticle,
		MpMsgTypeMsgMenu, MpMsgTypeWxCard, MpMsgTypeMiniProgramPage:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q %q %q %q %q]", v,
			MpMsgTypeText, MpMsgTypeImage, MpMsgTypeVoice, MpMsgTypeVideo,
			MpMsgTypeMusic, MpMsgTypeNews, MpMsgTypeMpNews, MpMsgTypeMpNewsArticle,
			MpMsgTypeMsgMenu, MpMsgTypeWxCard, MpMsgTypeMiniProgramPage)
	}
	return nil
}
