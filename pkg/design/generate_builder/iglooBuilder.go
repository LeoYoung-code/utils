package main

// IglooBuilder 是实现 IBuilder 接口的冰屋建造者
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// newIglooBuilder 创建一个新的冰屋建造者实例
func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

// setWindowType 设置冰屋的窗户类型
func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

// setDoorType 设置冰屋的门类型
func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

// setNumFloor 设置冰屋的楼层数量
func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

// getHouse 返回构建好的冰屋实例
func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
