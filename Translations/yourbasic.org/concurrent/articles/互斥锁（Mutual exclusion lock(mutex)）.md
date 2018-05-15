### 互斥锁

> 原文：[https://yourbasic.org/golang/mutex-explained/](https://yourbasic.org/golang/mutex-explained/)

**互斥让你可以通过显式锁定来同步数据访问，而不需要 channel**。

有时候通过显式锁定来同步数据访问会比 channel 来得方便。出于该目的，Go 标准库提供了互斥锁 sync.Mutex。

#### 谨慎使用

为了保证锁定的安全，所有对共享数据的读写访问只有在 goroutine 拥有该锁时才能执行显得尤为重要。单个 goroutine 中的一个错误就足以产生数据竞争
以致程序崩溃。

由此，应该考虑使用干净的 API 设计自定义数据结构确保所有同步操作都在内部完成。

该例中我们创建了一个安全易用的并发数据结构——AtomicInt，用来保存单个整数。任意数量的 goroutine 都可以通过 Add 和 Value 方法安全地访问该数字。

```
// AtomicInt 是一个用来保存一个整数的并发数据结构，它的零值是 0。
type AtomicInt struct {
	mu sync.Mutex	// 一次可以被一个 goroutine 保存一次
	n int
}

// Add 把以单个原子操作方式把整数 n 添加到 AtomicInt
func (a *AtomicInt) Add(n int) {
	a.mu.Lock()		// 等待锁释放后就获取该锁
	a.n += n
	a.mu.Unlock()	// 释放锁
}

// Value 返回 a 的值
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

func main() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1)	// 一个访问
		close(wait)
	}()
	n.Add(1)	// 另一个并发访问
	<- wait
	fmt.Println(n.Value())	// 2
}
```

source code: [mutex.go](../src/mutex.go)
