package token

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

type CmdTokenParams struct {
	UserAgent string
	AppID     string
	AppSecret string
}

// CmdGetAccessToken 获取微信接口调用凭证
func CmdGetAccessToken(arg CmdTokenParams) error {
	if arg.UserAgent != "" {
		client.UserAgent = arg.UserAgent
	}

	accessTokenResp, err := GetAccessToken(arg.AppID, arg.AppSecret)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, accessTokenResp))

	return nil
}
