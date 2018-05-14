### 数据竞争解释

> 原文：[https://yourbasic.org/golang/data-races-explained/](https://yourbasic.org/golang/data-races-explained/)

**当两个以上的 goroutine 并发访问同一个变量，且至少有一个访问是写操作时，就会发生数据竞争**。

数据竞争相当普遍而且可能会很难 debug。

以下函数存在数据竞争且它的行为是未定义的。例如，它可能会打印 1。尝试思考它是如何发生的，其中一个可能的原因在以下代码之后。

```
func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n ++	// 读取，递增，读入
		close(wait)
	}()
	n ++	// 访问冲突
	<- wait
	fmt.Println(n)	// 输出: <unspecified>
}
```

两个 goroutine，g1 和 g2 处于竞争状态，没有办法知道会出现哪种操作顺序。以下是众多可能输出结果中的一个。

| g1 | g2 |
| :- | :- |
| 从 n 中读取 0 值 | |
| | 从 n 中读取 0 值 |
| 值从 0 递增到 1 | |
| 把 1 写入 n 中 | |
| | 值从 0 递增到 1 |
| | 把 1 写入 n 中 |
| 打印 n，值为 1 | |

"数据竞争"这个名字会让人误解。不仅仅操作顺序是未定义的，而且几乎是没有保证的。编译器和硬件会经常把代码颠倒，以获取更好的性能。如果你在中间
过程观察一个线程，可能会看到很多东西。（这段真是糟糕）

#### 如何避免数据竞争

避免数据竞争的唯一方法是使用同步方式访问所有在线程间共享的可变数据。有几种方法可以实现这种方式。在 Go 语言中，通常使用 channel 或者 lock。
（底层原理可以在 sync 和 sync/atomic 包中了解）

在 Go 中更提倡使用 channel 在 goroutine 之间传递数据的方式来处理并发数据访问的问题。Go 的箴言："不要通过共享内存来通信；通过通信来共享内存。"

```
func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0	// 局部变量只在一个 goroutine 中可见
		n ++
		ch <- n	// 数据离开 goroutine
	}()
	n := <- ch	// ...（译者注：数据）安全地到达另一个。（译者注：goroutine）
	n ++
	fmt.Println(n)	// 输出：2
}
```

在这段代码中 channel 做了两件事：

*	把数据从一个 goroutine 传递到另一个 goroutine；
*	作为一个同步点；

作为发送方的 goroutine 要等待其他 goroutine 接收数据，而作为接收方的 goroutine 等待其他 goroutine 发送数据。

Go 的内存模型，在一个 goroutine 中，要保证在读取变量时可以监测到在不同的 goroutine 中对同一个变量的写入是相当复杂的，但只要你在 goroutine 
之间共享可变数据时使用 channel，就可以安全地处理数据竞争的问题。

souce code: [dataRace.go](../src/dataRace.go)
