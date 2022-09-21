package token

import (
	"fmt"
	"time"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

type AccessToken struct {
	AccessToken string    `json:"access_token"`        // 微信接口调用凭证
	ExpireIn    int64     `json:"expires_in"`          // 微信接口调用凭证有效时间，单位：秒
	ExpireAt    time.Time `json:"expire_at,omitempty"` // 微信接口调用凭证到期时间
}

func (t AccessToken) String() string {
	if t.ExpireAt.IsZero() {
		return fmt.Sprintf("access_token: %q, expires_in: %v", t.AccessToken, t.ExpireIn)
	}
	return fmt.Sprintf("access_token: %q, expires_in: %v, expire_at: %q", t.AccessToken, t.ExpireIn, t.ExpireAt.Format(time.RFC3339))
}

// AccessTokenResponse 响应
type AccessTokenResponse struct {
	weixin.ResponseCode
	AccessToken
}

const accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

// GetAccessToken 获取微信接口调用凭证
// 正常情况下，微信会返回下述 JSON
// {"access_token":"ACCESS_TOKEN","expires_in":7200}
// 错误时微信会返回错误码等信息，JSON数据包示例如下:
// {"errcode":40013,"errmsg":"invalid appid"}
func GetAccessToken(appID, appSecret string) (*AccessToken, error) {
	url := fmt.Sprintf(accessTokenURL, appID, appSecret)
	var resp AccessTokenResponse
	_, err := client.GetJSON(url, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; uri: %q, %v", weixin.ErrWeiXinRequest, url, resp.ResponseCode)
	}

	resp.AccessToken.ExpireAt = time.Now().Add(time.Second * time.Duration(resp.AccessToken.ExpireIn))

	return &resp.AccessToken, nil
}
