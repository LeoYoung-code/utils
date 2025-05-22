package string

import (
	"go/token"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/samber/lo"
)

// UcFirst 将字符串首字母大写
func UcFirst(str string) string {
	if str == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(str)
	return string(unicode.ToUpper(r)) + str[size:]
}

// LcFirst 将字符串首字母小写
func LcFirst(str string) string {
	if str == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(str)
	return string(unicode.ToLower(r)) + str[size:]
}

// GoSanitized converts a string to a valid Go identifier.
func GoSanitized(s string) string {
	// Fast path: check if we can skip sanitization
	valid := true
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			valid = false
			break
		}
	}
	if valid {
		// Check for keyword conflict
		if token.Lookup(s).IsKeyword() {
			return "_" + s
		}
		return s
	}

	// Sanitize the input using strings.Builder for better performance
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	s = b.String()

	// Prepend '_' in the event of a Go keyword conflict or if
	// the identifier is invalid (does not start in the Unicode L category).
	r, _ := utf8.DecodeRuneInString(s)
	if token.Lookup(s).IsKeyword() || !unicode.IsLetter(r) {
		return "_" + s
	}
	return s
}

func b2s(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// b2s_new.go
func s2b(s string) (b []byte) {
	if len(s) == 0 {
		return nil
	}
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}

func generateRandStr(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return lo.RandomString(length, []rune(letters))
}
