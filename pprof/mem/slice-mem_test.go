package main

import (
	"runtime"
	"testing"
)

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024)
		ans = append(ans, f(origin))
		runtime.GC()
	}
	printMem(t)
	_ = ans
}

func Test_lastNumsByCopy(t *testing.T) {
	testLastChars(t, lastNumsBySlice)
}

func Test_lastNumsBySlice(t *testing.T) {
	testLastChars(t, lastNumsByCopy)
}
