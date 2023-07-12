package utils

import (
	"go/token"
	"strings"
	"unicode"
	"unicode/utf8"
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
