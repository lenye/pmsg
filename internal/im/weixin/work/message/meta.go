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

// TextMeta 文本消息
type TextMeta struct {
	Content string `json:"content"` // 消息内容，最长不超过2048个字节，超过将截断（支持id转译）
}

// ImageMeta 图片消息
type ImageMeta struct {
	MediaID string `json:"media_id"`
}

// VoiceMeta 语音消息
type VoiceMeta struct {
	MediaID string `json:"media_id"`
}

// VideoMeta 视频消息
type VideoMeta struct {
	MediaID     string `json:"media_id"`              // 视频媒体文件id
	Title       string `json:"title,omitempty"`       // 视频消息的标题，不超过128个字节，超过会自动截断
	Description string `json:"description,omitempty"` // 视频消息的描述，不超过512个字节，超过会自动截断
}

// FileMeta 文件消息
type FileMeta struct {
	MediaID string `json:"media_id"`
}

// TextCardMeta 文本卡片消息
type TextCardMeta struct {
	Title       string `json:"title"`            // 标题，不超过128个字节，超过会自动截断（支持id转译）
	Description string `json:"description"`      // 描述，不超过512个字节，超过会自动截断（支持id转译）
	Url         string `json:"url"`              // 点击后跳转的链接。最长2048字节，请确保包含了协议头(http/https)
	BtnTxt      string `json:"btntxt,omitempty"` // 按钮文字。 默认为“详情”， 不超过4个文字，超过自动截断。
}

// NewsMeta 图文消息
type NewsMeta struct {
	Articles []NewsArticle `json:"articles"` // 一个图文消息支持1到8条图文
}

// NewsArticle 图文
type NewsArticle struct {
	Title       string `json:"title"`                 // 标题，不超过128个字节，超过会自动截断（支持id转译）
	Description string `json:"description,omitempty"` // 描述，不超过512个字节，超过会自动截断（支持id转译）
	Url         string `json:"url,omitempty"`         // 点击后跳转的链接
	PicUrl      string `json:"picurl,omitempty"`      // 图文消息的图片链接
	AppID       string `json:"appid"`                 // 小程序appid，必须是与当前应用关联的小程序，appid和pagepath必须同时填写，填写后会忽略url字段
	PagePath    string `json:"pagepath"`              // 点击消息卡片后的小程序页面，最长128字节，仅限本小程序内的页面
}

// MpNewsMeta 图文消息
type MpNewsMeta struct {
	Articles []MpNewsArticle `json:"articles"` // 一个图文消息支持1到8条图文
}

// MpNewsArticle 图文
type MpNewsArticle struct {
	Title            string `json:"title"`                        // 标题，不超过128个字节，超过会自动截断（支持id转译）
	ThumbMediaID     string `json:"thumb_media_id"`               // 图文消息缩略图的media_id
	Author           string `json:"author,omitempty"`             // 图文消息的作者，不超过64个字节
	ContentSourceUrl string `json:"content_source_url,omitempty"` // 图文消息点击“阅读原文”之后的页面链接
	Content          string `json:"content"`                      // 图文消息的内容，支持html标签，不超过666 K个字节（支持id转译）
	Digest           string `json:"digest,omitempty"`             // 图文消息的描述，不超过512个字节，超过会自动截断（支持id转译）
}

// MarkdownMeta markdown消息
type MarkdownMeta struct {
	Content string `json:"content"` // markdown内容，最长不超过2048个字节，必须是utf8编码
}

// MiniProgramNoticeMeta 小程序通知消息  小程序通知消息只允许绑定了小程序的应用发送 不支持@all全员发送
type MiniProgramNoticeMeta struct {
	AppID             string              `json:"appid"`                         // 小程序appid，必须是与当前应用关联的小程序
	Page              string              `json:"page,omitempty"`                // 点击消息卡片后的小程序页面，最长1024个字节，仅限本小程序内的页面。该字段不填则消息点击后不跳转。
	Title             string              `json:"title"`                         // 消息标题，长度限制4-12个汉字（支持id转译）
	Description       string              `json:"description,omitempty"`         // 消息描述，长度限制4-12个汉字（支持id转译）
	EmphasisFirstItem bool                `json:"emphasis_first_item,omitempty"` // 是否放大第一个content_item
	ContentItem       []NoticeContentItem `json:"content_item,omitempty"`        // 消息内容键值对，最多允许10个item
}

