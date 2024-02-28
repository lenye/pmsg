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

package asset

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/lenye/pmsg/internal/im/weixin"
	"github.com/lenye/pmsg/internal/im/weixin/client"
)

const (
	FieldName = "media"

	TypeImage = "image" // 图片（image）: 10M，支持PNG\JPEG\JPG\GIF格式
	TypeVoice = "voice" // 语音（voice）：2M，播放长度不超过60s，支持AMR\MP3格式
	TypeVideo = "video" // 视频（video）：10MB，支持MP4格式
	TypeThumb = "thumb" // 缩略图（thumb）：64KB，支持 JPG 格式
)

// ValidateMediaType 验证
func ValidateMediaType(v string) error {
	switch v {
	case TypeImage, TypeVoice, TypeVideo, TypeThumb:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q]", v,
			TypeImage, TypeVoice, TypeVideo, TypeThumb)
	}
	return nil
}

type MediaResponse struct {
	weixin.ResponseMeta
	MediaMeta
}

type MediaMeta struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

func (t MediaMeta) String() string {
	var sb []string

	if t.Type != "" {
		sb = append(sb, fmt.Sprintf("type: %q", t.Type))
	}
	if t.MediaID != "" {
		sb = append(sb, fmt.Sprintf("media_id: %q", t.MediaID))
	}
	locCreatedAt := time.Unix(t.CreatedAt, 0).Local()
	sb = append(sb, fmt.Sprintf("created_at: %v (%v)", t.CreatedAt, locCreatedAt.Format(time.RFC3339)))
	return strings.Join(sb, ", ")
}

const reqURL = weixin.Host + "/cgi-bin/media/upload?access_token="

// MediaUpload 微信公众号/小程序 新增临时素材 媒体文件在微信后台保存时间为3天，即3天后media_id失效。
func MediaUpload(accessToken, mediaType, filename string) (*MediaMeta, error) {
	u := reqURL + url.QueryEscape(accessToken) + "&type=" + url.QueryEscape(mediaType)
	var resp MediaResponse
	_, err := client.PostFileJSON(u, FieldName, filename, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; %v", weixin.ErrRequest, resp)
	}
	return &resp.MediaMeta, nil
}
