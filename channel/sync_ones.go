package channel

import (
	"log"
	"sync"
	"time"
)

type foo struct {
}

var once sync.Once
var instance *foo

func getInstance(id int) *foo {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("goroutine-%d: caught a panic: %s", id, e)
		}
	}()
	log.Printf("goroutine-%d: enter GetInstance\n", id)
	once.Do(func() {
		instance = &foo{}
		time.Sleep(3 * time.Second)
		log.Printf("goroutine-%d: the addr of instance is %p\n", id, instance)
		panic("panic in once.Do function")
	})
	return instance
}

func testOnes() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			inst := getInstance(i)
			log.Printf("goroutine-%d: the addr of instance returned is %p\n", i, inst)
			wg.Done()
		}(i + 1)
	}
	time.Sleep(5 * time.Second)
	inst := getInstance(0)
	log.Printf("goroutine-0: the addr of instance returned is %p\n", inst)

	wg.Wait()
	log.Printf("all goroutines exit\n")
}
