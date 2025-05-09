package main

// NormalBuilder 是实现 IBuilder 接口的普通房屋建造者
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// newNormalBuilder 创建一个新的普通房屋建造者实例
func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

// setWindowType 设置普通房屋的窗户类型
func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

// setDoorType 设置普通房屋的门类型
func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

// setNumFloor 设置普通房屋的楼层数量
func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

// getHouse 返回构建好的普通房屋实例
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
