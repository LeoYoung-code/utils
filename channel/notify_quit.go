package channel

import (
	"fmt"
	"time"

	utime "utils/time"
)

// notifyQuit 通知退出
func notifyQuit() {
	doneChan := make(chan int)
	for i := 0; i < 5; i++ {
		go func(worker int) {
			for {
				select {
				case <-doneChan: // 监听退出信号
					fmt.Println("worker", worker, "done")
					return
				default:
					fmt.Println("worker", worker, utime.TimeToString(nil))
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}
	time.Sleep(5 * time.Second)
	// 发出退出信号，这⾥关闭 channel 以后，上⾯的 case <-doneChan 就会发现已经不会阻塞 了，从⽽执⾏return
	close(doneChan)
	time.Sleep(2 * time.Second)
}
