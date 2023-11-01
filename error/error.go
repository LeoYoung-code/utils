package errorTest

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

const basePath = "/Users/staff/project"

func NewError(err error) error {
	_, file, line, _ := runtime.Caller(1)
	relFile, err2 := filepath.Rel(basePath, file)
	if err2 != nil {
		// 如果无法获取相对路径，就使用原始的文件路径
		relFile = file
	}
	return errors.Wrap(err, fmt.Sprintf("%s:%d", relFile, line))
}
