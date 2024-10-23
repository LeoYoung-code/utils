package version

import (
	"strings"

	"github.com/LeoYoung-code/cast"
)

// GetSysVersion 系统版本号 字符串格式转int 如"10.2"转10
func GetSysVersion(versionStr string) int64 {
	index := strings.Index(versionStr, ".")
	if index >= 0 {
		versionStr = versionStr[0:index]
	}
	versionInt, _ := cast.ToInt64E(versionStr)
	return versionInt
}
