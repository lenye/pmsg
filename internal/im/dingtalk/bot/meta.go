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

// TextMeta 文本
type TextMeta struct {
	Content string `json:"content"` // 消息内容
}

// AtMeta @用户
type AtMeta struct {
	AtMobiles []string `json:"atMobiles,omitempty"` // 被@人的手机号。
	AtUserIds []string `json:"atUserIds,omitempty"` // 被@人的用户userid
	IsAtAll   bool     `json:"isAtAll,omitempty"`   // 是否@所有人
}

// LinkMeta 链接
type LinkMeta struct {
	Title      string `json:"title"`            // 消息标题
	Text       string `json:"text"`             // 消息内容。如果太长只会部分展示。
	MessageUrl string `json:"messageUrl"`       // 点击消息跳转的URL
	PicUrl     string `json:"picUrl,omitempty"` // 图片URL
}

// MarkdownMeta markdown
type MarkdownMeta struct {
	Title string `json:"title"` // 首屏会话透出的展示内容
	Text  string `json:"text"`  // markdown格式的消息
}

// SingleActionCardMeta 整体跳转
type SingleActionCardMeta struct {
	Title          string `json:"title"`                    // 首屏会话透出的展示内容
	Text           string `json:"text"`                     // markdown格式的消息
	SingleTitle    string `json:"singleTitle"`              // 单个按钮的标题
	SingleURL      string `json:"singleURL"`                // 点击消息跳转的
	BtnOrientation string `json:"btnOrientation,omitempty"` // 0：按钮竖直排列	1：按钮横向排列
}

// ActionCardMeta 独立跳转
type ActionCardMeta struct {
	Title          string              `json:"title"`                    // 首屏会话透出的展示内容
	Text           string              `json:"text"`                     // markdown格式的消息
	Btns           []ActionCardBtnMeta `json:"btns"`                     // 按钮
	BtnOrientation string              `json:"btnOrientation,omitempty"` // 0：按钮竖直排列	1：按钮横向排列
}

// ActionCardBtnMeta 按钮
type ActionCardBtnMeta struct {
	Title     string `json:"title"`     // 按钮标题
	ActionURL string `json:"actionURL"` // 点击按钮触发的URL
}

// FeedCardMeta free card
type FeedCardMeta struct {
	Links []FeedCardLinkMeta `json:"links"`
}

// FeedCardLinkMeta 链接信息
type FeedCardLinkMeta struct {
	Title      string `json:"title"`      // 单条信息文本
	MessageURL string `json:"messageURL"` // 点击单条信息到跳转链接
	PicURL     string `json:"picURL"`     // 单条信息后面图片的URL
}
