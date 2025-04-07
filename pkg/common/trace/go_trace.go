package trace

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/trace"
)

func test() {

	// 从标准输入开始读取。
	r, err := trace.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var blocked int
	var blockedOnNetwork int
	for {
		// 读取事件。
		ev, err := r.ReadEvent()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// 处理它。
		if ev.Kind() == trace.EventStateTransition {
			st := ev.StateTransition()
			if st.Resource.Kind == trace.ResourceGoroutine {
				from, to := st.Goroutine()

				// 查找阻塞的goroutine，并计数。
				if from.Executing() && to == trace.GoWaiting {
					blocked++
					if strings.Contains(st.Reason, "network") {
						blockedOnNetwork++
					}
				}
			}
		}
	}
	// 打印我们发现的内容。
	p := 100 * float64(blockedOnNetwork) / float64(blocked)
	fmt.Printf("%2.3f%% instances of goroutines blocking were to block on the network\n", p)
}
