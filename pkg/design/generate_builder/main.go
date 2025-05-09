package main

import "fmt"

// main 函数演示建造者模式的使用方法
func main() {
	// 创建不同类型的建造者
	normalBuilder := GetBuilder(NormalBuilderType)
	iglooBuilder := GetBuilder(IglooBuilderType)

	// 创建指挥者并使用普通建造者构建房屋
	director := NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	// 输出普通房屋的信息
	fmt.Println("普通房屋信息:")
	fmt.Println(normalHouse)

	// 切换至冰屋建造者并构建冰屋
	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	// 输出冰屋的信息
	fmt.Println("\n冰屋信息:")
	fmt.Println(iglooHouse)
}
