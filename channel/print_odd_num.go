package channel

import (
	"fmt"
	"sync"
	"time"
)

func printOddNum() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 打印奇数
	go func() {
		for n := range ch1 {
			fmt.Printf("channel-1 print %d\n", n)
			// 奇数 + 1 等于下一个要打印的偶数
			// 发送到 channel 2
			ch2 <- n + 1

			if n == 9 {
				break
			}
		}
		close(ch1)
	}()

	// 打印偶数
	go func() {
		for n := range ch2 {
			fmt.Printf("channel-2 print %d\n", n)
			if n == 10 {
				break
			}

			// 偶数 + 1 等于下一个要打印的奇数
			// 发送到 channel 1
			ch1 <- n + 1
		}
		close(ch2)
	}()

	ch1 <- 1

	// 等待 goroutine 执行完成
	time.Sleep(time.Second)
}

// rev  接收 channel
// send 发送 channel
// no   channel 编号 (提高打印消息的可读性)
// last channel 接收的最后一个数字, 超过该数字时结束打印
func printChar(rev <-chan int, send chan<- int, no, last int) {
	for i := range rev {
		if i > last {
			break
		}
		fmt.Printf("channel-%d print %d\n", no, i)
		send <- i + 1
	}
	close(send)
}

// cond sync.Mutex 互斥锁
// cur  指定当前执行的 goroutine 编号
// no   channel 编号 (提高打印消息的可读性)
// last channel 接收的最后一个数字, 超过该数字时结束打印
func printChar1(mu *sync.Mutex, cur *int, no, last int) {
	for i := no; i <= last; {
		mu.Lock()
		if *cur != no {
			// 如果当前执行 goroutine 编号不是自身
			mu.Unlock()
			continue
		}

		fmt.Printf("channel-%d print %d\n", no, i)

		// 指定当前执行的 goroutine 编号为另外一个 goroutine
		*cur = 3 - *cur
		i = i + 2
		// 唤醒另外一个 goroutine
		mu.Unlock()
	}
}

// cond sync.Cond 条件变量
// cur  指定当前执行的 goroutine 编号
// no   channel 编号 (提高打印消息的可读性)
// last channel 接收的最后一个数字, 超过该数字时结束打印
func printChar2(cond *sync.Cond, cur *int, no, last int) {
	for i := no; i <= last; i = i + 2 {
		cond.L.Lock()
		for *cur != no {
			// 如果当前执行 goroutine 编号不是自身
			// 等待被唤醒
			cond.Wait()
		}

		fmt.Printf("channel-%d print %d\n", no, i)

		// 指定当前执行的 goroutine 编号为另外一个 goroutine
		*cur = 3 - *cur
		// 唤醒另外一个 goroutine
		cond.Signal()
		cond.L.Unlock()
	}
}

func printNum() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// 打印奇数
	go func() {
		defer wg.Done()
		printChar(ch1, ch2, 1, 9)
	}()

	// 打印偶数
	go func() {
		defer wg.Done()
		printChar(ch2, ch1, 2, 10)
	}()

	ch1 <- 1
	wg.Wait()
}

func printNumMutex() {
	mu := &sync.Mutex{}

	// 当前执行的 goroutine 编号
	cur := 1

	var wg sync.WaitGroup
	wg.Add(2)

	// 打印奇数
	go func() {
		defer wg.Done()
		printChar1(mu, &cur, 1, 9)
	}()

	// 打印偶数
	go func() {
		defer wg.Done()
		printChar1(mu, &cur, 2, 10)
	}()

	wg.Wait()
}

func printNumForCond() {
	cond := sync.NewCond(&sync.Mutex{})

	// 当前执行的 goroutine 编号
	cur := 1

	var wg sync.WaitGroup
	wg.Add(2)

	// 打印奇数
	go func() {
		defer wg.Done()
		printChar2(cond, &cur, 1, 9)
	}()

	// 打印偶数
	go func() {
		defer wg.Done()
		printChar2(cond, &cur, 2, 10)
	}()

	wg.Wait()
}
