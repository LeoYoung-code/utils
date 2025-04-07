package channel

import (
	"fmt"
)

func waitQuit() {
	waitChan := make(chan struct{}, 5)
	for i := 0; i < 5; i++ {
		go func(worker int) {
			for wi := 0; wi < 5; wi++ {
				fmt.Println("worker", worker, wi)
			}
			// 通知 worker 退出
			waitChan <- struct{}{}
		}(i)
	}
	// 监听所有 worker 退出信号
	for i := 0; i < 5; i++ {
		<-waitChan
	}

}
