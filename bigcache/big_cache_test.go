package bigcache

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSetGoCache(t *testing.T) {
	go printMemUsage()
	for i := 0; i < 200000; i++ {
		k := fmt.Sprintf("ab:s:1:1:10000:%d", i)
		SetGoCache(k, []byte{})
	}
}

func printMemUsage() {
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
		fmt.Printf("\tNumGC = %v\n", m.NumGC)
		time.Sleep(time.Microsecond * 200)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
