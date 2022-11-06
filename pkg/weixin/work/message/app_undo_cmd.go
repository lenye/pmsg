package message

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/work/token"
)

type CmdWorkUndoAppParams struct {
	UserAgent   string
	AccessToken string
	CorpID      string
	CorpSecret  string
	MsgID       string
}

// CmdWorkUndoApp 撤回企业微信应用消息
func CmdWorkUndoApp(arg *CmdWorkUndoAppParams) error {

	msg := UndoAppMessage{
		MsgID: arg.MsgID,
	}

	client.UserAgent = arg.UserAgent

	if arg.AccessToken == "" {
		accessTokenResp, err := token.GetAccessToken(arg.CorpID, arg.CorpSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if err := UndoApp(arg.AccessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}
