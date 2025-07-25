// Copyright 2022-2025 The pmsg Authors. All rights reserved.
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

package discord

import (
	"encoding/json"
	"errors"
)

const (
	CodeOK    = 0
	MessageOK = "ok"
)

var ErrRequest = errors.New("discord request failed")

type ResponseMeta struct {
	Code       *int     `json:"code,omitzero"`
	Message    string   `json:"message"`
	RetryAfter *float64 `json:"retry_after,omitzero"`
	Global     *bool    `json:"global,omitzero"`
}

func (t *ResponseMeta) String() string {
	buf, _ := json.Marshal(t)
	return string(buf)
}

// Succeed 操作是否成功
func (t *ResponseMeta) Succeed() bool {
	if t.Code == nil {
		return false
	}
	return *t.Code == CodeOK
}
