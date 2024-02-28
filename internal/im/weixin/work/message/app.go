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

// 企业微信应用消息 类型
const (
	AppMsgTypeText              = "text"               // 文本消息
	AppMsgTypeImage             = "image"              // 图片消息
	AppMsgTypeVoice             = "voice"              // 语音消息
	AppMsgTypeVideo             = "video"              // 视频消息
	AppMsgTypeFile              = "file"               // 文件消息
	AppMsgTypeTextCard          = "textcard"           // 文本卡片消息
	AppMsgTypeNews              = "news"               // 图文消息
	AppMsgTypeMpNews            = "mpnews"             // 图文消息
	AppMsgTypeMarkdown          = "markdown"           // markdown消息
	AppMsgTypeMiniProgramNotice = "miniprogram_notice" // 小程序通知消息
	AppMsgTypeTemplateCard      = "template_card"      // 模板卡片消息
)

// ValidateAppMsgType 验证
func ValidateAppMsgType(v string) error {
	switch v {
	case AppMsgTypeText, AppMsgTypeImage, AppMsgTypeVoice, AppMsgTypeVideo,
		AppMsgTypeFile, AppMsgTypeTextCard, AppMsgTypeNews, AppMsgTypeMpNews,
		AppMsgTypeMarkdown, AppMsgTypeMiniProgramNotice, AppMsgTypeTemplateCard:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q %q %q %q %q]", v,
			AppMsgTypeText, AppMsgTypeImage, AppMsgTypeVoice, AppMsgTypeVideo,
			AppMsgTypeFile, AppMsgTypeTextCard, AppMsgTypeNews, AppMsgTypeMpNews,
			AppMsgTypeMarkdown, AppMsgTypeMiniProgramNotice, AppMsgTypeTemplateCard)
	}
	return nil
}

// 企业微信应用的模板卡片消息 类型
const (
	AppTplCardTypeTextNotice        = "text_notice"        // 文本通知型
	AppTplCardTypeNewsNotice        = "news_notice"        // 图文展示型
	AppTplCardTypeButtonInteraction = "button_interaction" // 按钮交互型
	AppTplCardTypeVoteInteraction   = "vote_interaction"   // 投票选择型
)

// ValidateAppTemplateCardType 验证
func ValidateAppTemplateCardType(v string) error {
	switch v {
	case AppTplCardTypeTextNotice, AppTplCardTypeNewsNotice, AppTplCardTypeButtonInteraction, AppTplCardTypeVoteInteraction:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q]", v,
			AppTplCardTypeTextNotice, AppTplCardTypeNewsNotice, AppTplCardTypeButtonInteraction, AppTplCardTypeVoteInteraction)
	}
	return nil
}

// AppMessage 企业微信应用消息 touser、toparty、totag不能同时为空
type AppMessage struct {
	ToUser                 string                 `json:"touser,omitempty"`                   // 指定接收消息的成员，成员ID列表（多个接收者用‘|’分隔，最多支持1000个）。	特殊情况：指定为"@all"，则向该企业应用的全部成员发送
	ToParty                string                 `json:"toparty,omitempty"`                  // 指定接收消息的部门，部门ID列表，多个接收者用‘|’分隔，最多支持100个。	当touser为"@all"时忽略本参数
	ToTag                  string                 `json:"totag,omitempty"`                    // 指定接收消息的标签，标签ID列表，多个接收者用‘|’分隔，最多支持100个。	当touser为"@all"时忽略本参数
	AgentID                int64                  `json:"agentid"`                            // 企业应用的id
	MsgType                string                 `json:"msgtype"`                            // 消息类型
	Safe                   int                    `json:"safe,omitempty"`                     // 表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0
	EnableIDTrans          int                    `json:"enable_id_trans,omitempty"`          // 表示是否开启id转译，0表示否，1表示是，默认0。仅第三方应用需要用到，企业自建应用可以忽略。
	EnableDuplicateCheck   int                    `json:"enable_duplicate_check,omitempty"`   // 表示是否开启重复消息检查，0表示否，1表示是，默认0
	DuplicateCheckInterval int                    `json:"duplicate_check_interval,omitempty"` // 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	Text                   *TextMeta              `json:"text,omitempty"`                     // 文本消息
	Image                  *ImageMeta             `json:"image,omitempty"`                    // 图片消息
	Voice                  *VoiceMeta             `json:"voice,omitempty"`                    // 语音消息
	Video                  *VideoMeta             `json:"video,omitempty"`                    // 视频消息
	File                   *FileMeta              `json:"file,omitempty"`                     // 文件消息
	TextCard               *TextCardMeta          `json:"textcard,omitempty"`                 // 文本卡片消息
	News                   *NewsMeta              `json:"news,omitempty"`                     // 图文消息
	MpNews                 *MpNewsMeta            `json:"mpnews,omitempty"`                   // 图文消息
	Markdown               *MarkdownMeta          `json:"markdown,omitempty"`                 // markdown消息
	MiniProgramNotice      *MiniProgramNoticeMeta `json:"miniprogram_notice,omitempty"`       // 小程序通知消息
	TemplateCard           *TemplateCardMeta      `json:"template_card,omitempty"`            // 模板卡片消息
}

// AppMessageResponse 企业微信应用消息响应
type AppMessageResponse struct {
	weixin.ResponseMeta
	MsgID          string `json:"msgid,omitempty"`          // 消息id，用于撤回应用消息
	ResponseCode   string `json:"response_code,omitempty"`  // 仅消息类型为“按钮交互型”，“投票选择型”和“多项选择型”的模板卡片消息返回，应用可使用response_code调用更新模版卡片消息接口，24小时内有效，且只能使用一次
	InvalidUser    string `json:"invaliduser,omitempty"`    // 不合法的userid，不区分大小写，统一转为小写
	InvalidParty   string `json:"invalidparty,omitempty"`   // 不合法的partyid
	InvalidTag     string `json:"invalidtag,omitempty"`     // 不合法的标签id
	UnLicensedUser string `json:"unlicenseduser,omitempty"` // 没有基础接口许可(包含已过期)的userid
}

func (t AppMessageResponse) String() string {
	var sb []string

	if t.MsgID != "" {
		sb = append(sb, fmt.Sprintf("msgid: %q", t.MsgID))
	}
	if t.ResponseCode != "" {
		sb = append(sb, fmt.Sprintf("response_code: %q", t.ResponseCode))
	}
	if t.InvalidUser != "" {
		sb = append(sb, fmt.Sprintf("invaliduser: %q", t.InvalidUser))
	}
	if t.InvalidParty != "" {
		sb = append(sb, fmt.Sprintf("invalidparty: %q", t.InvalidParty))
	}
	if t.InvalidTag != "" {
		sb = append(sb, fmt.Sprintf("invalidtag: %q", t.InvalidTag))
	}
	if t.UnLicensedUser != "" {
		sb = append(sb, fmt.Sprintf("unlicenseduser: %q", t.UnLicensedUser))
	}

	if !t.ResponseMeta.Succeed() {
		sb = append([]string{fmt.Sprintf("%v", t.ResponseMeta)}, sb...)
	}
	return strings.Join(sb, ", ")
}

const appSendURL = work.Host + "/cgi-bin/message/send?access_token="

// SendApp 发送企业微信应用消息
func SendApp(accessToken string, msg *AppMessage) (*AppMessageResponse, error) {
	u := appSendURL + url.QueryEscape(accessToken)
	var resp AppMessageResponse
	_, err := client.PostJSON(u, msg, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; %v", weixin.ErrRequest, resp)
	}
	return &resp, nil
}
