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

package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"time"
)

// timestamp 当前时间戳，单位是秒

// Validate 验证
func Validate(signStr, timestamp, secret string) (bool, error) {
	t, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return false, err
	}

	timeGap := time.Since(time.Unix(t, 0))
	if math.Abs(timeGap.Hours()) > 1 {
		return false, fmt.Errorf("specified timestamp is expired")
	}

	ourSign := Sign(timestamp, secret)
	return ourSign == signStr, nil
}

// Sign 签名
func Sign(timestamp string, secret string) string {
	stringToSign := timestamp + "\n" + secret
	h := hmac.New(sha256.New, []byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
