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

package version

import (
	"fmt"
	"runtime"
	"time"
)

var (
	Version   = "dev"                  // 版本
	AppName   = "app"                  // 名称
	BuildTime = "2018-09-01T16:03:59Z" // 编译时间   $(date +%Y%m%d.%H%M%S)
	BuildGit  = "git"                  // 版本号     $(git rev-parse HEAD)
	StartTime = time.Now()
)

const versionTemplate = `%s
 Version:    %s
 Git commit: %s
 Built:      %s
 Go version: %s
 OS/Arch:    %s/%s`

func Print() string {
	return fmt.Sprintf(versionTemplate, AppName, Version, BuildGit, BuildTime, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
