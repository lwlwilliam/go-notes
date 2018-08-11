// TODO: 未完待续
// 漏桶算法
// 思考一下以下的 client-server 配置：客户端 goroutine 执行一个无限循环来从某些源接收数据，可能是网络。数据被读进缓冲类型 Buffer 中。
// 为了避免太多的分配及释放缓冲，会保存一个空闲缓冲队列，并使用缓冲 channel 来表示`var freeList = make(chan *Buffer, 100)`。
// 这个可重用的缓冲队列被服务器共享。当接收数据时，client 尝试从缓冲 channel 中获取缓冲；但如果 channel 是空的，就会重新分配一个缓
// 冲。一旦消息缓冲加载完，就会从 server channel 中发送到服务器：`var serverChan = make(chan *Buffer)`。
package main

import (
	"fmt"
	"bytes"
	"time"
)

var freeList = make(chan *bytes.Buffer, 100)
var serverChan = make(chan *bytes.Buffer)

// client 算法
func client() {
	for {
		var b *bytes.Buffer
		// 如果可以获取缓冲就获取，否则重新分配
		select {
		case b = <- freeList:
			// Got one; nothing more to do
			fmt.Println("client b", b)
		default:
			// None free, so allocate a new one
			b = new(bytes.Buffer)
		}
		loadInto(b)		// 从网络中读取下一条消息
		serverChan <- b
	}
}

func loadInto(b *bytes.Buffer) {
	b.Write([]byte{'a'})
}

func process(b *bytes.Buffer) {
//	fmt.Println(*b)
}

func server() {
	for {
		b := <- serverChan
		process(b)

		select {
		case freeList <- b:
			// Reuse buffer is free slot on freeList; nothing more to do
		default:
			// Free list full, just carry on: the buffer is 'dropped'
		}
	}
}

func main() {
	go server()
	go client()

	time.Sleep(1010 * time.Nanosecond)
}
