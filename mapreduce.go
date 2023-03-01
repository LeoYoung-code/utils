package utils

// Map 数组元素遍历
func Map[T1 any, T2 any](arr []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(arr))
	for i, elem := range arr {
		result[i] = f(elem)
	}
	return result
}

// func MapFormat[T1 any, T2 int](arr []T1, f func(T1, T2) T1) []T1 {

// MapFormat 数组元素遍历
func MapFormat(arr []int, f func(T1 int, T2 ...any) int) []int {
	list := make([]int, 0, len(arr))
	for _, elem := range arr {
		list = append(list, f(elem, elem))
	}
	return list
}

// Reduce 数组数据合成
func Reduce[T1 any, T2 any](arr []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, elem := range arr {
		result = f(result, elem)
	}
	return result
}

// Filter 数组元素过滤
func Filter[T any](arr []T, in bool, f func(T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, elem := range arr {
		choose := f(elem)
		if (in && choose) || (!in && !choose) {
			result = append(result, elem)
		}
	}
	return result
}

// Set 去重
func Set[T comparable](arr []T) []T {
	m := make(map[T]struct{}, len(arr))
	list := make([]T, 0, len(arr))
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			list = append(list, v)
			m[v] = struct{}{}
		}
	}
	return list
}

// slice2 中是否包含 slice1 元素
type mapKeys interface {
	string | int | int8 | int16 | int32 | int64 | float32 | float64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func SliceIntersect[T mapKeys](slice1, slice2 []T) []T {
	m := make(map[T]bool, 0)
	r := make([]T, 0)
	for _, s1 := range slice1 {
		m[s1] = true
	}
	for _, s2 := range slice2 {
		if _, ok := m[s2]; ok {
			r = append(r, s2)
		}
	}
	return r
}

// FilterByFunc 数组元素过滤  符合函数条件的元素
func FilterByFunc[T any](arr []T, f func(T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, elem := range arr {
		if f(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// GetSliceDiff  获取两个 slice 的差集
func GetSliceDiff[T any, C comparable](p1, p2 []T, f func(T) C) []T {
	m2 := make(map[C]struct{}, len(p2))
	for _, item := range p2 {
		m2[f(item)] = struct{}{}
	}
	p1temp := make([]T, 0, len(p1))
	for _, item := range p1 {
		if _, ok := m2[f(item)]; !ok {
			p1temp = append(p1temp, item)
		}
	}
	return p1temp
}

// GetSliceIntersect [T any, C comparable]获取两个 slice 的交集
func GetSliceIntersect[T any, C comparable](p1, p2 []T, f func(T) C) []T {
	m2 := make(map[C]struct{}, len(p2))
	for _, item := range p2 {
		m2[f(item)] = struct{}{}
	}
	p1temp := make([]T, 0, len(p1))
	for _, item := range p1 {
		if _, ok := m2[f(item)]; ok {
			p1temp = append(p1temp, item)
		}
	}
	return p1temp
}
