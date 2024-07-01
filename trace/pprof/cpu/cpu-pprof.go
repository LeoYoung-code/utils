package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	err := pprof.StartCPUProfile(f)
	if err != nil {
		return
	}
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

// go run cpu-pprof.go

// go tool pprof cpu.pprof  // 进入交互模式 top  查看消耗最多的函数 list bubbleSort  查看函数的具体代码 quit 退出

// go tool pprof -http=:8080 cpu.pprof  // 生成一个web页面
