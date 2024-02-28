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

type ContentMeta struct {
	Text        string    `json:"text,omitempty"`          // 文本
	ImageKey    string    `json:"image_key,omitempty"`     // 图片
	ShareChatID string    `json:"share_chat_id,omitempty"` // 分享群名片
	Post        *PostMeta `json:"post,omitempty"`          // 富文本
}

// PostMeta 富文本
type PostMeta struct {
	ZhCn PostZhCn `json:"zh_cn"`
}

type PostZhCn struct {
	Title   string              `json:"title"`
	Content [][]PostZhCnContent `json:"content"`
}

type PostZhCnContent struct {
	Tag       string `json:"tag"`
	Text      string `json:"text,omitempty"`
	Href      string `json:"href,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	ImageKey  string `json:"image_key,omitempty"`
	FileKey   string `json:"file_key,omitempty"`
	EmojiType string `json:"emoji_type,omitempty"`
}

// CardMeta 消息卡片
type CardMeta struct {
	Elements []CardElement `json:"elements"`
	Header   CardHeader    `json:"header"`
}

type CardHeader struct {
	Title CardHeaderTitle `json:"title"`
}

type CardHeaderTitle struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type CardElement struct {
	Tag     string              `json:"tag"`
	Text    CardElementText     `json:"text,omitempty"`
	Actions []CardElementAction `json:"actions,omitempty"`
}

type CardElementText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type CardElementAction struct {
	Tag   string                `json:"tag"`
	Text  CardElementActionText `json:"text"`
	Url   string                `json:"url"`
	Type  string                `json:"type"`
	Value struct {
	} `json:"value"`
}

type CardElementActionText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}
