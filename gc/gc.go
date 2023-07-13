package gc

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func PrintGCStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf(
				"GC: NumGC: %d,  PauseTotal: %d, Pause: %d, PauseEnd: %v, PauseQuantiles: %d, LastGC: %v\n",
				s.NumGC,
				s.PauseTotal,
				s.Pause,
				s.PauseEnd,
				s.PauseQuantiles,
				s.LastGC,
			)
		}
	}
}

func PrintMemStats() {
	t := time.NewTicker(time.Second)
	s := runtime.MemStats{}
	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(&s)
			fmt.Printf(
				"Mem: Alloc: %d, TotalAlloc: %d, Sys: %d, Lookups: %d, Mallocs: %d, Frees: %d, HeapAlloc: %d, HeapSys: %d, HeapIdle: %d, HeapInuse: %d, HeapReleased: %d, HeapObjects: %d, StackInuse: %d, StackSys: %d, MSpanInuse: %d, MSpanSys: %d, MCacheInuse: %d, MCacheSys: %d, BuckHashSys: %d, GCSys: %d, OtherSys: %d, NextGC: %d, LastGC: %d, PauseTotalNs: %d, PauseNs: %d, PauseEnd: %d, NumGC: %d, NumForcedGC: %d, GCCPUFraction: %f, DebugGC: %t, BySize: %v\n",
				s.Alloc,
				s.TotalAlloc,
				s.Sys,
				s.Lookups,
				s.Mallocs,
				s.Frees,
				s.HeapAlloc,
				s.HeapSys,
				s.HeapIdle,
				s.HeapInuse,
				s.HeapReleased,
				s.HeapObjects,
				s.StackInuse,
				s.StackSys,
				s.MSpanInuse,
				s.MSpanSys,
				s.MCacheInuse,
				s.MCacheSys,
				s.BuckHashSys,
				s.GCSys,
				s.OtherSys,
				s.NextGC,
				s.LastGC,
				s.PauseTotalNs,
				s.PauseNs,
				s.PauseEnd,
				s.NumGC,
				s.NumForcedGC,
				s.GCCPUFraction,
				s.DebugGC,
				s.BySize,
			)
		}
	}
}
