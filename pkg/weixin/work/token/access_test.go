package token

import (
	"testing"

	"github.com/lenye/pmsg/pkg/weixin/work"
)

func TestGetAccessToken(t *testing.T) {
	corpID := work.CorpID
	corpSecret := work.CorpSecret
	accessTokenResp, err := GetAccessToken(corpID, corpSecret)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", accessTokenResp)
}
