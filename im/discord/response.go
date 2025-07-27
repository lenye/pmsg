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
	"errors"
	"fmt"
	"strings"
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
	var sb []string

	if t.Code != nil {
		sb = append(sb, fmt.Sprintf("code: %d", *t.Code))
	}
	if t.Message != "" {
		sb = append(sb, fmt.Sprintf("message: %s", t.Message))
	}
	if t.RetryAfter != nil {
		sb = append(sb, fmt.Sprintf("retry_after: %f", *t.RetryAfter))
	}
	if t.Global != nil {
		sb = append(sb, fmt.Sprintf("global: %t", *t.Global))
	}

	return strings.Join(sb, ", ")
}

// Succeed 操作是否成功
func (t *ResponseMeta) Succeed() bool {
	if t.Code == nil {
		return false
	}
	return *t.Code == CodeOK
}
