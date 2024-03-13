package encrypt

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strconv"
	"strings"
)

func signature() string {
	// 请自行填入以下参数
	securityKey := "a0565dbfd68889e600d4225a9151c588"
	// timestamp := time.Time{}.Unix()
	timestamp := 111
	// 生成随机的 nonce
	// nonce, err := generateRandomNonce()
	// if err != nil {
	// 	return ""
	// }
	nonce := 123

	values := []string{
		securityKey,
		strconv.Itoa(int(timestamp)),
		strconv.Itoa(nonce),
	}
	sort.Strings(values)

	sb := strings.Join(values, "")
	hash := sha1.Sum([]byte(sb))
	signature := hex.EncodeToString(hash[:])
	return signature
}

func generateRandomNonce() (int, error) {
	// 生成一个64位的随机数
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}

	// 将字节切片转换为整数
	nonce := int(uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 |
		uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56)

	return nonce, nil
}
