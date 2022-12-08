// Copyright 2022 The pmsg Authors. All rights reserved.
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

package webhook

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
