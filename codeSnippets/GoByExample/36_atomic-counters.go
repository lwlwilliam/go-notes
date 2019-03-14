/*
Go 中最主要的状态管理方式是通过通道间的沟通来完成的，在 34_worker-pools.go 工作池中碰到过，
还有一些其他的方法来管理状态的。

这里使用 sync/atomic 包在多个 Go 协程中进行原子计数
 */
package main

import (
	"sync/atomic"
	"time"
	"fmt"
	"runtime"
)

func main() {
	// 使用一个无符号整型数来表示（永远是正整数）这个计数器
	var ops uint64 = 0

	// 为了模拟并发更新，启动 50 个 Go 协程，对计数器每隔 1ms（应为非准确时间）进行一次加一操作
	for i := 0; i < 50; i ++ {
		go func() {
			for {
				// 使用 AddUint64 来让计数器自动增加，使用 & 语法来给出 ops 的内存地址
				atomic.AddUint64(&ops, 1)

				// 允许其他 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}

	// 等待一秒，让 ops 的自加操作执行一会
	time.Sleep(time.Second)

	// 为了在计数器还在被其他 Go 协程更新时，全安的使用它，通过 LoadUint64 将当前值的拷贝提取到 opsFinal 中
	// 和上面一样，需要给这个函数所取修士的内存地址 &ops
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
