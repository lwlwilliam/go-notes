// 这个示例程序展示如何用无缓冲通道来模拟 4 个 goroutine 间的接力比赛
package main

import (
	"fmt"
	"sync"
	"time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 为最后一位跑步者将计数加 1
	// 等待一个 goroutine 的返回
	wg.Add(1)

	// 第一位跑步者持有接力棒
	go Runner(baton)

	// 开始比赛
	baton <- 1

	// 等待比赛结束
	wg.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	// 新的接力队员
	var newRunner int

	// 等待接力棒
	runner := <- baton

	// 开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// 当前不是最后一个接力队员，下一个队员准备
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}