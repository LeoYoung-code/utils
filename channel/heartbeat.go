package channel

import (
	"log"
	"os"
	"time"

	utime "utils/time"
)

func heartbeat() {
	heartbeat := time.NewTicker(300 * time.Microsecond)
	stop := time.NewTicker(1 * time.Second)
	defer heartbeat.Stop()

	for {
		select {
		case <-heartbeat.C:
			// 处理心跳
			log.Println("heartbeat")
		case <-stop.C:
			// 处理停止信号
			log.Println("stop")
			os.Exit(1)
		default:
			log.Println("worker", utime.TimeToString(nil))
		}
	}
}
