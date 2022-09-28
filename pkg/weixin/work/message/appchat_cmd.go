package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/work/token"
)

type CmdWorkSendAppChatParams struct {
	UserAgent   string
	AccessToken string
	CorpID      string
	CorpSecret  string
	ChatID      string
	MsgType     string
	Safe        int
	Data        string
}

func (t *CmdWorkSendAppChatParams) Validate() error {
	if t.AccessToken == "" && t.CorpID == "" {
		return flags.ErrWeixinWorkAccessToken
	}

	if t.Safe != 0 && t.Safe != 1 {
		return fmt.Errorf("invalid %v", flags.Safe)
	}

	if err := ValidateAppChatMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdWorkSendAppChat 发送企业微信群聊推送消息
func CmdWorkSendAppChat(arg *CmdWorkSendAppChatParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := AppChatMessage{
		ChatID:  arg.ChatID,
		MsgType: arg.MsgType,
		Safe:    arg.Safe,
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case AppChatMsgTypeText:
		var msgMeta TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case AppChatMsgTypeImage:
		var msgMeta ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case AppChatMsgTypeVoice:
		var msgMeta VoiceMeta
		msgMeta.MediaID = buf.String()
		msg.Voice = &msgMeta
	case AppChatMsgTypeVideo:
		var msgMeta VideoMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MediaID == "" {
			return errors.New("media_id is empty")
		}
		msg.Video = &msgMeta
	case AppChatMsgTypeFile:
		var msgMeta FileMeta
		msgMeta.MediaID = buf.String()
		msg.File = &msgMeta
	case AppChatMsgTypeTextCard:
		var msgMeta TextCardMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		msg.TextCard = &msgMeta
	case AppChatMsgTypeNews:
		var msgMeta NewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.News = &msgMeta
	case AppChatMsgTypeMpNews:
		var msgMeta MpNewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.MpNews = &msgMeta
	case AppChatMsgTypeMarkdown:
		var msgMeta MarkdownMeta
		msgMeta.Content = buf.String()
		msg.Markdown = &msgMeta
	}

	client.UserAgent = arg.UserAgent

	if arg.AccessToken == "" {
		accessTokenResp, err := token.GetAccessToken(arg.CorpID, arg.CorpSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if err := SendAppChat(arg.AccessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
