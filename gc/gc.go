package gc

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"

	json "github.com/bytedance/sonic"
)

func printGCStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			m := map[string]interface{}{
				"NumGC":      s.NumGC,
				"PauseTotal": s.PauseTotal,
				// "Pause":          s.Pause,
				// "PauseEnd":       s.PauseEnd,
				"PauseQuantiles": s.PauseQuantiles,
				"LastGC":         s.LastGC,
			}
			b, _ := json.Marshal(m)
			fmt.Println(string(b))
		}
	}
}

func printMemStats() {
	t := time.NewTicker(time.Second)
	s := runtime.MemStats{}
	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(&s)
			m := map[string]interface{}{
				"Alloc":         s.Alloc,
				"TotalAlloc":    s.TotalAlloc,
				"Sys":           s.Sys,
				"Lookups":       s.Lookups,
				"Mallocs":       s.Mallocs,
				"Frees":         s.Frees,
				"HeapAlloc":     s.HeapAlloc,
				"HeapSys":       s.HeapSys,
				"HeapIdle":      s.HeapIdle,
				"HeapInuse":     s.HeapInuse,
				"HeapReleased":  s.HeapReleased,
				"HeapObjects":   s.HeapObjects,
				"StackInuse":    s.StackInuse,
				"StackSys":      s.StackSys,
				"MSpanInuse":    s.MSpanInuse,
				"MSpanSys":      s.MSpanSys,
				"MCacheInuse":   s.MCacheInuse,
				"MCacheSys":     s.MCacheSys,
				"BuckHashSys":   s.BuckHashSys,
				"GCSys":         s.GCSys,
				"OtherSys":      s.OtherSys,
				"NextGC":        s.NextGC,
				"LastGC":        s.LastGC,
				"PauseTotalNs":  s.PauseTotalNs,
				"PauseNs":       s.PauseNs,
				"PauseEnd":      s.PauseEnd,
				"NumGC":         s.NumGC,
				"NumForcedGC":   s.NumForcedGC,
				"GCCPUFraction": s.GCCPUFraction,
				"DebugGC":       s.DebugGC,
				"BySize":        s.BySize,
			}
			b, _ := json.Marshal(m)
			fmt.Println(string(b))
		}
	}
}
