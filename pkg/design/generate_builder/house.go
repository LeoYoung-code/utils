package main

import (
	"fmt"
)

// House 表示通过建造者模式创建的房屋
// 包含窗户类型、门类型和楼层数量等属性
type House struct {
	windowType string // 窗户类型
	doorType   string // 门类型
	floor      int    // 楼层数量
}

// String 方法用于返回房屋的字符串表示
func (h House) String() string {
	return fmt.Sprintf("门类型: %s, 窗户类型: %s, 楼层数: %d", h.doorType, h.windowType, h.floor)
}
