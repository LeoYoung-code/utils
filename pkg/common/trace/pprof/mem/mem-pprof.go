package main

import (
	"math/rand"
	"strings"

	"github.com/pkg/profile"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	concat(100)
	// concat1(100)
}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func concat1(n int) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randomString(n))
	}
	return s.String()
}

// go run mem-pprof.go

// 2023/03/05 14:47:51 profile: memory profiling enabled (rate 1), /var/folders/yh/2mmdl1bn1jd99p0f21wbbch40000gp/T/profile2972808088/mem.pprof
// 2023/03/05 14:47:51 profile: memory profiling disabled, /var/folders/yh/2mmdl1bn1jd99p0f21wbbch40000gp/T/profile2972808088/mem.pprof

// go tool pprof -http=:8080 /var/folders/yh/2mmdl1bn1jd99p0f21wbbch40000gp/T/profile2972808088/mem.pprof
