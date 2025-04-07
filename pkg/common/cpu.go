package common

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func printCPU() {
	done := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			// 获取 CPU 利用率 (每 100 毫秒获取一次)
			percent, err := cpu.Percent(100*time.Millisecond, false)
			if err != nil {
				log.Fatalf("get CPU usage: %v\n", err)
				return
			}

			for _, v := range percent {
				fmt.Printf("CPU usage is %.2f%%\n", v)
			}
		}

		// 模拟程序结束后通过 channel 发送通知
		done <- struct{}{}
	}()

	// 模拟一些 CPU 密集型任务
	for i := 0; i < 10000000; i++ {
		_ = math.Sqrt(rand.Float64())
	}

	<-done
	close(done)
}
