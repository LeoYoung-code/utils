package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
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

// Encrypt 加密参数
func Encrypt(content string) string {
	secretSalt := "242ccb8230d709e1"
	randNum := getRandNumber()
	process := func(crypt []byte) []byte {
		return []byte(fmt.Sprintf("%v%v", randNum, string(crypt)))
	}
	result, _ := AesEncrypt(content, secretSalt, randNum, process)
	return result
}

func getRandNumber() string {
	numSlice := make([]string, 0)
	for i := 0; i < 16; i++ {
		num := rand.Int31n(10)
		numSlice = append(numSlice, strconv.Itoa(int(num)))
	}
	return strings.Join(numSlice, "")
}
