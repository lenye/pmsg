package media

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

const (
	FieldName = "media"

	TypeImage = "image" // 图片（image）: 10M，支持PNG\JPEG\JPG\GIF格式
	TypeVoice = "voice" // 语音（voice）：2M，播放长度不超过60s，支持AMR\MP3格式
	TypeVideo = "video" // 视频（video）：10MB，支持MP4格式
	TypeThumb = "thumb" // 缩略图（thumb）：64KB，支持 JPG 格式
)

type Response struct {
	weixin.ResponseMeta
	Meta
}

type Meta struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// Upload 新增临时素材
func Upload(accessToken, mediaType, filename string) (*Meta, error) {
	url := "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=" + accessToken + "&type=" + mediaType
	var resp Response
	_, err := client.PostFileJSON(url, FieldName, filename, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; %v", weixin.ErrWeiXinRequest, resp)
	}
	return &resp.Meta, nil
}
