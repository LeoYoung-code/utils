package channel

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		// 模拟工作的耗时
		// 这里可以是实际的工作逻辑
		result := job * 2
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- result
	}
}

func doWork() {
	jobs := make(chan int)
	results := make(chan int)

	// 创建三个工作协程
	numWorkers := 3
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			worker(workerID, jobs, results)
			wg.Done()
		}(i)
	}

	// 获取结果 先读取结果(不阻塞),否则生产者数量超过消费者数量后会panic (向无缓冲通道 超过等待消费者的写入数据会阻塞)
	go func() {
		for result := range results {
			fmt.Printf("Received result: %d\n", result)
		}
	}()

	// 发送任务到通道中
	numJobs := 10
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs) // 关闭通道，表示任务发送完成

	// 等待所有工作协程完成
	wg.Wait()
	close(results) // 关闭结果通道，表示结果接收完成

}
