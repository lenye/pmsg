package token

import (
	"fmt"
	"time"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

type AccessToken struct {
	AccessToken string    `json:"access_token"`        // 接口调用凭证
	ExpireIn    int64     `json:"expires_in"`          // 接口调用凭证有效时间，单位：秒
	ExpireAt    time.Time `json:"expire_at,omitempty"` // 到期时间
}

// AccessTokenResponse 响应
type AccessTokenResponse struct {
	weixin.ResponseCode
	AccessToken
}

const accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

// GetAccessToken 获取接口调用凭证
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
		return nil, fmt.Errorf("weixin request failed, uri=%q, response=%+v", url, resp.ResponseCode)
	}

	resp.AccessToken.ExpireAt = time.Now().Add(time.Second * time.Duration(resp.AccessToken.ExpireIn))

	return &resp.AccessToken, nil
}
