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

package variable

var (
	UserAgent string

	Secret      string
	AccessToken string

	AppID     string
	AppSecret string

	ToUser          string
	ToParty         string
	ToTag           string
	ToParentUserID  string
	ToStudentUserID string
	ToMobile        string

	Mini map[string]string

	TemplateID       string
	Url              string
	ClientMsgID      string
	Color            string
	MsgType          string
	KfAccount        string
	Page             string
	MiniProgramState string
	Language         string
	Scene            string
	Title            string
	MediaType        string
	MediaID          string

	CorpID     string
	CorpSecret string

	AgentID                int64
	Safe                   int
	EnableIDTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
	ChatID                 string
	ToAll                  int
	RecvScope              int
	OpenKfID               string
	MsgID                  string

	AtUser   string
	AtMobile string
	IsAtAll  bool

	IsRaw bool
)
