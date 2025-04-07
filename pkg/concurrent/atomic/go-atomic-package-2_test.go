package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

/**
利用原子操作的无锁并发写的性能随着并发量的增大而小幅下降；

利用原子操作的无锁并发读的性能随着并发量增大有持续提升的趋势，并且性能约为读锁的100倍。
*/

type Config struct {
	sync.RWMutex
	data string
}

func BenchmarkRWMutexSet(b *testing.B) {
	config := Config{}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Lock()
			config.data = "hello"
			config.Unlock()
		}
	})
}

func BenchmarkRWMutexGet(b *testing.B) {
	config := Config{data: "hello"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.RLock()
			_ = config.data
			config.RUnlock()
		}
	})
}

func BenchmarkAtomicSet(b *testing.B) {
	var config atomic.Value
	c := Config{data: "hello"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Store(c)
		}
	})
}

func BenchmarkAtomicGet(b *testing.B) {
	var config atomic.Value
	config.Store(Config{data: "hello"})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = config.Load().(Config)
		}
	})
}
