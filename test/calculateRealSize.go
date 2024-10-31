package test

import (
	"fmt"
	"reflect"
)

// 递归计算结构体大小，包括动态分配部分
func calculateRealSize(value interface{}) uintptr {
	return calculateSize(reflect.ValueOf(value))
}

func calculateSize(v reflect.Value) uintptr {
	size := uintptr(0)

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		if !v.IsNil() {
			size += calculateSize(v.Elem())
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			size += calculateSize(v.Field(i))
		}
	case reflect.Slice, reflect.Map, reflect.String:
		size += v.Type().Size() // 包含切片/映射/字符串结构体头的大小
		if v.Kind() == reflect.Slice {
			for i := 0; i < v.Len(); i++ {
				size += calculateSize(v.Index(i))
			}
		} else if v.Kind() == reflect.Map {
			for _, key := range v.MapKeys() {
				size += calculateSize(key)
				size += calculateSize(v.MapIndex(key))
			}
		} else if v.Kind() == reflect.String {
			size += uintptr(len(v.String()))
		}
	default:
		size += v.Type().Size()
	}

	return size
}

// Example 示例结构体
type Example struct {
	Name    string
	Age     int
	Friends []string
}

func main() {
	example := Example{
		Name:    "Alice",
		Age:     30,
		Friends: []string{"Bob", "Charlie", "Daisy"},
	}

	realSize := calculateRealSize(example)
	fmt.Printf("结构体 Example 的真实大小: %d 字节\n", realSize)
}
