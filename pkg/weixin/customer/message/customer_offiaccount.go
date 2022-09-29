package message

import "fmt"

// 公众号 msgtype 的合法值
const (
	MpMsgTypeText            = "text"            // 文本消息
	MpMsgTypeImage           = "image"           // 图片消息
	MpMsgTypeVoice           = "voice"           // 语音消息
	MpMsgTypeVideo           = "video"           // 视频消息
	MpMsgTypeMusic           = "music"           // 音乐消息
	MpMsgTypeNews            = "news"            // 图文消息（点击跳转到外链）
	MpMsgTypeMpNews          = "mpnews"          // 图文消息（点击跳转到图文消息页面）
	MpMsgTypeMpNewsArticle   = "mpnewsarticle"   // 图文消息（点击跳转到图文消息页面）使用通过 “发布” 系列接口得到的 article_id
	MpMsgTypeMsgMenu         = "msgmenu"         // 菜单消息
	MpMsgTypeWxCard          = "wxcard"          // 卡券
	MpMsgTypeMiniProgramPage = "miniprogrampage" // 小程序卡片（要求小程序与公众号已关联）
)

// ValidateMpMsgType 验证
func ValidateMpMsgType(v string) error {
	switch v {
	case MpMsgTypeText, MpMsgTypeImage, MpMsgTypeVoice, MpMsgTypeVideo,
		MpMsgTypeMusic, MpMsgTypeNews, MpMsgTypeMpNews, MpMsgTypeMpNewsArticle,
		MpMsgTypeMsgMenu, MpMsgTypeWxCard, MpMsgTypeMiniProgramPage:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q %q %q %q %q]", v,
			MpMsgTypeText, MpMsgTypeImage, MpMsgTypeVoice, MpMsgTypeVideo,
			MpMsgTypeMusic, MpMsgTypeNews, MpMsgTypeMpNews, MpMsgTypeMpNewsArticle,
			MpMsgTypeMsgMenu, MpMsgTypeWxCard, MpMsgTypeMiniProgramPage)
	}
	return nil
}
