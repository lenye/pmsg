// Copyright 2022-2024 The pmsg Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package weixin

import (
	"errors"
	"fmt"
)

const (
	CodeOK    = 0
	MessageOK = "ok"
)

var ErrRequest = errors.New("weixin request error")

// ResponseMeta 响应操作信息
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
