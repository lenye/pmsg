package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin/mp/message"
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
	data         string
	clientMsgID  string
	color        string
	mini         map[string]string
	miniAppID    string
	miniPagePath string
)

// weiXinMpTplCmd 微信公众号模板消息
var weiXinMpTplCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "publish weixin mp template message",
	Long:    `publish weixin mp template message`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinMpSendTemplate(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinMpCmd.AddCommand(weiXinMpTplCmd)

	weiXinMpTplCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinMpTplCmd.MarkFlagRequired(nameOpenID)

	weiXinMpTplCmd.Flags().StringVarP(&templateID, nameTemplateID, "p", "", "weixin template id (required)")
	weiXinMpTplCmd.MarkFlagRequired(nameTemplateID)

	weiXinMpTplCmd.Flags().StringVarP(&url, nameUrl, "u", "", "url")
	weiXinMpTplCmd.Flags().StringVar(&clientMsgID, nameClientMsgID, "", "weixin template client msg id")
	weiXinMpTplCmd.Flags().StringVarP(&color, nameColor, "c", "", "weixin template color")

	weiXinMpTplCmd.Flags().StringToStringVarP(&mini, nameMini, "m", nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}

// WeiXinMpSendTemplate 发送微信公众号模板消息
func WeiXinMpSendTemplate(args []string) error {

	if userAgent != "" {
		client.UserAgent = userAgent
	}

	if accessToken == "" {
		if appID == "" {
			return ErrMultiRequiredOne
		}
	}

	var dataItem map[string]message.TemplateDataItem
	data = strings.Join(args, "")
	if data != "" {
		if err := json.Unmarshal([]byte(data), &dataItem); err != nil {
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

		msg.MiniProgram = &message.MiniProgram{
			AppID:    miniAppID,
			PagePath: miniPagePath,
		}
	}

	if accessToken == "" {
		accessTokenResp, err := token.GetAccessToken(appID, appSecret)
		if err != nil {
			return err
		}
		accessToken = accessTokenResp.AccessToken.Token
	}

	if gotMsgID, err := message.SendTemplate(accessToken, &msg); err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("msgid: %v", gotMsgID))
	}

	return nil
}
