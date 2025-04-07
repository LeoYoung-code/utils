package atomic

import (
	"sync/atomic"
)

// ID 原子ID生成器
type ID struct {
	counter int64
}

// NewID 创建一个新的ID生成器
func NewID(start int64) *ID {
	return &ID{
		counter: start,
	}
}

// Next 获取下一个ID，原子增加
func (i *ID) Next() int64 {
	return atomic.AddInt64(&i.counter, 1)
}

// Get 获取当前ID值
func (i *ID) Get() int64 {
	return atomic.LoadInt64(&i.counter)
}

// Set 设置ID值
func (i *ID) Set(value int64) {
	atomic.StoreInt64(&i.counter, value)
}

// Bool 原子布尔值
type Bool struct {
	value uint32
}

// NewBool 创建一个新的原子布尔值
func NewBool(value bool) *Bool {
	var i uint32 = 0
	if value {
		i = 1
	}
	return &Bool{value: i}
}

// Set 设置布尔值
func (b *Bool) Set(value bool) {
	var i uint32 = 0
	if value {
		i = 1
	}
	atomic.StoreUint32(&b.value, i)
}

// Get 获取布尔值
func (b *Bool) Get() bool {
	return atomic.LoadUint32(&b.value) != 0
}

// Toggle 切换布尔值并返回新值
func (b *Bool) Toggle() bool {
	var old uint32
	var new uint32

	for {
		old = atomic.LoadUint32(&b.value)
		new = 1 - old
		if atomic.CompareAndSwapUint32(&b.value, old, new) {
			break
		}
	}

	return new != 0
}
