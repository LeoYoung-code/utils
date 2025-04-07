package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main1() {

	// 创建 trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	// 启动 trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()
	// main
	fmt.Println("Hello World")
}
