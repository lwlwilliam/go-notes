### Goroutine 是轻量级的线程

> 原文: [https://yourbasic.org/golang/goroutines-explained/](https://yourbasic.org/golang/goroutines-explained/)

**go 语句会让在一个函数在隔离的（另外的）执行线程中运行**。

我们可以用 go 语句创建一个新的执行线程 goroutine。该语句会在一个新建的 goroutine 中运行函数。单独的程序中，所有 goroutine 共享相同的地址空间。

```
go list.Sort()  //  并发运行 list.Sort；不需要等待
```

以下程序会打印"Hello from main goroutine"。也有一定概率会把"Hello from another goroutine"也打印出来，这取决于两个 goroutine 中哪一个先完成。

```
func main() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	// 至此，程序结束，所有执行中的 goroutine 均被终止。
}
```

以下来的程序极有可能会同时打印"Hello from main goroutine"和"Hello from another goroutine"。它们打印的顺序是随机的。然而也有可能第二个 goroutine 执行
非常慢以至无法在程序结束前打印它的信息。

```
func main() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	time.Sleep(time.Second)  // 等待其他 goroutine 执行完成
}
```

以下是一个更贴近实际的示例，我们定义一个并发函数来延缓执行一个事件。

```
// Publish 在给定时间向 stdout 输出文本，它不但不会阻塞面且会立即返回
func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREADING NEWS:", text)
	}()  // 注意这里的`()`，必须调用这个匿名函数
}
```

我们可以如此使用 Publish 函数。

```
func main() {
	Publish("A goroutine starts a new thread.", 5 * time.Second)
	fmt.Println("Lets hope the news will published before I leave.")

	// 等待消息的发送
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I'm leaving now.")
}
```

程序极有可能会以每行间隔 5 秒的速度并按顺序打印以下三行结果。

```
$ go run publish1.go
Let's hope the news will published before I leave.
BREAKING NEWS: A goroutine starts a new thread.
Ten seconds later: I'm leaving now.
```

通常情况下是不可能让线程通过休眠等待其他线程的。Go 主要的同步方法是通过 channels 实现的。

#### 实现

** TODO: 这段死活翻译不出来。。。看来知识储备还差远 **

```
Goroutines are lightweight, costing little more than the allocation of stack space. The stacks start small and grow by allocating and freeing heap storage as required.


Goroutine 是轻量级的，开销比栈空间的分配要小得多。栈刚开始的时候开销比较小，并且随着堆的分配和释放而增长。
```

Goroutine 的内部实现类似于 coroutine，在多个系统线程间复用。如果一个 goroutine 在系统线程中阻塞，例如等待输入，在同一线程中的其他 goroutine 会迁移（到
其他线程），因此它们可以继续运行。

source code: [publish1.go](../src/publish1.go)
