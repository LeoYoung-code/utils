package template

import (
	"bytes"
	"testing"
	"text/template"
)

// ...

func TestGenerateTemplate(t *testing.T) {
	// 定义期望的输出
	expectedOutput := "Name: John Doe, Age: 30, City: Example City, Email: john.doe@example.com"

	// 创建一个模板字符串
	templateString := "Name: {{.Name}}, Age: {{.Age}}, City: {{.City}}, Email: {{.Email}}"

	// 解析模板字符串
	tmpl, err := template.New("personTemplate").Parse(templateString)
	if err != nil {
		t.Fatal("Error parsing template:", err)
	}

	// 创建一个包含数据的结构体实例
	person := Person{
		Name:  "John Doe",
		Age:   30,
		City:  "Example City",
		Email: "john.doe@example.com",
	}

	// 使用字节数组缓冲区捕获模板执行结果
	var resultBuffer bytes.Buffer
	err = tmpl.Execute(&resultBuffer, person)
	if err != nil {
		t.Fatal("Error applying template:", err)
	}

	// 检查实际输出是否与预期输出相匹配
	actualOutput := resultBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("Expected: %s, Got: %s", expectedOutput, actualOutput)
	}
}
