package string

import (
	"bytes"
	"strings"
	"unicode"
	"unsafe"
)

// StrToBytes 将字符串转换为字节切片，零拷贝
func StrToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// TrimSpace 去除字符串首尾空格
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// Split 分割字符串
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join 拼接字符串
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// ToUpper 转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Contains 判断字符串是否包含子串
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// CamelToSnake 驼峰命名转蛇形命名
func CamelToSnake(s string) string {
	var buffer bytes.Buffer
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// SnakeToCamel 蛇形命名转驼峰命名
func SnakeToCamel(s string) string {
	var buffer bytes.Buffer
	toUpper := false
	for _, r := range s {
		if r == '_' {
			toUpper = true
			continue
		}
		if toUpper {
			buffer.WriteRune(unicode.ToUpper(r))
			toUpper = false
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// FirstCharToUpper 首字母大写
func FirstCharToUpper(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
