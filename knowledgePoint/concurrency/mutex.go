// 了解 mutex 之前，先理解 critical section (https://en.wikipedia.org/wiki/Critical_section) 的概念很重要。
// 一个并发程序，修改共享资源的代码不应该被多个 goroutine 在同一时间访问。

// 假设有一段代码使变量 x 递增：x = x + 1。这段代码被单独一个 goroutine 访问不会出现任何问题。如果多个 goroutine 
// 并发执行，为了简化就先假设有两个 goroutine 并发执行以上代码。把系统执行的过程进行简化，该段代码执行步骤如下：
//		1.	获取当前的 x 值；
//		2.	计算 x + 1；
// 		3.	把第二步得到的值赋给 x；

// 当这三个步骤由一个 goroutine 执行，一切正常。下面来讨论两个 goroutine 并发执行这段代码会发生什么。
//		1.	x 初始值为 0；
//		2.	goroutine1 获取当前 x 值为 0，计算 x + 1；
//		3.	系统上下文切换到 goroutine2，goroutine2 获取当前 x 值为 0，计算 x + 1；
//		4.	系统上下文切换到 goroutine1，goroutine1 把计算后的值 1 赋给 x，x 的值变为 1；
//		5.	系统上下文切换到 goroutine2，goroutine2 把计算后的值 1 赋给 x，x 的值变为 1；
//		6.	goroutine1 和 goroutine2 都执行结束，x 的最终值为 1；

// 以上只是其中的一种并发情景，还可能会发生以下情景。
//		1.	x 初始值为 0；
//		2.	goroutine1 获取当前 x 值为 0，计算 x + 1，把计算后的结果赋给 x，x 的值变为 1；
//		3.	系统上下文切换到 goroutine2，goroutine2 获取当前 x 值为 1，计算 x + 1，把计算后的结果赋给 x，x 的值变为 2；
//		4.	goroutine1 和 goroutine2 都扫行结束，x 的最终值为 2；

// 以上两种情景，x 的最终结果取决于系统上下文的切换，这不是期望的结果，这也称为 race condition（竞争状态）。

// 使用 mutex 可以避免竞争状态。
// mutex 用于提供一种锁机制以确保在任何时间点都只有唯一一个 goroutine 执行 critical section（临界资源），从而避免竞争状态。
// mutex 在 sync 包，在 mutex 定义了两个方法 Lock 和 Unlock。任何在 Lock 和 Unlock 方法之间的代码都只会被一个 goroutine
// 执行，因此避免了竞争状态。
// 如果一个 goroutine 已经持有了锁，如果另一个 goroutine 试图请求锁，这个新的 goroutine 会阻塞直到锁取消。
package main

import (
	"fmt"
	"sync"
)

var x = 0
// 注意：m 一定要传址。如果传值，每个 goroutine 都会有自己的 mutext，竞争状态还会出现。
func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i ++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}