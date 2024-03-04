package test

import (
	"strings"

	"github.com/spf13/cast"
)

func testBase(i int64, base int64) int64 {
	if base == 0 {
		return i
	}
	res := make([]int64, 0, 5)
	num := i
	for {
		quotient := num / base
		remainder := num % base
		res = append(res, remainder)
		if quotient == 0 {
			break
		}
		num = quotient
	}
	reversed := reverseSlice(res)
	return cast.ToInt64(strings.Join(cast.ToStringSlice(reversed), ""))
}

func reverseSlice(slice []int64) []int64 {
	length := len(slice)
	reversed := make([]int64, length)

	for i, v := range slice {
		reversed[length-i-1] = v
	}

	return reversed
}
