package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 计算接口返回内容的MD5值
func getMD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// 模拟接口
func mockAPI() []byte {
	return []byte("Hello, World!")
}

func main() {
	// 获取接口返回内容
	data := mockAPI()

	// 计算MD5值
	md5Value := getMD5(data)

	// 如果接口内容不变，则返回空字符串
	if md5Value == "6a262e4271d49957c319e486a7e84b26" {
		fmt.Println("接口内容不变，返回空字符串")
	} else {
		fmt.Println("接口内容已更改，返回MD5值:", md5Value)
	}
}
