package gc

import (
	"fmt"
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
			fmt.Printf("GC: NumGC: %d,  PauseTotal: %d, Pause: %d, PauseEnd: %v, PauseQuantiles: %d, LastGC: %v\n", s.NumGC, s.PauseTotal, s.Pause, s.PauseEnd, s.PauseQuantiles, s.LastGC)
		}
	}
}
