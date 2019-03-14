/*
在前面的例子中，用互斥锁进行了明确的锁定来让共享的 state 跨多个 Go 协程同步访问。另一个选择是使用内置的 Go 协程
和通道的同步特性来达到同样的效果。这个基于通道的方法和 Go 通过通信以及每个 Go 协程间通过通讯来共享内存，确保每块
数据有单独的 Go 协程，所有的思路是一致的。

 */
package main

import (
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

// 在这个例子中，state 将被一个单独的 Go 协程拥有。这就能保证数据在并行读取时不会混乱。为了对 state 进行读取或者写入，
// 然后接收对应的回复。结构体 readOp 和 writeOp 封装这些请求，并且是拥有 Go 协程响应的一个方式
type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	// 分别计算执行读写操作的次数
	var readOps uint64 = 0
	var writeOps uint64 = 0

	var readTest uint64 = 0
	var writeTest uint64 = 0

	// reads 和 writes 通道分别将被其他 Go 协程用来发布读和写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 这个就是拥有 state 的 Go 协程，反复响应到达的请求。先响应到达的请求，然后返回一个值到响应通道 resp 表示操作成功
	go func() {
		var state = make(map[int]int)
		var i, j int
		for {
			select {
			// 获取读请求中的指定的 state 中对应 readOp.key 的值，发送回 readOp 结构体中的 resp 通道中
			case read := <- reads:
				i += 1
				fmt.Printf("i = %d; j = %d\n", i, j)
				read.resp <- state[read.key]
				// 获取写请求中指定的 key 和 val，并写入 state 中，成功之后往 writeOp 结构体中的 resp 通道发送 true
			case write := <- writes:
				j += 1
				fmt.Printf("i = %d; j = %d\n", i, j)
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 并行发布 100 个读请求
	for r := 0; r < 1; r ++ {
		go func() {
			for {
				read := &readOp {
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<- read.resp

				readTest += 1
				atomic.AddUint64(&readOps, 1)
				//time.Sleep(time.Millisecond)
			}
		}()
	}

	// 并行发布 100 个写请求
	for w := 0; w < 1; w ++ {
		go func() {
			for {
				write := &writeOp {
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<- write.resp

				writeTest += 1
				atomic.AddUint64(&writeOps, 1)
				//time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)


	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readText:", readTest)
	fmt.Println("readOps :", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeTest:", writeTest)
	fmt.Println("writeOps :", writeOpsFinal)
}
