package version

import (
	"strconv"
	"strings"
)

// GetSysVersion 系统版本号 字符串格式转int 如"10.2"转10
func GetSysVersion(versionStr string) int64 {
	if versionStr == "" {
		return 0
	}

	parts := strings.SplitN(versionStr, ".", 2)
	versionStr = parts[0]

	versionInt, err := strconv.ParseInt(versionStr, 10, 64)
	if err != nil {
		return 0
	}
	return versionInt
}
