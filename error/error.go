package errorTest

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

var basePath string

func init() {
	fullPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 分割路径
	parts := strings.Split(fullPath, "/")

	// 关键字
	key := "project"
	// 重新组合路径直到关键字
	basePath = "/"
	for _, part := range parts {
		if part == "" {
			continue
		}
		basePath = basePath + part
		if part == key {
			break
		}
		basePath = basePath + "/"
	}
}

func NewError(err error) error {
	_, file, line, _ := runtime.Caller(1)
	relFile, err2 := filepath.Rel(basePath, file)
	if err2 != nil {
		// 如果无法获取相对路径，就使用原始的文件路径
		relFile = file
	}
	return errors.Wrap(err, fmt.Sprintf("%s:%d", relFile, line))
}