// NoticeContentItem 消息内容键值对
type NoticeContentItem struct {
	Key   string `json:"key"`   // 长度10个汉字以内
	Value string `json:"value"` // 长度30个汉字以内（支持id转译）
}

// LinkMeta 图文链接
type LinkMeta struct {
	Title        string `json:"title"`
	Desc         string `json:"desc,omitempty"`
	Url          string `json:"url"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// MiniProgramMeta 小程序消息
type MiniProgramMeta struct {
	Title        string `json:"title,omitempty"`
	AppID        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// MsgMenuMeta 菜单消息
type MsgMenuMeta struct {
	HeadContent string        `json:"head_content,omitempty"`
	List        []MsgMenuItem `json:"list"`
	TailContent string        `json:"tail_content,omitempty"`
}

// MsgMenuItem 菜单内容
type MsgMenuItem struct {
	Type        string              `json:"type"`
	Click       *MsgMenuClick       `json:"click,omitempty"`
	View        *MsgMenuView        `json:"view,omitempty"`
	MiniProgram *MsgMenuMiniProgram `json:"miniprogram,omitempty"`
	Text        *MsgMenuText        `json:"text,omitempty"`
}

type MsgMenuClick struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content"`
}

type MsgMenuView struct {
	Url     string `json:"url"`
	Content string `json:"content"`
}

type MsgMenuMiniProgram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Content  string `json:"content"`
}

type MsgMenuText struct {
	Content string `json:"content"`
}

// LocationMeta 地理位置消息
type LocationMeta struct {
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}

// TemplateCardMeta 模板卡片消息
type TemplateCardMeta struct {
	CardType              string                         `json:"card_type"` // 模板卡片类型
	Source                *TplCardSource                 `json:"source,omitempty"`
	ActionMenu            *TplCardActionMenu             `json:"action_menu,omitempty"`
	TaskID                string                         `json:"task_id,omitempty"`
	MainTitle             *TplCardMainTitle              `json:"main_title,omitempty"`
	QuoteArea             *TplCardQuoteArea              `json:"quote_area,omitempty"`
	EmphasisContent       *TplCardEmphasisContent        `json:"emphasis_content,omitempty"`
	SubTitleText          string                         `json:"sub_title_text,omitempty"`
	HorizontalContentList []TplCardHorizontalContentItem `json:"horizontal_content_list,omitempty"`
	JumpList              []TplCardJumpItem              `json:"jump_list,omitempty"`
	CardAction            *TplCardAction                 `json:"card_action,omitempty"`
	ImageTextArea         *TplImageTextArea              `json:"image_text_area,omitempty"`
	CardImage             *TplCardImage                  `json:"card_image,omitempty"`
	VerticalContentList   []TplVerticalContentItem       `json:"vertical_content_list,omitempty"`
	ButtonSelection       *TplButtonSelection            `json:"button_selection,omitempty"`
	ButtonList            []TplButtonItem                `json:"button_list,omitempty"`
	Checkbox              *TplCheckbox                   `json:"checkbox,omitempty"`
	SubmitButton          *TplSubmitButton               `json:"submit_button,omitempty"`
	SelectList            []TplSelectItem                `json:"select_list,omitempty"`
}

type TplCardTextNotice struct {
	SelectList   []TplSelectItem `json:"select_list,omitempty"`
	SubmitButton struct {
		Text string `json:"text"`
		Key  string `json:"key"`
	} `json:"submit_button"`
}

// TplCardSource 卡片来源样式信息，不需要来源样式可不填写
type TplCardSource struct {
	IconUrl   string `json:"icon_url,omitempty"`   // 来源图片的url，来源图片的尺寸建议为72*72
	Desc      string `json:"desc,omitempty"`       // 来源图片的描述，建议不超过20个字，（支持id转译）
	DescColor int    `json:"desc_color,omitempty"` // 来源文字的颜色，目前支持：0(默认) 灰色，1 黑色，2 红色，3 绿色
}

