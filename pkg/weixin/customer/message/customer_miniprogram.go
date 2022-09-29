package message

import "fmt"

// 小程序 msgtype 的合法值
const (
	MiniProgramMsgTypeText            = "text"            // 文本消息
	MiniProgramMsgTypeImage           = "image"           // 图片消息
	MiniProgramMsgTypeLink            = "link"            // 图文链接
	MiniProgramMsgTypeMiniProgramPage = "miniprogrampage" // 小程序卡片
)

// ValidateMiniProgramMsgType 验证
func ValidateMiniProgramMsgType(v string) error {
	switch v {
	case MiniProgramMsgTypeText, MiniProgramMsgTypeImage, MiniProgramMsgTypeLink, MiniProgramMsgTypeMiniProgramPage:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q]", v,
			MiniProgramMsgTypeText, MiniProgramMsgTypeImage, MiniProgramMsgTypeLink, MiniProgramMsgTypeMiniProgramPage)
	}
	return nil
}
