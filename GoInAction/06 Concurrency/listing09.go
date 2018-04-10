// 这个示例程序展示如何在程序里造成竞争状态
// 实际上不希望出现这种情况
package main

import (
	"fmt"
	"runtime"

	//"runtime"
	"sync"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter int

	// wg 用来等待程序结束
	wg sync.WaitGroup
)

func main() {
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	// 函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

	for count := 0; count < 2; count ++ {
		// 捕获 counter 的值
		value := counter
		fmt.Printf("id: %d; value: %d\n", id, value)

		fmt.Printf("id: %d **********\n", id)

		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()

		// 增加本地 value 变量的值
		value ++

		// 将该值保存回 counter
		counter = value
		fmt.Printf("id: %d; counter: %d\n", id, counter)
	}
}