// TplCardActionMenu 卡片右上角更多操作按钮
type TplCardActionMenu struct {
	Desc       string              `json:"desc,omitempty"` // 更多操作界面的描述
	ActionList []TplCardActionItem `json:"action_list"`    // 操作列表，列表长度取值范围为 [1, 3]
}

// TplCardActionItem 操作
type TplCardActionItem struct {
	Text string `json:"text"` // 操作的描述文案
	Key  string `json:"key"`  // 操作key值，用户点击后，会产生回调事件将本参数作为EventKey返回，回调事件会带上该key值，最长支持1024字节，不可重复
}

// TplCardMainTitle 一级标题
type TplCardMainTitle struct {
	Title string `json:"title,omitempty"` // 一级标题，建议不超过36个字，文本通知型卡片本字段非必填，但不可本字段和sub_title_text都不填，（支持id转译）
	Desc  string `json:"desc,omitempty"`  // 标题辅助信息，建议不超过44个字，（支持id转译）
}

// TplCardQuoteArea 引用文献样式
type TplCardQuoteArea struct {
	Type      int    `json:"type"`               // 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	Url       string `json:"url"`                // 点击跳转的url，quote_area.type是1时必填
	AppID     string `json:"appid,omitempty"`    // 点击跳转的小程序的appid，必须是与当前应用关联的小程序，quote_area.type是2时必填
	PagePath  string `json:"pagepath,omitempty"` // 点击跳转的小程序的pagepath，quote_area.type是2时选填
	Title     string `json:"title"`              // 引用文献样式的标题
	QuoteText string `json:"quote_text"`         // 引用文献样式的引用文案
}

// TplCardEmphasisContent 关键数据样式
type TplCardEmphasisContent struct {
	Title string `json:"title"` // 关键数据样式的数据内容，建议不超过14个字
	Desc  string `json:"desc"`  // 关键数据样式的数据描述内容，建议不超过22个字
}

// TplCardHorizontalContentItem 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
type TplCardHorizontalContentItem struct {
	KeyName string `json:"keyname"`            // 二级标题，建议不超过5个字
	Value   string `json:"value,omitempty"`    // 二级文本，如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过30个字，（支持id转译）
	Type    int    `json:"type,omitempty"`     //	链接类型，0或不填代表不是链接，1 代表跳转url，2 代表下载附件，3 代表点击跳转成员详情
	Url     string `json:"url,omitempty"`      // 链接跳转的url，horizontal_content_list.type是1时必填
	MediaID string `json:"media_id,omitempty"` // 附件的media_id，horizontal_content_list.type是2时必填
	UserID  string `json:"userid,omitempty"`   // 成员详情的userid，horizontal_content_list.type是3时必填
}

// TplCardJumpItem 	跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
type TplCardJumpItem struct {
	Type     int    `json:"type,omitempty"`     // 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
	Title    string `json:"title"`              // 跳转链接样式的文案内容，建议不超过18个字
	Url      string `json:"url,omitempty"`      // 跳转链接的url，jump_list.type是1时必填
	AppID    string `json:"appid,omitempty"`    // 跳转链接的小程序的appid，必须是与当前应用关联的小程序，jump_list.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 跳转链接的小程序的pagepath，jump_list.type是2时选填
}

// TplCardAction 整体卡片的点击跳转事件，text_notice必填本字段
type TplCardAction struct {
	Type     int    `json:"type"`               // 跳转事件类型，1 代表跳转url，2 代表打开小程序。text_notice卡片模版中该字段取值范围为[1,2]
	Url      string `json:"url,omitempty"`      // 跳转事件的url，card_action.type是1时必填
	AppID    string `json:"appid,omitempty"`    // 跳转事件的小程序的appid，必须是与当前应用关联的小程序，card_action.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 跳转事件的小程序的pagepath，card_action.type是2时选填
}

// TplImageTextArea 左图右文样式
type TplImageTextArea struct {
	Type     int    `json:"type,omitempty"`     // 左图右文样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	Url      string `json:"url,omitempty"`      // 点击跳转的url，image_text_area.type是1时必填
	Title    string `json:"title,omitempty"`    // 左图右文样式的标题
	Desc     string `json:"desc,omitempty"`     // 左图右文样式的描述
	AppID    string `json:"appid,omitempty"`    // 点击跳转的小程序的appid，必须是与当前应用关联的小程序，image_text_area.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 点击跳转的小程序的pagepath，image_text_area.type是2时选填
	ImageUrl string `json:"image_url"`          // 左图右文样式的图片url
}

