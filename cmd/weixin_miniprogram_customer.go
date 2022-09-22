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

// weiXinMiniProgramCustomerCmd 发送微信小程序客服消息
var weiXinMiniProgramCustomerCmd = &cobra.Command{
	Use:     "customer",
	Aliases: []string{"kf"},
	Short:   "publish weixin miniprogram customer message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinMiniSendCustomer(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinMiniProgramCmd.AddCommand(weiXinMiniProgramCustomerCmd)

	weiXinSetAccessTokenFlags(weiXinMiniProgramCustomerCmd)

	weiXinMiniProgramCustomerCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinMiniProgramCustomerCmd.MarkFlagRequired(nameOpenID)

	weiXinMiniProgramCustomerCmd.Flags().StringVar(&msgType, nameMsgType, "", "message type (required)")
	weiXinMiniProgramCustomerCmd.MarkFlagRequired(nameMsgType)
}

// WeiXinMiniSendCustomer 发送微信小程序客服消息
func WeiXinMiniSendCustomer(args []string) error {

	if accessToken == "" {
		if appID == "" {
			return ErrMultiRequiredOne
		}
	}

	if err := message.ValidateMiniProgramMsgType(msgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", nameMsgType, err)
	}

	msg := message.CustomerMessage{
		ToUser:  openID,
		MsgType: msgType,
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(args[0])
	switch msgType {
	case message.MiniProgramMsgTypeText:
		var msgMeta message.TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case message.MiniProgramMsgTypeImage:
		var msgMeta message.ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case message.MiniProgramMsgTypeLink:
		var msgMeta message.LinkMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.Description == "" {
			return errors.New("description is empty")
		}
		if msgMeta.Url == "" {
			return errors.New("url is empty")
		}
		if msgMeta.ThumbUrl == "" {
			return errors.New("thumb_url is empty")
		}
		msg.Link = &msgMeta
	case message.MiniProgramMsgTypeMiniProgramPage:
		var msgMeta message.MiniProgramPageMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.Title == "" {
			return errors.New("title is empty")
		}
		if msgMeta.AppID != "" {
			return errors.New("no appid required")
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
