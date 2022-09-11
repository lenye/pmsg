package version

import (
	"time"
)

var (
	Version   = "dev"
	CodeName  = "app"             // app名称
	BuildTime = "20180303.192030" // 编译时间   $(date +%Y%m%d.%H%M%S)
	BuildGit  = "git"             // 版本号     $(git rev-parse HEAD)
	StartTime = time.Now()
)
