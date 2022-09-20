package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/mp/message"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

// weiXinMpSubCmd 微信公众号订阅通知消息
var weiXinMpSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin mp subscribe message",
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinMpBizSendSubscribe(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinMpCmd.AddCommand(weiXinMpSubCmd)

	weiXinMpSubCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinMpSubCmd.MarkFlagRequired(nameOpenID)

	weiXinMpSubCmd.Flags().StringVarP(&templateID, nameTemplateID, "p", "", "weixin template id (required)")
	weiXinMpSubCmd.MarkFlagRequired(nameTemplateID)

	weiXinMpSubCmd.Flags().StringVar(&page, namePage, "", "page")
	weiXinMpSubCmd.Flags().StringToStringVarP(&mini, nameMini, "m", nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}

// WeiXinMpBizSendSubscribe 发送微信公众号订阅通知消息
func WeiXinMpBizSendSubscribe(args []string) error {

	if accessToken == "" {
		if appID == "" {
			return ErrMultiRequiredOne
		}
	}

	var dataItem map[string]message.SubscribeDataItem
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

	msg := message.SubscribeMessage{
		ToUser:     openID,
		TemplateID: templateID,
		Page:       page,
		Data:       dataItem,
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

	if err := message.BizSendSubscribe(accessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
