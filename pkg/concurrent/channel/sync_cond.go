package channel

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

var ready bool

func workerCond(i int) {
	fmt.Printf("workerCond %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("workerCond %d: works done\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal *sync.Cond) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			groupSignal.L.Lock()
			for !ready {
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			fmt.Printf("workerCond %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- struct{}{}
	}()
	return c
}

func testCond() {
	fmt.Println("start a group of workers...")
	groupSignal := sync.NewCond(&sync.Mutex{})
	c := spawnGroup(workerCond, 5, groupSignal)

	time.Sleep(5 * time.Second) // 模拟ready前的准备工作
	fmt.Println("the group of workers start to work...")

	groupSignal.L.Lock()
	ready = true
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-c
	fmt.Println("the group of workers work done!")
}
