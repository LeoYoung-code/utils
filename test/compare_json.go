package test

import (
	"encoding/json"
	"fmt"
)

// 比较两个JSON对象的差异
func compareJSON(json1, json2 string) []string {
	var data1, data2 any
	// 解析第一个JSON字符串
	if err := json.Unmarshal([]byte(json1), &data1); err != nil {
		panic(err)
	}
	// 解析第二个JSON字符串
	if err := json.Unmarshal([]byte(json2), &data2); err != nil {
		panic(err)
	}

	// 比较两个数据结构
	return compareData(data1, data2, "")
}

// 递归比较两个数据结构
func compareData(data1, data2 any, path string) []string {
	var differences []string

	switch v1 := data1.(type) {
	case map[string]any:
		switch v2 := data2.(type) {
		case map[string]any:
			// 比较两个map的key
			keySet1 := make(map[string]struct{})
			for k := range v1 {
				keySet1[k] = struct{}{}
			}
			for k := range v2 {
				if _, ok := keySet1[k]; !ok {
					differences = append(differences, fmt.Sprintf("%s.%s", path, k))
				} else {
					delete(keySet1, k)
				}
			}
			for k := range keySet1 {
				differences = append(differences, fmt.Sprintf("%s.%s", path, k))
			}

			// 递归比较相同的key
			for k := range v1 {
				if _, ok := v2[k]; ok {
					differences = append(differences, compareData(v1[k], v2[k], fmt.Sprintf("%s.%s", path, k))...)
				} else {
					differences = append(differences, fmt.Sprintf("%s.%s", path, k))
				}
			}
		default:
			differences = append(differences, path)
		}
	case []any:
		switch v2 := data2.(type) {
		case []any:
			if len(v1) != len(v2) {
				differences = append(differences, path)
			} else {
				for i := range v1 {
					differences = append(differences, compareData(v1[i], v2[i], fmt.Sprintf("%s[%d]", path, i))...)
				}
			}
		default:
			differences = append(differences, path)
		}
	default:
		if data1 != data2 {
			differences = append(differences, path)
		}
	}

	return differences
}
