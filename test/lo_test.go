package test

import (
	"fmt"
	"strconv"
	"testing"

	lop "github.com/samber/lo/parallel"
)

func TestTimes(t *testing.T) {
	s := lop.Times(100, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	fmt.Println(s)
}

func TestGroupBy(t *testing.T) {
	r := lop.GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		a := i % 3
		return a
	})
	fmt.Println(r)
}

func TestPartitionBy(t *testing.T) {
	partitions := lop.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Println(partitions)
}
