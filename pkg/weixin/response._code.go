package weixin

const CodeOK = 0

// ResponseCode 微信响应操作错误信息
type ResponseCode struct {
	ErrorCode    int64  `json:"errcode"`          // 出错返回码，为0表示成功，非0表示调用失败
	ErrorMessage string `json:"errmsg,omitempty"` // 返回码提示语
}

// Succeed 操作是否成功
func (p ResponseCode) Succeed() bool {
	return p.ErrorCode == CodeOK
}
