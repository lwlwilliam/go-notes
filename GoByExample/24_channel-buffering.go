/*
默认通道是无缓冲的，这意味着只有在对应的接收(<- chan)通道准备好接收时，才允许进行发送(chan <-)。
可缓冲通道允许在没有对应接收方的情况下，缓存限定数量的值
 */
package main

import "fmt"

func main() {
	// 创建一个通道，最多允许缓冲 2 个值
	messages := make(chan string, 2)

	// 这个通道是有缓冲区的，即使没有一个对应的并发接收方，仍然可以发送这些值
	messages <- "buffered"
	messages <- "channel"
	//messages <- "test"  // 超出缓冲数量会死锁

	fmt.Println(<- messages)
	fmt.Println(<- messages)
}
