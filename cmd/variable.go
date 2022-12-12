// Copyright 2022 The pmsg Authors. All rights reserved.
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

package cmd

var (
	userAgent string

	secret      string
	accessToken string

	appID     string
	appSecret string

	toUser          string
	toParty         string
	toTag           string
	toParentUserID  string
	toStudentUserID string
	toMobile        string

	mini map[string]string

	templateID       string
	url              string
	clientMsgID      string
	color            string
	msgType          string
	kfAccount        string
	page             string
	miniProgramState string
	language         string
	scene            string
	title            string
	mediaType        string
	mediaID          string

	corpID     string
	corpSecret string

	agentID                int64
	safe                   int
	enableIDTrans          int
	enableDuplicateCheck   int
	duplicateCheckInterval int
	chatID                 string
	toAll                  int
	recvScope              int
	openKfID               string
	msgID                  string

	atUser   string
	atMobile string
	isAtAll  bool
)
