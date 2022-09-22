package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/customer/message"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

// weiXinOfficialAccountCustomerCmd 微信公众号客服
var weiXinOfficialAccountCustomerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "publish weixin offiaccount customer message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinOfficialAccountSendCustomer(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinOfficialAccountCmd.AddCommand(weiXinOfficialAccountCustomerCmd)

	weiXinSetAccessTokenFlags(weiXinOfficialAccountCustomerCmd)

	weiXinOfficialAccountCustomerCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountCustomerCmd.MarkFlagRequired(nameOpenID)

	weiXinOfficialAccountCustomerCmd.Flags().StringVar(&msgType, nameMsgType, "", "message type (required)")
	weiXinOfficialAccountCustomerCmd.MarkFlagRequired(nameMsgType)

	weiXinOfficialAccountCustomerCmd.Flags().StringVarP(&kfAccount, nameKfAccount, "k", "", "customer account")
}

// WeiXinOfficialAccountSendCustomer 发送微信公众号客服消息
func WeiXinOfficialAccountSendCustomer(args []string) error {

	if accessToken == "" {
		if appID == "" {
			return ErrMultiRequiredOne
		}
	}

	if err := message.ValidateMpMsgType(msgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", nameMsgType, err)
	}

	msg := message.CustomerMessage{
		ToUser:  openID,
		MsgType: msgType,
	}

	if kfAccount != "" {
		msg.CustomService = &message.ServiceMeta{KfAccount: kfAccount}
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(args[0])
	switch msgType {
	case message.MpMsgTypeText:
		var msgMeta message.TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case message.MpMsgTypeImage:
		var msgMeta message.ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case message.MpMsgTypeVoice:
		var msgMeta message.VoiceMeta
		msgMeta.MediaID = buf.String()
		msg.Voice = &msgMeta
	case message.MpMsgTypeVideo:
		var msgMeta message.VideoMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MediaID == "" {
			return errors.New("media_id is empty")
		}
		if msgMeta.ThumbMediaID == "" {
			return errors.New("thumb_media_id is empty")
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.Description == "" {
			return errors.New("description is empty")
		}
		msg.Video = &msgMeta
	case message.MpMsgTypeMusic:
		var msgMeta message.MusicMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MusicUrl == "" {
			return errors.New("musicurl is empty")
		}
		if msgMeta.HqmusicUrl == "" {
			return errors.New("hqmusicurl is empty")
		}
		if msgMeta.ThumbMediaID == "" {
			return errors.New("thumb_media_id is empty")
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.Description == "" {
			return errors.New("description is empty")
		}
		msg.Music = &msgMeta
	case message.MpMsgTypeNews:
		var msgMeta message.NewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if len(msgMeta.Articles) != 1 {
			return errors.New("articles length != 1")
		}
		article := msgMeta.Articles[0]
		if article.Url == "" {
			return errors.New("url is empty")
		}
		if article.PicUrl == "" {
			return errors.New("picurl is empty")
		}
		if article.Title == "" {
			return errors.New("title is empty")
		}
		if article.Description == "" {
			return errors.New("description is empty")
		}
		msg.News = &msgMeta
	case message.MpMsgTypeMpNews:
		var msgMeta message.MpNewsMeta
		msgMeta.MediaID = buf.String()
		msg.MpNews = &msgMeta
	case message.MpMsgTypeMpNewsArticle:
		var msgMeta message.MpNewsArticleMeta
		msgMeta.ArticleID = buf.String()
		msg.MpNewsArticle = &msgMeta
	case message.MpMsgTypeMsgMenu:
		var msgMeta message.MsgMenuMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.HeadContent == "" {
			return errors.New("head_content is empty")
		}
		if msgMeta.TailContent == "" {
			return errors.New("tail_content is empty")
		}
		lenList := len(msgMeta.List)
		if lenList == 0 {
			return errors.New("list is empty")
		}
		for i := 0; i < lenList; i++ {
			if msgMeta.List[i].ID == "" {
				return fmt.Errorf("list[%v].id is empty", i)
			}
			if msgMeta.List[i].Content == "" {
				return fmt.Errorf("list[%v].content is empty", i)
			}
		}
		msg.MsgMenu = &msgMeta
	case message.MpMsgTypeWxCard:
		var msgMeta message.WxCardMeta
		msgMeta.CardID = buf.String()
		msg.WxCard = &msgMeta
	case message.MiniProgramMsgTypeMiniProgramPage:
		var msgMeta message.MiniProgramPageMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.AppID == "" {
			return errors.New("appid is empty")
		}
		if msgMeta.PagePath == "" {
			return errors.New("pagepath is empty")
		}
		if msgMeta.ThumbMediaID == "" {
			return errors.New("thumb_media_id is empty")
		}
		msg.MiniProgramPage = &msgMeta
	}

	if userAgent != "" {
		client.UserAgent = userAgent
	}

	if accessToken == "" {
		accessTokenResp, err := token.GetAccessToken(appID, appSecret)
		if err != nil {
			return err
		}
		accessToken = accessTokenResp.AccessToken
	}

	if err := message.SendCustomer(accessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
