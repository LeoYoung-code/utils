package bigcache

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
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

// mockBigcache is a mock implementation of Bigcache
type mockBigcache struct {
}

func (*mockBigcache) New(ctx context.Context, config bigcache.Config) (*bigcache.BigCache, error) {
	return &bigcache.BigCache{}, nil
}

func TestNewCache(t *testing.T) {
	tests := []struct {
		name string
		mock func() *mockBigcache
		want *bigcache.BigCache
	}{
		{
			name: "Successful initialization",
			mock: func() *mockBigcache {
				return &mockBigcache{}
			},
			want: &bigcache.BigCache{},
		},
		// Here you can add more tests cases, for example, when the initialization fails
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			got := NewCache(ctx)
			if got == nil {
				t.Error("Cache should not be nil")
			}
		})
	}
}
