package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	words := []string{"1", "2", "3", "4"}
	quoted := Map(words, func(s string) int64 {
		i, _ := strconv.ParseInt(s, 10, 64)
		return i
	})
	t.Log(quoted)
}

func TestReduce(t *testing.T) {
	numbers := []int{4, 9, 16, 25}
	quoted := Reduce(numbers, 0, func(i, j int) int {
		return i + j
	})
	t.Log(quoted)
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dual := Filter(numbers, true, func(in int) bool {
		return in%2 == 0
	})
	t.Log(dual)
}

func TestGetStructName(t *testing.T) {
	type bar struct{}
	foo := bar{}
	t.Log(GetStructName(foo))
}

func TestSet(t *testing.T) {
	numbers := []int{1, 2, 3, 3, 1}
	set := Set(numbers)
	assert.Equal(t, []int{1, 2, 3}, set)

	s := []string{"1", "1", "3", "2", "3"}
	ss := Set(s)
	assert.Equal(t, []string{"1", "3", "2"}, ss)
}

func Test_SliceIntersect_string(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{"string1", []string{"1", "2", "3", "1"}, []string{"2", "10"}, []string{"2"}},
		{"string2", []string{"1", "2", "3"}, []string{"5", "10"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := SliceIntersect(tt.slice1, tt.slice2)
			assert.Equalf(t, tt.want, w, "hasIntersectString(%v, %v)", tt.slice1, tt.slice2)
		})
	}
}

func Test_SliceIntersect_int(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{"int1", []int{1, 2, 3, 1}, []int{2, 10}, []int{2}},
		{"int2", []int{1, 2, 3}, []int{5, 10}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := SliceIntersect(tt.slice1, tt.slice2)
			assert.Equalf(t, tt.want, w, "hasIntersectString(%v, %v)", tt.slice1, tt.slice2)
		})
	}
}

func TestMapFormat(t *testing.T) {
	type args struct {
		arr []int
		f   func(T1 int, T2 ...any) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MapFormat(tt.args.arr, tt.args.f), "MapFormat(%v, %v)", tt.args.arr, tt.args.f)
		})
	}
}
