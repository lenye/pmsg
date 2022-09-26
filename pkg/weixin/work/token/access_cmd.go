package token

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

type CmdWorkTokenParams struct {
	UserAgent  string
	CorpID     string
	CorpSecret string
}

// CmdWorkGetAccessToken 获取企业微信接口调用凭证
func CmdWorkGetAccessToken(arg *CmdWorkTokenParams) error {

	client.UserAgent = arg.UserAgent

	accessTokenResp, err := token.GetAccessToken(arg.CorpID, arg.CorpSecret)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, accessTokenResp))

	return nil
}
