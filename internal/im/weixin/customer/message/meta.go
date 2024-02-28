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

// ServiceMeta 客服帐号
// 如果需要以某个客服帐号来发消息（在微信6.0.2及以上版本中显示自定义头像），则需在 JSON 数据包的后半部分加入 customservice 参数
type ServiceMeta struct {
	KfAccount string `json:"kf_account"`
}

// TextMeta 文本
type TextMeta struct {
	Content string `json:"content"`
}

// ImageMeta 图片
type ImageMeta struct {
	MediaID string `json:"media_id"`
}

// VoiceMeta 语音
type VoiceMeta struct {
	MediaID string `json:"media_id"`
}

// VideoMeta 视频
type VideoMeta struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// MusicMeta 音乐
type MusicMeta struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HqmusicUrl   string `json:"hqmusicurl"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// NewsMeta 图文（点击跳转到外链）
type NewsMeta struct {
	Articles []Article `json:"articles"`
}

// Article 图文内容（点击跳转到外链）
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

// MpNewsMeta 图文（点击跳转到图文消息页面）
type MpNewsMeta struct {
	MediaID string `json:"media_id"`
}

// MpNewsArticleMeta 图文消息（点击跳转到图文消息页面）使用通过 “发布” 系列接口得到的 article_id
type MpNewsArticleMeta struct {
	ArticleID string `json:"article_id"`
}

// MsgMenuMeta 菜单消息
type MsgMenuMeta struct {
	HeadContent string        `json:"head_content"`
	List        []MsgMenuItem `json:"list"`
	TailContent string        `json:"tail_content"`
}

// MsgMenuItem 菜单内容
type MsgMenuItem struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// WxCardMeta 卡券消息 特别注意客服消息接口投放卡券仅支持非自定义 Code 码和导入 code 模式的卡券的卡券
type WxCardMeta struct {
	CardID string `json:"card_id"`
}

// MiniProgramPageMeta 小程序卡片（要求小程序与公众号已关联）
type MiniProgramPageMeta struct {
	Title        string `json:"title"`
	AppID        string `json:"appid,omitempty"` // 小程序发送不需要填写
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// LinkMeta 图文链接
type LinkMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ThumbUrl    string `json:"thumb_url"`
}
