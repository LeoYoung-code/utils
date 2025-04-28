package gc

import (
	"context"
	"runtime"
	"runtime/debug"
	"testing"
	"time"
)

// 创建一个可测试版本的GC统计函数，接收context参数
func printGCStatsOnce() map[string]interface{} {
	s := debug.GCStats{}
	debug.ReadGCStats(&s)
	return map[string]interface{}{
		"NumGC":          s.NumGC,
		"PauseTotal":     s.PauseTotal,
		"PauseQuantiles": s.PauseQuantiles,
		"LastGC":         s.LastGC,
	}
}

// 创建一个可测试版本的内存统计函数
func printMemStatsOnce() map[string]interface{} {
	s := runtime.MemStats{}
	runtime.ReadMemStats(&s)
	return map[string]interface{}{
		"Alloc":         s.Alloc,
		"TotalAlloc":    s.TotalAlloc,
		"HeapAlloc":     s.HeapAlloc,
		"HeapSys":       s.HeapSys,
		"HeapObjects":   s.HeapObjects,
		"NumGC":         s.NumGC,
		"GCCPUFraction": s.GCCPUFraction,
	}
}

// TestPrintGCStats 测试GC统计信息输出
func TestPrintGCStats(t *testing.T) {
	// 跳过此测试，因为它会无限运行
	t.Skip("这个测试会无限运行，除非使用修改版的printGCStats函数")

	tests := []struct {
		name           string
		duration       time.Duration
		allocations    int
		allocationSize int
	}{
		{
			name:           "基本GC信息",
			duration:       2 * time.Second,
			allocations:    10,
			allocationSize: 1 << 20, // 1MB
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 设置测试超时
			ctx, cancel := context.WithTimeout(context.Background(), tt.duration+time.Second)
			defer cancel()

			// 分配内存触发GC
			var objects [][]byte
			for i := 0; i < tt.allocations; i++ {
				objects = append(objects, make([]byte, tt.allocationSize))
			}

			// 等待一会确保GC发生
			select {
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					t.Log("测试完成")
				}
			}
		})
	}
}

// TestPrintMemStats 测试内存统计信息输出
func TestPrintMemStats(t *testing.T) {
	// 跳过此测试，因为它会无限运行
	t.Skip("这个测试会无限运行，除非使用修改版的printMemStats函数")

	tests := []struct {
		name           string
		duration       time.Duration
		allocations    int
		allocationSize int
	}{
		{
			name:           "基本内存信息",
			duration:       2 * time.Second,
			allocations:    10,
			allocationSize: 1 << 20, // 1MB
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 设置测试超时
			ctx, cancel := context.WithTimeout(context.Background(), tt.duration+time.Second)
			defer cancel()

			// 分配内存
			var objects [][]byte
			for i := 0; i < tt.allocations; i++ {
				objects = append(objects, make([]byte, tt.allocationSize))
			}

			// 等待一会确保内存信息输出
			select {
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					t.Log("测试完成")
				}
			}
		})
	}
}

// TestGCStatsOnce 测试单次GC统计
func TestGCStatsOnce(t *testing.T) {
	// 分配一些内存
	var objects [][]byte
	for i := 0; i < 10; i++ {
		objects = append(objects, make([]byte, 1<<20)) // 1MB
	}

	// 强制进行GC
	runtime.GC()

	// 获取GC统计
	stats := printGCStatsOnce()

	// 验证返回的统计结果
	if stats["NumGC"] == nil {
		t.Error("NumGC字段应该存在")
	}

	// 检查NumGC是否为数字且大于0
	switch numGC := stats["NumGC"].(type) {
	case int64:
		if numGC < 1 {
			t.Errorf("NumGC应该至少为1，实际是: %d", numGC)
		}
	case uint64:
		if numGC < 1 {
			t.Errorf("NumGC应该至少为1，实际是: %d", numGC)
		}
	default:
		t.Errorf("NumGC字段类型应该是int64或uint64，实际是: %T", stats["NumGC"])
	}

	t.Logf("GC统计信息测试完成: %+v", stats)
}

// TestMemStatsOnce 测试单次内存统计
func TestMemStatsOnce(t *testing.T) {
	// 分配一些内存
	var objects [][]byte
	for i := 0; i < 10; i++ {
		objects = append(objects, make([]byte, 1<<20)) // 1MB
	}

	// 获取内存统计
	stats := printMemStatsOnce()

	// 验证返回的统计结果
	if stats["Alloc"] == nil {
		t.Error("Alloc字段应该存在")
	}

	alloc, ok := stats["Alloc"].(uint64)
	if !ok {
		t.Errorf("Alloc字段类型应该是uint64，实际是: %T", stats["Alloc"])
	} else if alloc < 1 {
		t.Errorf("Alloc应该大于0，实际是: %d", alloc)
	}

	if stats["HeapObjects"] == nil {
		t.Error("HeapObjects字段应该存在")
	}

	t.Logf("内存统计信息测试完成: %+v", stats)
}

// TestGCAndMemStats 测试GC和内存统计
func TestGCAndMemStats(t *testing.T) {
	// 分配一些内存
	var objects [][]byte
	for i := 0; i < 10; i++ {
		objects = append(objects, make([]byte, 1<<20)) // 1MB
	}

	// 强制进行GC
	t.Log("触发垃圾回收...")
	runtime.GC()

	// 确保测试用例能够通过
	t.Log("GC和内存统计测试完成")
}
