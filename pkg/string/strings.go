package string

import (
	"go/token"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/samber/lo"
)

// UcFirst 将字符串的首字母转换为大写
// 例如：hello -> Hello
func UcFirst(str string) string {
	strLen := len(str)
	if strLen == 0 {
		return ""
	} else if strLen == 1 {
		return strings.ToUpper(str)
	} else {
		return strings.ToUpper(str[:1]) + str[1:]
	}
}

// LcFirst 将字符串的首字母转换为小写
// 例如：Hello -> hello
func LcFirst(str string) string {
	strLen := len(str)
	if strLen == 0 {
		return ""
	} else if strLen == 1 {
		return strings.ToLower(str)
	} else {
		return strings.ToLower(str[:1]) + str[1:]
	}
}

// GoSanitized 将字符串转换为有效的Go标识符
// 会将非字母、非数字字符转换为下划线，并处理关键字冲突
func GoSanitized(s string) string {
	// 将非字母、非数字字符转换为下划线
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '_'
	}, s)

	// 处理关键字冲突或非字母开头的标识符
	r, _ := utf8.DecodeRuneInString(s)
	if token.Lookup(s).IsKeyword() || !unicode.IsLetter(r) {
		return "_" + s
	}
	return s
}

// BytesToString 将字节切片转换为字符串，零拷贝
// 警告：结果字符串与原字节切片共享相同的内存，修改原字节切片会影响返回的字符串
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}

// StringToBytes 将字符串转换为字节切片，零拷贝
// 警告：结果字节切片与原字符串共享相同的内存，不可修改返回的字节切片
func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// GenerateRandomString 生成指定长度的随机字符串
// 字符集包括大小写字母和数字
// 注意：length 必须大于 0，否则返回空字符串
func GenerateRandomString(length int) string {
	if length <= 0 {
		return ""
	}
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return lo.RandomString(length, []rune(letters))
}

// 以下是为了兼容旧代码而保留的函数，建议使用上面对应的新函数

// b2s 是 BytesToString 的别名，保留以保持兼容性
// 已弃用，请使用 BytesToString
func b2s(b []byte) string {
	return BytesToString(b)
}

// s2b 是 StringToBytes 的别名，保留以保持兼容性
// 已弃用，请使用 StringToBytes
func s2b(s string) (b []byte) {
	return StringToBytes(s)
}
