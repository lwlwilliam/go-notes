/*
使用 Go 协程和通道实现一个工作池

执行这个程序，显示 9 个任务被多个 worker 执行。整个程序处理所有的任务仅执行了 3s 而不是 9s，
是因为 3 个 worker 是并行的
 */
package main

import (
	"fmt"
	"time"
)

// 这是将要在多个并发实例中支持的任务。
// 从 jobs 通道接收 int 类型的任务，把 int 类型的结果发送到 results 通道
func worker(id int, jobs <- chan int, results chan <- int) {
	// 每隔一秒执行一次 jobs 通道中的任务并把结果写入 results 通道
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		// 每个任务间隔 1s 来模仿一个耗时的任务
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 为了使用 worker 工作池并且收集他们的结果，需要 2 个通道
	// 任务通道
	jobs := make(chan int, 100)
	// 结果通道
	results := make(chan int, 100)

	// 开启 3 个 Go 协程处理 jobs 任务，初始是阻塞的，因为还没有传递任务
	for w := 1; w <= 3; w ++ {
		go worker(w, jobs, results)
	}

	// 往 jobs 通道中发送 9 个任务
	for j := 1; j <= 9; j ++ {
		jobs <- j
	}

	// 通过 close 这些通道来表示这些就是所有的任务了
	close(jobs)

	// 收集所有任务的返回值
	for a := 1; a <= 9; a ++ {
		<- results
	}
}