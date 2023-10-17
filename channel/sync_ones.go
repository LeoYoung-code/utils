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

/*
*
onceDo会等待f执行完毕后才返回，这期间其他执行once。Do函数的goroutine（如上面运行结果中的goroutine 2~5)将会阻塞等待；

Do函数返回后，后续的goroutine再执行Do函数将不再执行f并立即返回（如上面运行结果中的goroutine0）；

即便在函数f中出现panic，sync。Once原语也会认为once。Do 执行完毕，后续对once。Do的调用将不再执行f。
*/
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
