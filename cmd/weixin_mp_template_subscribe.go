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
	nameScene = "scene"
	nameTitle = "title"
)

var (
	scene string
	title string
)

// weiXinMpTplSubCmd 微信公众号一次性订阅消息
var weiXinMpTplSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin mp template subscribe message",
	Long:    `publish weixin mp template subscribe message`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinMpSendTemplateSubscribe(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinMpTplCmd.AddCommand(weiXinMpTplSubCmd)

	weiXinMpTplSubCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinMpTplSubCmd.MarkFlagRequired(nameOpenID)

	weiXinMpTplSubCmd.Flags().StringVarP(&templateID, nameTemplateID, "p", "", "weixin template id (required)")
	weiXinMpTplSubCmd.MarkFlagRequired(nameTemplateID)

	weiXinMpTplSubCmd.Flags().StringVar(&scene, nameScene, "", "weixin subscribe scene (required)")
	weiXinMpTplSubCmd.MarkFlagRequired(nameScene)

	weiXinMpTplSubCmd.Flags().StringVar(&title, nameTitle, "", "weixin message title (required)")
	weiXinMpTplSubCmd.MarkFlagRequired(nameTitle)

	weiXinMpTplSubCmd.Flags().StringVarP(&url, nameUrl, "u", "", "url")
	weiXinMpTplSubCmd.Flags().StringToStringVarP(&mini, nameMini, "m", nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}

// WeiXinMpSendTemplateSubscribe 发送微信公众号一次性订阅消息
func WeiXinMpSendTemplateSubscribe(args []string) error {

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
			if len(v.Value) > 200 {
				return fmt.Errorf("data %v.value maximum length is within 200", k)
			}
		}
	}

	if len(title) > 15 {
		return fmt.Errorf("flag %q maximum length is within 15", nameTitle)
	}

	msg := message.TemplateSubscribeMessage{
		ToUser:     openID,
		TemplateID: templateID,
		Data:       dataItem,
		URL:        url,
		Scene:      scene,
		Title:      title,
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

	if err := message.SendTemplateSubscribe(accessToken, &msg); err != nil {
		return err
	}

	return nil
}
