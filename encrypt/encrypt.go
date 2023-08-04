package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(content string) (md string) {
	if content == "" {
		return ""
	}
	h := md5.New()
	_, _ = io.WriteString(h, content)
	md = fmt.Sprintf("%x", h.Sum(nil))
	return
}
