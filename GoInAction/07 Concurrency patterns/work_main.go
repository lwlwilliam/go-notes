// 这个示例程序展示如何使用 work 包
// 创建一个 goroutine 池并完成工作
package main

import (
	"./work"
	"log"
	"sync"
	"time"
)

// names 提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现 Worker 接口（必须实现 Worker 接口才能使用工作池）
func (m *namePrinter) Task() {
	log.Println(m.name)

	// 等待一秒是为了让测试程序运行的速度慢一些，以便看到并发效果
	time.Sleep(time.Second)
}

func main() {
	// 使用两个 goroutine 来创建工作池，也就是工作池只有两个 goroutine 来处理工作
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	// 提交 10*len(names) 个工作
	// names 切片里的每个名字都会创建 10 个 goroutine 来提交任务，
	// 这样就会有一堆 goroutine 互相竞争，将任务提交到池里
	for i := 0; i < 10; i ++ {
		// 迭代 names 切片
		for _, name := range names {
			// 创建一个 namePrinter 并提供指定的名字
			np := namePrinter {
				name: name,
			}

			go func() {
				// 将任务提交执行，当 Run 返回时就知道任务已经处理完成
				p.Run(&np)

				// Run 返回后，goroutine 将 WaitGroup 的计数递减，并终止 goroutine
				wg.Done()
			}()
		}
	}

	// 等待所有创建的 goroutine 提交它们的工作
	wg.Wait()

	// 一旦 Wait 返回，就会调用工作池的 Shutdown 来关闭工作池。
	// 让工作池停止工作，等待所有现有的工作完成
	// 该例中，最多只会等待两个工作的完成
	p.Shutdown()
}
