package main

import (
	"fmt"
	"sync"
)

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
