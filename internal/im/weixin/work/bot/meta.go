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
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// TextMeta 文本消息
type TextMeta struct {
	Content             string   `json:"content"`                         // 文本内容，最长不超过2048个字节，必须是utf8编码
	MentionedList       []string `json:"mentioned_list,omitempty"`        // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
}

// MarkdownMeta markdown消息
type MarkdownMeta struct {
	Content string `json:"content"` // markdown内容，最长不超过4096个字节，必须是utf8编码
}

// NewsMeta 图文消息
type NewsMeta struct {
	Articles []NewsArticle `json:"articles"` // 一个图文消息支持1到8条图文
}

// NewsArticle 图文
type NewsArticle struct {
	Title       string `json:"title"`                 // 标题，不超过128个字节，超过会自动截断
	Description string `json:"description,omitempty"` // 描述，不超过512个字节，超过会自动截断
	Url         string `json:"url"`                   // 点击后跳转的链接
	PicUrl      string `json:"picurl,omitempty"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150
}

// FileMeta 文件消息
type FileMeta struct {
	MediaID string `json:"media_id"` //	文件id，通过文件上传接口获取
}

// ImageMeta 图片消息
type ImageMeta struct {
	Base64 string `json:"base64"` // 图片内容的base64编码
	MD5    string `json:"md5"`    // 图片内容（base64编码前）的md5值
}

// ImageFile2Meta 图片文件转换为图片消息
func ImageFile2Meta(filename string) (*ImageMeta, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("file: %q, open file failed: %w", filename, err)
	}
	defer f.Close()

	base64Buf := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, base64Buf)
	if _, err := io.Copy(encoder, f); err != nil {
		return nil, fmt.Errorf("file: %q, base64 io.Copy failed: %w", filename, err)
	}
	if err := encoder.Close(); err != nil {
		return nil, fmt.Errorf("base64 encoder close failed: %w", err)
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("file: %q, file seek start failed: %w", filename, err)
	}

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return nil, fmt.Errorf("file: %q, md5 io.Copy failed: %w", filename, err)
	}

	return &ImageMeta{
		Base64: base64Buf.String(),
		MD5:    hex.EncodeToString(hash.Sum(nil)),
	}, nil
}

// TemplateCardText 文本通知模版卡片
type TemplateCardText struct {
	CardType              string         `json:"card_type"`                  // 模版卡片的模版类型，文本通知模版卡片的类型为text_notice
	Source                *CardSource    `json:"source,omitempty"`           // 卡片来源样式信息，不需要来源样式可不填写
	MainTitle             CardTitle      `json:"main_title"`                 // 模版卡片的主要内容，包括一级标题和标题辅助信息
	EmphasisContent       *CardTitle     `json:"emphasis_content,omitempty"` // 关键数据样式
	QuoteArea             *CardQuoteArea `json:"quote_area,omitempty"`       // 引用文献样式，建议不与关键数据共用
	SubTitleText          string         `json:"sub_title_text,omitempty"`   // 二级普通文本，建议不超过112个字。模版卡片主要内容的一级标题main_title.title和二级普通文本sub_title_text必须有一项填写
	HorizontalContentList []CardContent  `json:"horizontal_content_list"`    // 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
	JumpList              []JumpContent  `json:"jump_list"`                  // 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
	CardAction            CardAction     `json:"card_action"`                // 整体卡片的点击跳转事件，text_notice模版卡片中该字段为必填项
}

// TemplateCardNews 图文展示模版卡片
type TemplateCardNews struct {
	CardType              string             `json:"card_type"`               // 模版卡片的模版类型，文本通知模版卡片的类型为text_notice
	Source                *CardSource        `json:"source,omitempty"`        // 卡片来源样式信息，不需要来源样式可不填写
	MainTitle             CardTitle          `json:"main_title"`              // 模版卡片的主要内容，包括一级标题和标题辅助信息
	CardImage             CardImage          `json:"card_image"`              // 图片样式
	ImageTextArea         *CardImageTextArea `json:"image_text_area"`         // 左图右文样式
	QuoteArea             *CardQuoteArea     `json:"quote_area,omitempty"`    // 引用文献样式，建议不与关键数据共用
	VerticalContentList   []CardTitle        `json:"vertical_content_list"`   // 卡片二级垂直内容，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过4
	HorizontalContentList []CardContent      `json:"horizontal_content_list"` // 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
	JumpList              []JumpContent      `json:"jump_list"`               // 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
	CardAction            CardAction         `json:"card_action"`             // 整体卡片的点击跳转事件，text_notice模版卡片中该字段为必填项
}

// CardSource 卡片来源样式信息，不需要来源样式可不填写
type CardSource struct {
	IconURL   string `json:"icon_url,omitempty"`   // 来源图片的url
	Desc      string `json:"desc,omitempty"`       // 来源图片的描述，建议不超过13个字
	DescColor int    `json:"desc_color,omitempty"` // 来源文字的颜色，目前支持：0(默认) 灰色，1 黑色，2 红色，3 绿色
}

// CardTitle 标题和标题辅助信息
type CardTitle struct {
	Title string `json:"title"`          // 标题，建议不超过26个字。模版卡片主要内容的一级标题main_title.title和二级普通文本sub_title_text必须有一项填写
	Desc  string `json:"desc,omitempty"` // 标题辅助信息，建议不超过30个字
}

// CardImage 图片样式
type CardImage struct {
	Url         string  `json:"url"`                    // 图片的url
	AspectRatio float64 `json:"aspect_ratio,omitempty"` // 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
}

// CardImageTextArea 左图右文样式
type CardImageTextArea struct {
	Type     int    `json:"type,omitempty"`     // 左图右文样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	Url      string `json:"url,omitempty"`      // 点击跳转的url，image_text_area.type是1时必填
	Title    string `json:"title,omitempty"`    // 左图右文样式的标题
	Desc     string `json:"desc,omitempty"`     // 左图右文样式的描述
	ImageUrl string `json:"image_url"`          // 左图右文样式的图片url
	AppID    string `json:"appid,omitempty"`    // 点击跳转的小程序的appid，必须是与当前应用关联的小程序，image_text_area.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 点击跳转的小程序的pagepath，image_text_area.type是2时选填
}

// CardQuoteArea 引用文献样式，建议不与关键数据共用
type CardQuoteArea struct {
	Type      int    `json:"type,omitempty"`       // 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	URL       string `json:"url,omitempty"`        // 点击跳转的url，quote_area.type是1时必填
	Title     string `json:"title,omitempty"`      // 引用文献样式的标题
	QuoteText string `json:"quote_text,omitempty"` // 引用文献样式的引用文案
	AppID     string `json:"appid,omitempty"`      // 点击跳转的小程序的appid，quote_area.type是2时必填
	PagePath  string `json:"pagepath,omitempty"`   // 点击跳转的小程序的pagepath，quote_area.type是2时选填
}

// CardContent 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
type CardContent struct {
	KeyName string `json:"keyname"`            // 链接类型，0或不填代表是普通文本，1 代表跳转url，2 代表下载附件，3 代表@员工
	Value   string `json:"value,omitempty"`    // 二级标题，建议不超过5个字
	Type    int    `json:"type,omitempty"`     // 二级文本，如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过26个字
	URL     string `json:"url,omitempty"`      // 链接跳转的url，horizontal_content_list.type是1时必填
	MediaID string `json:"media_id,omitempty"` // 附件的media_id，horizontal_content_list.type是2时必填
	UserID  string `json:"userid,omitempty"`   // 被@的成员的userid，horizontal_content_list.type是3时必填
}

// JumpContent 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
type JumpContent struct {
	Type     int    `json:"type,omitempty"`     // 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
	URL      string `json:"url,omitempty"`      // 跳转链接的url，jump_list.type是1时必填
	Title    string `json:"title"`              // 跳转链接样式的文案内容，建议不超过13个字
	AppID    string `json:"appid,omitempty"`    // 跳转链接的小程序的appid，jump_list.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 跳转链接的小程序的pagepath，jump_list.type是2时选填
}

// CardAction 整体卡片的点击跳转事件，text_notice模版卡片中该字段为必填项
type CardAction struct {
	Type     int    `json:"type"`               // 卡片跳转类型，1 代表跳转url，2 代表打开小程序。text_notice模版卡片中该字段取值范围为[1,2]
	URL      string `json:"url,omitempty"`      // 跳转事件的url，card_action.type是1时必填
	AppID    string `json:"appid,omitempty"`    // 跳转事件的小程序的appid，card_action.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 跳转事件的小程序的pagepath，card_action.type是2时选填
}
