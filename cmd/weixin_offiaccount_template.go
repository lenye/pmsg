package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/offiaccount/message"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

const (
	nameOpenID      = "open_id"
	nameTemplateID  = "template_id"
	nameUrl         = "url"
	nameClientMsgID = "client_msg_id"
	nameColor       = "color"

	nameMini         = "mini"
	nameMiniAppID    = "app_id"
	nameMiniPagePath = "page_path"
)

var (
	openID       string
	templateID   string
	url          string
	clientMsgID  string
	color        string
	mini         map[string]string
	miniAppID    string
	miniPagePath string
)

// weiXinOfficialAccountTplCmd 微信公众号模板消息
var weiXinOfficialAccountTplCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "publish weixin offiaccount template message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinOfficialAccountSendTemplate(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinOfficialAccountCmd.AddCommand(weiXinOfficialAccountTplCmd)

	weiXinSetAccessTokenFlags(weiXinOfficialAccountTplCmd)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountTplCmd.MarkFlagRequired(nameOpenID)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&templateID, nameTemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountTplCmd.MarkFlagRequired(nameTemplateID)

	weiXinOfficialAccountTplCmd.Flags().StringVarP(&url, nameUrl, "u", "", "url")
	weiXinOfficialAccountTplCmd.Flags().StringVar(&clientMsgID, nameClientMsgID, "", "weixin template client msg id")
	weiXinOfficialAccountTplCmd.Flags().StringVarP(&color, nameColor, "c", "", "weixin template color")

	weiXinOfficialAccountTplCmd.Flags().StringToStringVarP(&mini, nameMini, "m", nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}

// WeiXinOfficialAccountSendTemplate 发送微信公众号模板消息
func WeiXinOfficialAccountSendTemplate(args []string) error {

	if accessToken == "" {
		if appID == "" {
			return ErrMultiRequiredOne
		}
	}

	var dataItem map[string]message.TemplateDataItem
	buf := bytes.NewBufferString("")
	buf.WriteString(args[0])
	if buf.String() != "" {
		if err := json.Unmarshal(buf.Bytes(), &dataItem); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		for k, v := range dataItem {
			if v.Value == "" {
				return fmt.Errorf("data %v.value not set", k)
			}
		}
	}

	msg := message.TemplateMessage{
		ToUser:      openID,
		TemplateID:  templateID,
		Data:        dataItem,
		URL:         url,
		ClientMsgID: clientMsgID,
		Color:       color,
	}

	// 跳小程序
	if mini != nil {
		var ok bool
		miniAppID, ok = mini[nameMiniAppID]
		if !ok {
			return fmt.Errorf("mini flag %q not set", nameMiniAppID)
		}
		if miniAppID == "" {
			return fmt.Errorf("mini flag %q not set", nameMiniAppID)
		}

		miniPagePath, ok = mini[nameMiniPagePath]
		if !ok {
			return fmt.Errorf("mini flag %q not set", nameMiniPagePath)
		}
		if miniPagePath == "" {
			return fmt.Errorf("mini flag %q not set", nameMiniPagePath)
		}

		msg.MiniProgram = &message.MiniProgramMeta{
			AppID:    miniAppID,
			PagePath: miniPagePath,
		}
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

	if gotMsgID, err := message.SendTemplate(accessToken, &msg); err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("%v; msgid: %v", weixin.MessageOK, gotMsgID))
	}

	return nil
}
