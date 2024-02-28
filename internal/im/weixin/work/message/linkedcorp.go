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

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/internal/im/weixin/client"
	"github.com/lenye/pmsg/internal/im/weixin/work"
)

// 企业微信互联企业消息 类型
const (
	LinkedCorpMsgTypeText              = "text"               // 文本消息
	LinkedCorpMsgTypeImage             = "image"              // 图片消息
	LinkedCorpMsgTypeVoice             = "voice"              // 语音消息
	LinkedCorpMsgTypeVideo             = "video"              // 视频消息
	LinkedCorpMsgTypeFile              = "file"               // 文件消息
	LinkedCorpMsgTypeTextCard          = "textcard"           // 文本卡片消息
	LinkedCorpMsgTypeNews              = "news"               // 图文消息
	LinkedCorpMsgTypeMpNews            = "mpnews"             // 图文消息
	LinkedCorpMsgTypeMarkdown          = "markdown"           // markdown消息
	LinkedCorpMsgTypeMiniProgramNotice = "miniprogram_notice" // 小程序通知消息
)

// ValidateLinkedCorpMsgType 验证
func ValidateLinkedCorpMsgType(v string) error {
	switch v {
	case LinkedCorpMsgTypeText, LinkedCorpMsgTypeImage, LinkedCorpMsgTypeVoice, LinkedCorpMsgTypeVideo,
		LinkedCorpMsgTypeFile, LinkedCorpMsgTypeTextCard, LinkedCorpMsgTypeNews, LinkedCorpMsgTypeMpNews,
		LinkedCorpMsgTypeMarkdown, LinkedCorpMsgTypeMiniProgramNotice:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q %q %q %q]", v,
			LinkedCorpMsgTypeText, LinkedCorpMsgTypeImage, LinkedCorpMsgTypeVoice, LinkedCorpMsgTypeVideo,
			LinkedCorpMsgTypeFile, LinkedCorpMsgTypeTextCard, LinkedCorpMsgTypeNews, LinkedCorpMsgTypeMpNews,
			LinkedCorpMsgTypeMarkdown, LinkedCorpMsgTypeMiniProgramNotice)
	}
	return nil
}

// LinkedCorpMessage 企业微信互联企业消息
type LinkedCorpMessage struct {
	ToUser            []string               `json:"touser,omitempty"`             // 成员ID列表（消息接收者，最多支持1000个）。每个元素的格式为： corpid/userid，其中， corpid为该互联成员所属的企业，userid为该互联成员所属企业中的帐号。如果是本企业的成员，则直接传userid即可
	ToParty           []string               `json:"toparty,omitempty"`            // 部门ID列表，最多支持100个。partyid在互联圈子内唯一。每个元素都是字符串类型，格式为：linked_id/party_id，其中linked_id是互联id，party_id是在互联圈子中的部门id。如果是本企业的部门，则直接传party_id即可
	ToTag             []string               `json:"totag,omitempty"`              // 本企业的标签ID列表，最多支持100个
	ToAll             int                    `json:"toall,omitempty"`              // 1表示发送给应用可见范围内的所有人（包括互联企业的成员），默认为0
	AgentID           int64                  `json:"agentid"`                      // 企业应用的id，整型。可在应用的设置页面查看
	MsgType           string                 `json:"msgtype"`                      // 消息类型
	Safe              int                    `json:"safe,omitempty"`               // 表示是否是保密消息，0表示否，1表示是，默认0
	Text              *TextMeta              `json:"text,omitempty"`               // 文本消息
	Image             *ImageMeta             `json:"image,omitempty"`              // 图片消息
	Voice             *VoiceMeta             `json:"voice,omitempty"`              // 语音消息
	Video             *VideoMeta             `json:"video,omitempty"`              // 视频消息
	File              *FileMeta              `json:"file,omitempty"`               // 文件消息
	TextCard          *TextCardMeta          `json:"textcard,omitempty"`           // 文本卡片消息
	News              *NewsMeta              `json:"news,omitempty"`               // 图文消息
	MpNews            *MpNewsMeta            `json:"mpnews,omitempty"`             // 图文消息
	Markdown          *MarkdownMeta          `json:"markdown,omitempty"`           // markdown消息
	MiniProgramNotice *MiniProgramNoticeMeta `json:"miniprogram_notice,omitempty"` // 小程序通知消息
}

type LinkedCorpMessageResponse struct {
	weixin.ResponseMeta
	InvalidUser  []string `json:"invaliduser,omitempty"`
	InvalidParty []string `json:"invalidparty,omitempty"`
	InvalidTag   []string `json:"invalidtag,omitempty"`
}

func (t LinkedCorpMessageResponse) String() string {
	var sb []string

	if t.InvalidUser != nil && len(t.InvalidUser) != 0 {
		sb = append(sb, fmt.Sprintf("invaliduser: %q", strings.Join(t.InvalidUser, ",")))
	}
	if t.InvalidParty != nil && len(t.InvalidParty) != 0 {
		sb = append(sb, fmt.Sprintf("invalidparty: %q", strings.Join(t.InvalidParty, ",")))
	}
	if t.InvalidTag != nil && len(t.InvalidTag) != 0 {
		sb = append(sb, fmt.Sprintf("invalidtag: %q", strings.Join(t.InvalidTag, ",")))
	}

	if !t.ResponseMeta.Succeed() {
		sb = append([]string{fmt.Sprintf("%v", t.ResponseMeta)}, sb...)
	}
	return strings.Join(sb, ", ")
}

const linkedCorpSendURL = work.Host + "/cgi-bin/linkedcorp/message/send?access_token="

// SendLinkedCorp 发送企业微信互联企业消息
func SendLinkedCorp(accessToken string, msg *LinkedCorpMessage) (*LinkedCorpMessageResponse, error) {
	u := linkedCorpSendURL + url.QueryEscape(accessToken)
	var resp LinkedCorpMessageResponse
	_, err := client.PostJSON(u, msg, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; %v", weixin.ErrRequest, resp)
	}
	return &resp, nil
}
