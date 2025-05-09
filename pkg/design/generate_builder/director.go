package main

// Director 是建造者模式中的指挥者，负责组织建造过程
// 它使用 IBuilder 接口来构建对象，而不关心具体的建造细节
type Director struct {
	builder IBuilder
}

// NewDirector 创建一个新的指挥者实例
func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// SetBuilder 设置指挥者使用的建造者
func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

// BuildHouse 通过建造者构建一个房子
// 它按特定顺序调用建造者的方法
func (d *Director) BuildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// DismantleHouse 通过建造者拆除一个房子
// 这是一个示例方法，显示建造者也可以用于拆解过程
func (d *Director) DismantleHouse() House {
	// 拆除步骤的实现
	// 实际应用中这里会有不同的顺序或操作
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
