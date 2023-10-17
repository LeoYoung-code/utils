package channel

import (
	"bytes"
	"sync"
)

/*
*
它是goroutine并发安全的，可以被多个goroutine同时使用

放入该缓存池中的数据对象的生命是暂时的，随时都可能被垃圾回收掉；

缓存池中的数据对象是可以重复利用的，这样可以在一定程度上降低数据对象重新分配的频度，减轻GC的压力；

syncPool为每个P（goroutine调度模型中的P）单独建立一个local缓存池，进一步降低高并发下对锁的争抢。
*/
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func writeBufFromPool(data string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(data)
	bufPool.Put(b)
}
func writeBufFromNew(data string) *bytes.Buffer {
	b := new(bytes.Buffer)
	b.WriteString(data)
	return b
}
