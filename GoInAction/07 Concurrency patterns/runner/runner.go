/*
Gabriel Aszalos 协助完成了这个示例
runner 包管理处理任务的运行和生命周期

该程序展示了依据调度运行的无人值守的面向任务的程序，及其所使用的并发模式。在设计上，可支持以下终止点：
* 程序可以在分配的时间内完成工作，正常终止；
* 程序没有及时完成工作，"自杀"；
* 接收到操作系统发送的中断事件，程序立刻试图清理状态并停止工作；
 */
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
// 声明了 3 个通道，用来辅助管理程序的生命周期，以及用来表示顺序执行的不同任务的函数切片
type Runner struct {
	// interrupt 通道报告从操作系统发送的信号
	// Signal 用来描述操作系统发送的信号，其底层实现通常会依赖操作系统的具体实现：在 Unix 系统上是 syscall.Signal
	interrupt chan os.Signal

	// complete 通道报告处理任务已经完成
	// 如果执行任务时发生了错误，会发回一个 error 接口类型的值。如果没有发生错误，会通过这个通道发回一个 nil 值作为 error 接口值
	complete chan error

	// timeout 报告处理任务已经超时
	// 这个通道用来管理执行任务的时间。如果从这个通道接收到一个 time.Time 的值，这个程序就会试图清理状态并停止工作
	timeout <- chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会有接收到操作系统的中断事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	return &Runner{
		// 被初始化为缓冲区容量为 1 的通道。这可以保证通道至少能接收一个来自语言运行时的 os.Signal 值，确保语言运行时发送这个事件的时候
		// 不会被阻塞。如果 goroutine 没有准备好接收这个值，这个值就会被丢弃。
		interrupt: make(chan os.Signal, 1),
		// 通道 complete 被初始化为无缓冲的通道。当执行任务的 goroutine 完成时，会向这个通道发送一个 error 类型的值或者 nil 值。
		// 之后就会等待 main 函数接收这个值。一旦 main 接收了这个 error 值，goroutine 就可以安全地终止了
		complete: make(chan error),
		// 用 time 包的 After 函数初始化，After 返回一个 time.Time 类型的通道。语言运行时会在指定的 duration 时间到期之后，向这个通道
		// 发送一个 time.Time 的值
		timeout: time.After(d),
	}
}

// Add 将一个任务附加到 Runner 上。这个任务是一个接收一个 int 类型的 ID 作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
// 方法 Start 实现了程序的主流程
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	// 设置了 goInterrupt 方法要从操作系统接收的中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的 goroutine 执行不同的任务
	go func() {
		// run 方法将返回的 error 接口值发送到 complete 通道，一旦 error 接口的值被接收，该 goroutine 就会
		// 通过通道将这个值返回给调用者
		r.complete <- r.run()
	}()

	// 阻塞等待两个事件中的任意一个。
	select {
		// 当任务处理完成时发出的信号
		// 从 complete 通道接收到 error 接口值，那么该 goroutine 要么在规定的时间内完成了分配的工作，
		// 要么收到了操作系统的中断信号。
		case err := <- r.complete:
			return err
		// 当任务处理程序运行超时时发出的信号
		case <- r.timeout:
			return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner) run() error {
	// 迭代 tasks 切片，并按顺序执行每个函数
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		// 在执行之前，检查是否有要从操作系统接收的事件
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}

	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
		// 当中断事件被触发时发出的信号
		// 试图从 interrupt 通道去接收信号。一般来说，select 语句在没有任何要接收的数据时会阻塞
		// 不过，有了 default 分支就不会阻塞了。default 分支会将接收 interrupt 通道的阻塞调用转变为非阻塞的。
		// 如果 interrupt 通道在中断信号需要接收，就会接收这个中断，调用 Stop 方法来停止接收之后的所有事件
		// 之后返回 true
		case <- r.interrupt:
			// 停止接收后续的任何信号
			signal.Stop(r.interrupt)
			return true
		// 继续正常运行
		default:
			return false
	}
}
