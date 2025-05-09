package main

// IBuilder 定义了建造者模式的接口
// 包含设置窗户类型、设置门类型、设置楼层数量和获取房屋实例的方法
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

// BuilderType 定义了建造者类型的常量
type BuilderType string

const (
	// NormalBuilderType 普通房屋建造者类型
	NormalBuilderType BuilderType = "normal"
	// IglooBuilderType 冰屋建造者类型
	IglooBuilderType BuilderType = "igloo"
)

// GetBuilder 根据建造者类型返回对应的建造者实例
// 工厂方法模式的应用
func GetBuilder(builderType BuilderType) IBuilder {
	switch builderType {
	case NormalBuilderType:
		return newNormalBuilder()
	case IglooBuilderType:
		return newIglooBuilder()
	default:
		return nil
	}
}
