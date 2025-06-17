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

// IsEmpty 检查字符串是否为空或只包含空白字符
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsBlank 检查字符串是否为空
func IsBlank(s string) bool {
	return len(s) == 0
}

// Reverse 反转字符串（支持Unicode）
func Reverse(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ContainsIgnoreCase 检查字符串是否包含子字符串（忽略大小写）
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// Truncate 截断字符串到指定长度，如果超出长度则添加省略号
func Truncate(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}

// PadLeft 左侧填充字符串到指定长度
func PadLeft(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	pad := strings.Repeat(string(padChar), length-len(s))
	return pad + s
}

// PadRight 右侧填充字符串到指定长度
func PadRight(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	pad := strings.Repeat(string(padChar), length-len(s))
	return s + pad
}

// RemoveSpaces 移除字符串中的所有空格
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// IsNumeric 检查字符串是否只包含数字
func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// CamelCase 转换为驼峰命名（首字母小写）
func CamelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	// 先转换为蛇形命名，再转换为驼峰
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	if len(words) == 0 {
		return s
	}

	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			result += UcFirst(strings.ToLower(words[i]))
		}
	}
	return result
}

// PascalCase 转换为帕斯卡命名（首字母大写的驼峰）
func PascalCase(s string) string {
	camel := CamelCase(s)
	return UcFirst(camel)
}

// SnakeCase 转换为蛇形命名（下划线分隔，全小写）
func SnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 && result[len(result)-1] != '_' {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		} else {
			if len(result) > 0 && result[len(result)-1] != '_' {
				result = append(result, '_')
			}
		}
	}

	// 清理首尾的下划线
	resultStr := string(result)
	resultStr = strings.Trim(resultStr, "_")
	return resultStr
}

// KebabCase 转换为短横线命名（连字符分隔，全小写）
func KebabCase(s string) string {
	snake := SnakeCase(s)
	return strings.ReplaceAll(snake, "_", "-")
}

// Title 将字符串转换为标题格式（每个单词首字母大写）
func Title(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = UcFirst(strings.ToLower(word))
	}
	return strings.Join(words, " ")
}

// Center 居中对齐字符串
func Center(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}

	totalPad := length - len(s)
	leftPad := totalPad / 2
	rightPad := totalPad - leftPad

	return strings.Repeat(string(padChar), leftPad) + s + strings.Repeat(string(padChar), rightPad)
}

// Count 统计子字符串在字符串中出现的次数
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// HasPrefix 检查字符串是否以指定前缀开始（忽略大小写）
func HasPrefixIgnoreCase(s, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix))
}

// HasSuffix 检查字符串是否以指定后缀结束（忽略大小写）
func HasSuffixIgnoreCase(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}

// WordCount 统计字符串中的单词数量
func WordCount(s string) int {
	return len(strings.Fields(s))
}

// Left 获取字符串左侧指定长度的子字符串
func Left(s string, length int) string {
	if length <= 0 {
		return ""
	}
	runes := []rune(s)
	if length >= len(runes) {
		return s
	}
	return string(runes[:length])
}

// Right 获取字符串右侧指定长度的子字符串
func Right(s string, length int) string {
	if length <= 0 {
		return ""
	}
	runes := []rune(s)
	if length >= len(runes) {
		return s
	}
	return string(runes[len(runes)-length:])
}

// IsAlpha 检查字符串是否只包含字母
func IsAlpha(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric 检查字符串是否只包含字母和数字
func IsAlphaNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
