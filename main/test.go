package main

import (
	"fmt"
	"strconv"

	"github.com/LeoYoung-code/cast"
	oldCast "github.com/spf13/cast"
)

func main() {
	for i := 0; i <= 20; i++ {
		fmt.Println("旧包: ", oldCast.ToInt(fmt.Sprintf("0%d", i)))
		fmt.Println("新包: ", cast.ToInt(fmt.Sprintf("0%d", i)))
		nums, _ := strconv.Atoi(fmt.Sprintf("0%d", i))
		fmt.Println("原始: ", nums)
		fmt.Println("-------------------")
	}
}
