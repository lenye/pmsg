// Copyright 2022-2025 The pmsg Authors. All rights reserved.
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
	"github.com/lenye/pmsg/im/mattermost/client"
)

// Message Mattermost机器人消息
type Message struct {
	Text        string       `json:"text"`
	Username    string       `json:"username,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Text     string `json:"text,omitempty"`
	Fallback string `json:"fallback,omitempty"`
	Pretext  string `json:"pretext,omitempty"`

	Color string `json:"color,omitempty"`

	AuthorName string `json:"author_name,omitempty"`
	AuthorLink string `json:"author_link,omitempty"`
	AuthorIcon string `json:"author_icon,omitempty"`

	Title     string `json:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty"`

	Fields []Field `json:"fields,omitempty"`

	ImageURL string `json:"image_url,omitempty"`
	ThumbURL string `json:"thumb_url,omitempty"`
}

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

// Send 发送消息
func Send(webhookUrl string, msg *Message) error {
	_, err := client.PostJSON(webhookUrl, msg)
	if err != nil {
		return err
	}
	return nil
}
