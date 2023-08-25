package test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/stretchr/testify/assert"
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

func TestReplace(t *testing.T) {
	in := []int{0, 1, 0, 1, 2, 3, 0}

	s1 := lo.Replace(in, 0, 42, 1)
	// []int{42dd, 1, 0, 1, 2, 3, 0}
	assert.Equal(t, []int{42, 1, 0, 1, 2, 3, 0}, s1)

	s2 := lo.Replace(in, -1, 42, 1)
	// []int{0, 1, 0, 1, 2, 3, 0}
	assert.Equal(t, []int{0, 1, 0, 1, 2, 3, 0}, s2)

	s3 := lo.Replace(in, 0, 42, 2)
	// []int{42, 1, 42, 1, 2, 3, 0}
	assert.Equal(t, []int{42, 1, 42, 1, 2, 3, 0}, s3)

	s4 := lo.Replace(in, 0, 42, -1)
	// []int{42, 1, 42, 1, 2, 3, 42}
	assert.Equal(t, []int{42, 1, 42, 1, 2, 3, 42}, s4)
}

func TestCompact(t *testing.T) {
	in := []string{"", "foo", "", "bar", ""}
	s1 := lo.Compact[string](in)
	// []string{"foo", "bar"}
	assert.Equal(t, []string{"foo", "bar"}, s1)

	in2 := []int{0, 1, 0, 1, 2, 3, 0}
	s2 := lo.Compact[int](in2)
	// []int{1, 1, 2, 3}
	assert.Equal(t, []int{1, 1, 2, 3}, s2)
}

func TestMapKeys(t *testing.T) {
	m := lo.MapKeys(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(_ int, v int) string {
		return strconv.FormatInt(int64(v)*100, 10)
	})
	assert.Equal(t, map[string]int{"100": 1, "200": 2, "300": 3, "400": 4}, m)
	// map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
}
