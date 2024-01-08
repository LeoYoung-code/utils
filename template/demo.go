package template

import (
	"fmt"
	"os"
	"text/template"
)

// 定义一个结构体，用于在模板中存储数据
type Person struct {
	Name  string
	Age   int
	City  string
	Email string
}

func Echo() {
	// 创建一个模板字符串
	templateString := "Name: {{.Name}}, Age: {{.Age}}, City: {{.City}}, Email: {{.Email}}"

	// 解析模板字符串
	tmpl, err := template.New("personTemplate").Parse(templateString)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// 创建一个包含数据的结构体实例
	person := Person{
		Name:  "John Doe",
		Age:   30,
		City:  "Example City",
		Email: "john.doe@example.com",
	}

	// 应用模板并将结果写入标准输出
	err = tmpl.Execute(os.Stdout, person)
	if err != nil {
		fmt.Println("Error applying template:", err)
		return
	}
}
