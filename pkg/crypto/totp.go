package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"time"
)

// TOTP 配置参数
const (
	interval   = 30                 // 时间窗口（秒）
	codeDigits = 6                  // 验证码位数（可以调整为 8、7 等）
	base32Key  = "JBSWY3DPEHPK3PXP" // 示例密钥（需要与客户端同步）
)

// 计算时间片（Time Step）
func timeStep() int64 {
	return time.Now().Unix() / interval
}

// 生成 HMAC-SHA1 哈希
func generateHMAC(secret []byte, counter int64) []byte {
	// 将时间片转换为 8 字节的 big-endian 字节数组
	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(counter))

	// 使用密钥和时间片进行 HMAC-SHA1 加密
	hmacSha1 := hmac.New(sha1.New, secret)
	hmacSha1.Write(counterBytes)
	return hmacSha1.Sum(nil)
}

// 截取哈希值并生成 TOTP 验证码
func truncateAndGenerateCode(hmacResult []byte) int {
	// 获取最后一个字节的低4位作为偏移量
	offset := hmacResult[len(hmacResult)-1] & 0x0F

	// 从偏移位置取 4 字节，并转换为 31 位正整数
	code := binary.BigEndian.Uint32(hmacResult[offset:offset+4]) & 0x7FFFFFFF

	// 计算 10 的 codeDigits 次幂，用于取模生成验证码
	mod := int(math.Pow10(codeDigits))
	return int(code % uint32(mod))
}

// 生成 TOTP 验证码
func generateTOTP(secret string) (string, error) {
	if len(secret) == 0 {
		return "", fmt.Errorf("密钥不能为空")
	}

	// 解码 Base32 格式密钥
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// 生成 HMAC 并计算验证码
	hmacResult := generateHMAC(key, timeStep())
	code := truncateAndGenerateCode(hmacResult)

	// 格式化为指定位数的验证码（前置补零）
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", codeDigits), code), nil
}

func main1() {
	// 生成 TOTP 验证码
	code, err := generateTOTP(base32Key)
	if err != nil {
		log.Fatalf("生成 TOTP 出错: %v", err)
	}
	fmt.Printf("当前 %d 位 TOTP 验证码: %s\n", codeDigits, code)
}