// TplCardImage 图片样式，news_notice类型的卡片，card_image和image_text_area两者必填一个字段，不可都不填
type TplCardImage struct {
	Url         string  `json:"url"`          // 图片的url
	AspectRatio float64 `json:"aspect_ratio"` // 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
}

// TplVerticalContentItem 卡片二级垂直内容，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过4
type TplVerticalContentItem struct {
	Title string `json:"title"`          // 卡片二级标题，建议不超过38个字
	Desc  string `json:"desc,omitempty"` // 二级普通文本，建议不超过160个字
}

// TplButtonSelection 下拉式的选择器
type TplButtonSelection struct {
	QuestionKey string                `json:"question_key"`          // 下拉式的选择器的key，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
	Title       string                `json:"title,omitempty"`       // 下拉式的选择器左边的标题
	OptionList  []TplButtonOptionItem `json:"option_list"`           // 选项列表，下拉选项不超过 10 个，最少1个
	SelectedID  string                `json:"selected_id,omitempty"` // 默认选定的id，不填或错填默认第一个
}

// TplButtonOptionItem 下拉式的选择器选项
type TplButtonOptionItem struct {
	ID   string `json:"id"`   // 下拉式的选择器选项的id，用户提交后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
	Text string `json:"text"` // 下拉式的选择器选项的文案，建议不超过16个字
}

// TplButtonItem 按钮 按钮列表，列表长度不超过6
type TplButtonItem struct {
	Text  string `json:"text"`            // 按钮文案，建议不超过10个字
	Style int    `json:"style,omitempty"` // 按钮样式，目前可填1~4，不填或错填默认1
	Key   string `json:"key,omitempty"`   // 按钮key值，用户点击后，会产生回调事件将本参数作为EventKey返回，回调事件会带上该key值，最长支持1024字节，不可重复，button_list.type是0时必填
	Type  string `json:"type,omitempty"`  // 按钮点击事件类型，0 或不填代表回调点击事件，1 代表跳转url
	Url   string `json:"url,omitempty"`   // 跳转事件的url，button_list.type是1时必填
}

// TplCheckbox 选择题样式
type TplCheckbox struct {
	QuestionKey string                  `json:"question_key"`   // 选择题key值，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
	OptionList  []TplCheckboxOptionItem `json:"option_list"`    // 选项list，选项个数不超过 20 个，最少1个
	Mode        int                     `json:"mode,omitempty"` // 选择题模式，单选：0，多选：1，不填默认0
}

// TplCheckboxOptionItem 选项
type TplCheckboxOptionItem struct {
	ID        string `json:"id"`         // 选项id，用户提交选项后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
	Text      string `json:"text"`       // 选项文案描述，建议不超过17个字
	IsChecked bool   `json:"is_checked"` // 该选项是否要默认选中
}

// TplSubmitButton 提交按钮样式
type TplSubmitButton struct {
	Text string `json:"text"` // 按钮文案，建议不超过10个字，不填默认为提交
	Key  string `json:"key"`  // 提交按钮的key，会产生回调事件将本参数作为EventKey返回，最长支持1024字节
}

// TplSelectItem 下拉式的选择器列表，multiple_interaction类型的卡片该字段不可为空，一个消息最多支持 3 个选择器
type TplSelectItem struct {
	QuestionKey string                `json:"question_key"`          // 下拉式的选择器题目的key，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节，不可重复
	Title       string                `json:"title,omitempty"`       // 下拉式的选择器上面的title
	SelectedID  string                `json:"selected_id,omitempty"` // 默认选定的id，不填或错填默认第一个
	OptionList  []TplSelectOptionItem `json:"option_list"`           // 选项列表，下拉选项不超过 10 个，最少1个
}

// TplSelectOptionItem 下拉式的选择器选项
type TplSelectOptionItem struct {
	ID   string `json:"id"`   // 下拉式的选择器选项的id，用户提交选项后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节，不可重复
	Text string `json:"text"` // 下拉式的选择器选项的文案，建议不超过16个字
}
