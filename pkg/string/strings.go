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

// GoSanitized converts a string to a valid Go identifier.
func GoSanitized(s string) string {
	// Sanitize the input to the set of valid characters,
	// which must be '_' or be in the Unicode L or N categories.
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '_'
	}, s)

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
