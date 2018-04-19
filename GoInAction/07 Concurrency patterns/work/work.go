// Jason Waldrip 协助完成了这个示例
// work 包管理一个 goroutine 池来完成工作
package work

import "sync"

// Worker 必须满足接口类型才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个 goroutine 池，这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

// New 使用固定数量的 goroutine 来创建一个新工作池
func New(maxGoroutines int) *Pool {
	// 创建一个 Pool 类型的值，并使用无缓冲的通道来初始化 work 字段
	p := Pool {
		work: make(chan Worker),
	}

	// 初始化 WaitGroup 需要等待的数量
	p.wg.Add(maxGoroutines)
	// 创建同样数量的 goroutine，这些 goroutine 只接收 Worker 类型的接口值，并调用这个值的 Task 方法
	for i := 0; i < maxGoroutines; i ++ {
		go func() {
			// 会一起阻塞，走到从 work 通道收到一个 Worker 接口值。如果接收到一个值，就会执行这个值的 Task 方法。
			// 一旦 work 通道被关闭，for range 循环就会结束，并调用 WaitGroup 的 Done 方法。然后 goroutine 终止
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run 提交工作到工作池
// 接受一个 Worker 类型的接口值，并将这个值通过 work 通道发送。
// 由于 work 通道是一个无缓冲的通道，调用者必须等待工作池里的某个 goroutine 接收到这个值才会返回。
// 这正是想要的，这样可以保证调用的 Run 返回时，提交的工作已经开始执行
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有 goroutine 停止工作
// 首先，关闭 work 通道，这会导致所有池里的 goroutine 停止工作，并调用 WaitGroup 的 Done 方法；
// 然后，Shutdown 方法调用 WaitGroup 的 Wait 方法，这会让 Shutdown 方法等待所有 goroutine 终止
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
