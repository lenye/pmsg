package version

import (
	"fmt"
	"runtime"
	"time"
)

var (
	Version   = "dev"
	AppName   = "app"                  // app名称
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
