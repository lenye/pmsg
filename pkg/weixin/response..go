package weixin

import (
	"errors"
	"fmt"
)

const (
	CodeOK    = 0
	MessageOK = "ok"
)

var ErrWeiXinRequest = errors.New("weixin request error")

// ResponseMeta 微信响应操作信息
type ResponseMeta struct {
	ErrorCode    int64  `json:"errcode"`          // 出错返回码，为0表示成功，非0表示调用失败
	ErrorMessage string `json:"errmsg,omitempty"` // 返回码提示语
}

func (t ResponseMeta) String() string {
	return fmt.Sprintf("errcode: %v, errmsg: %q", t.ErrorCode, t.ErrorMessage)
}

// Succeed 操作是否成功
func (t ResponseMeta) Succeed() bool {
	return t.ErrorCode == CodeOK
}
