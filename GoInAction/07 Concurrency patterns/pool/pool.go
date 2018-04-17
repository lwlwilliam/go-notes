// Fatih Arslan 和 Gabriel Aszalos 协助完成了这个示例
//  包 pool 管理用户定义的一组资源
package pool

import (
	"errors"
	"log"
	"io"
	"sync"
)

// Pool 管理一组可以安全地在多个 goroutine 间共享的资源
// 被管理的资源必须实现 io.Closer 接口
// 该结构允许调用者根据所需数量创建不同的资源池。
// 只要某类资源实现了 io.Closer 接口，就可以用这个资源池来管理
type Pool struct {
	// 这个互斥锁用来保证在多个 goroutine 访问资源池时，池内的资源是安全的
	m			sync.Mutex
	// 被声明为 io.Closer 接口类型的通道。这个通道是作为一个有缓冲的通道创建的，用来保存共享资源
	// 由于通道的类型是一个接口，所以池可以管理任意实现了 io.Closer 接口的资源类型
	resources	chan io.Closer
	// 函数类型。任何一个没有输入参数且返回一个 io.Closer 和一个 error 接口值的函数，都可以赋值给这个字段。
	// 这个函数的目的是，当池需要一个新资源时，可以用这个函数创建。这个函数的实现细节超出了 pool 包的范围，并且需要由包的使用者实现并提供
	factory		func() (io.Closer, error)
	// 这个字段是一个标志，表示 Pool 是否已经被关闭
	closed		bool
}

// ErrPoolClosed 表示请求(Acquire)了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")

// New 创建一个用来管理资源的池，这个池需要一个可以分配新资源的函数，并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	// 为了保存资源而创建的有缓冲的通道的缓冲区大小
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	// 初始化一个新的 Pool
	return &Pool{
		factory:	fn,
		resources:	make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲的资源
	case r, ok := <- p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
		// 因为没有空闲资源可用，所以提供一个新资源
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 将一个使用后的资源放回池里
// 如果不再需要已经获得的资源，必须将这个资源放回资源池里
func (p *Pool) Release(r io.Closer) {
	// 对互斥量进行加锁，并在函数返回时解锁
	// 保证本操作和 Close 操作的安全
	// 这和 Close 方法中的互斥量是同一个互斥量。这样可以阻止这两个方法在不同 goroutine 里同时运行
	// 使用互斥量有两个目的：第一，可以保护读取 closed 标志的行为，保证同一时刻不会有其他 goroutine 调用
	// Close 方法写同一个标志；第二，不想往一个已经关闭的通道里发送数据，因为那样会引起崩溃。
	// 如果 closed 标志是 true，就知道 resources 通道已经被关闭
	p.m.Lock()
	defer p.m.Unlock()

	// 如果池已经被关闭，销毁这个资源
	// 因为这时已经清空关闭了池，所以无法将资源重新放回到该资源池里。对 closed 标志的读写必须进行同步，
	// 否则可能误导其他 goroutine，让其认为该资源池依旧是打开的，并试图对通道进行无效的操作
	if p.closed {
		r.Close()
		return
	}

	select {
	// 试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	// 如果队列已满，则关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close 会让资源池停止工作，并关闭所有现有的资源，一旦程序不再使用资源池，需要调用这个资源池的 Close 方法
// 需要注意，在同一时刻只能有一个 goroutine 执行这段代码。事实上，当这段代码被执行时，必须保证其他 goroutine 中
// 没有同时执行 Release 方法
func (p *Pool) Close() {
	// 互斥量被加锁，并在函数返回时解锁
	// 保证本操作与 Release 操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	// 如果 pool 已经被关闭，返回并释放锁
	if p.closed {
		return
	}

	// 将池关闭
	p.closed = true

	// 在清空通道里的资源之前，将通道关闭
	// 如果不这样做，会发生死锁
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}