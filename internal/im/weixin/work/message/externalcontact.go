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

// 企业微信家校消息 类型
const (
	ExternalContactMsgTypeText              = "text"               // 文本消息
	ExternalContactMsgTypeImage             = "image"              // 图片消息
	ExternalContactMsgTypeVoice             = "voice"              // 语音消息
	ExternalContactMsgTypeVideo             = "video"              // 视频消息
	ExternalContactMsgTypeFile              = "file"               // 文件消息
	ExternalContactMsgTypeNews              = "news"               // 图文消息
	ExternalContactMsgTypeMpNews            = "mpnews"             // 图文消息
	ExternalContactMsgTypeMiniProgramNotice = "miniprogram_notice" // 小程序通知消息
)

// ValidateExternalContactMsgType 验证
func ValidateExternalContactMsgType(v string) error {
	switch v {
	case ExternalContactMsgTypeText, ExternalContactMsgTypeImage, ExternalContactMsgTypeVoice, ExternalContactMsgTypeVideo,
		ExternalContactMsgTypeFile, ExternalContactMsgTypeNews, ExternalContactMsgTypeMpNews,
		ExternalContactMsgTypeMiniProgramNotice:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q %q]", v,
			ExternalContactMsgTypeText, ExternalContactMsgTypeImage, ExternalContactMsgTypeVoice, ExternalContactMsgTypeVideo,
			ExternalContactMsgTypeFile, ExternalContactMsgTypeNews, ExternalContactMsgTypeMpNews,
			ExternalContactMsgTypeMiniProgramNotice)
	}
	return nil
}

// ExternalContactMessage 企业微信家校消息
type ExternalContactMessage struct {
	RecvScope              int                    `json:"recv_scope,omitempty"`               // 指定发送对象，0表示发送给家长，1表示发送给学生，2表示发送给家长和学生，默认为0。
	ToParentUserID         []string               `json:"to_parent_userid,omitempty"`         // recv_scope为0或2表示发送给对应的家长，recv_scope为1忽略，（最多支持1000个）
	ToStudentUserID        []string               `json:"to_student_userid,omitempty"`        // recv_scope为0表示发送给学生的所有家长，recv_scope为1表示发送给学生，recv_scope为2表示发送给学生和学生的所有家长（最多支持1000个）
	ToParty                []string               `json:"to_party,omitempty"`                 // recv_scope为0表示发送给班级的所有家长，recv_scope为1表示发送给班级的所有学生，recv_scope为2表示发送给班级的所有学生和家长（最多支持100个）
	ToAll                  int                    `json:"toall,omitempty"`                    // 1表示字段生效，0表示字段无效。recv_scope为0表示发送给学校的所有家长，recv_scope为1表示发送给学校的所有学生，recv_scope为2表示发送给学校的所有学生和家长，默认为0
	MsgType                string                 `json:"msgtype"`                            // 消息类型
	AgentID                int64                  `json:"agentid"`                            // 企业应用的id
	EnableIDTrans          int                    `json:"enable_id_trans,omitempty"`          // 表示是否开启id转译，0表示否，1表示是，默认0
	EnableDuplicateCheck   int                    `json:"enable_duplicate_check,omitempty"`   // 表示是否开启重复消息检查，0表示否，1表示是，默认0
	DuplicateCheckInterval int                    `json:"duplicate_check_interval,omitempty"` // 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	Text                   *TextMeta              `json:"text,omitempty"`                     // 文本消息
	Image                  *ImageMeta             `json:"image,omitempty"`                    // 图片消息
	Voice                  *VoiceMeta             `json:"voice,omitempty"`                    // 语音消息
	Video                  *VideoMeta             `json:"video,omitempty"`                    // 视频消息
	File                   *FileMeta              `json:"file,omitempty"`                     // 文件消息
	News                   *NewsMeta              `json:"news,omitempty"`                     // 图文消息
	MpNews                 *MpNewsMeta            `json:"mpnews,omitempty"`                   // 图文消息
	MiniProgramNotice      *MiniProgramNoticeMeta `json:"miniprogram_notice,omitempty"`       // 小程序通知消息
}

type ExternalContactMessageResponse struct {
	weixin.ResponseMeta
	InvalidParentUserID  []string `json:"invalid_parent_userid,omitempty"`
	InvalidStudentUserID []string `json:"invalid_student_userid,omitempty"`
	InvalidParty         []string `json:"invalid_party,omitempty"`
}

func (t ExternalContactMessageResponse) String() string {
	var sb []string

	if t.InvalidParentUserID != nil && len(t.InvalidParentUserID) != 0 {
		sb = append(sb, fmt.Sprintf("invalid_parent_userid: %q", strings.Join(t.InvalidParentUserID, ",")))
	}
	if t.InvalidStudentUserID != nil && len(t.InvalidStudentUserID) != 0 {
		sb = append(sb, fmt.Sprintf("invalid_student_userid: %q", strings.Join(t.InvalidStudentUserID, ",")))
	}
	if t.InvalidParty != nil && len(t.InvalidParty) != 0 {
		sb = append(sb, fmt.Sprintf("invalid_party: %q", strings.Join(t.InvalidParty, ",")))
	}

	if !t.ResponseMeta.Succeed() {
		sb = append([]string{fmt.Sprintf("%v", t.ResponseMeta)}, sb...)
	}
	return strings.Join(sb, ", ")
}

const externalContactSendURL = work.Host + "/cgi-bin/externalcontact/message/send?access_token="

// SendExternalContact 发送企业微信家校消息
func SendExternalContact(accessToken string, msg *ExternalContactMessage) (*ExternalContactMessageResponse, error) {
	u := externalContactSendURL + url.QueryEscape(accessToken)
	var resp ExternalContactMessageResponse
	_, err := client.PostJSON(u, msg, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; %v", weixin.ErrRequest, resp)
	}
	return &resp, nil
}
