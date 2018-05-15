### Channel 提供同步通信机制

> 原文: [https://yourbasic.org/golang/channels-explained/](https://yourbasic.org/golang/channels-explained/)

**Channel 是 goroutine 之间通过传递值来同步执行以及通信的机制**。

使用内置的 make 函数可以创建一个新的 channel。

```
// 无缓冲 int 类型的 channel
ic := make(chan int)

// 容量为 10 的有缓冲字符串 channel
sc := make(chan string, 10)
```

往 channel 发送值，使用`<-`二元操作符。从 channel 接收值，使用一元操作符`<-`。（译者注：`<-`既是二元操作符，也是一元操作符）

```
ic <- 3		// 往 channel 发送 3
n := <- sc	// 从 channel 接收一个字符串
```

`<-`运算符指定了 channel 的方向，发送或者接收。如果没有指定方向，channel 是双向的。

```
chan Sushi		// 可以用来发送或接收 Sushi 类型的值
chan <- string	// 只能用来发送字符串
<- chan int		// 只能用来接收整数
```

#### 有缓冲和无缓冲的 channel

*	如果 channel 的容量为 0 或者没有指定，它就是无缓冲的，并且发送者会阻塞直到发送值被接收者接收；
*	如果 channel 有缓冲，发送者只有值已经被复制到缓冲时才会阻塞；如果缓冲已经满了，意味着需要一直等待到 channel 中的值被接收者接收；
*	接收者总是阻塞的，直至有数据接收为止；
*	从 nil channel 中发送或者接收操作会永远阻塞（译者注：死锁）；

#### 关闭 channel

close 函数会标记不再有值会发送到 channel 中。注意，只有在接收者在寻找关闭时才有需要关闭 channel。

*	调用 close 以及之前发送的所有数据都被接收了，接收操作将无阻塞地返回一个零值；
*	多值接收操作会额外返回一个表明 channel 是否已关闭的值；
*	往一个已关闭的 channel 中发送值或者试图对该 channel 进行关闭操作会导致运行时 panic；试图关闭一个 nil channel 也会导致运行时 panic；

```
ch := make(chan string)
go func() {
	ch <- "Hello!"
	close(ch)
}()

fmt.Println(<- ch)	// 打印"Hello!"
fmt.Println(<- ch)	// 打印零值""，无阻塞
fmt.Println(<- ch)	// 再次打印""
v, ok := <- ch		// v 为""，ok 为 false

// 一直从 ch 中接收值直至其被关闭
for v := range ch {
	fmt.Println(v)	// 不会被执行
}
```

#### 示例

在以下示例中，让 Publish 函数返回一个 channel，当文本被发布时，该函数会广播消息。

```
// Publish 在给定时间向 stdout 打印文本，当文本被发布时，它会关闭等待的 channel
func Publish(text string, delay time.Duration) (wait <- chan struct {}) {
	ch := make(chan struct {})
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)
	}()
	return ch
}
```

注意，我们使用空结构体类型的 channel 表明该 channel 只用于发送信息，并没有传递数据。可以这样使用以上函数。

```
wait := Publish("important news", 2 * time.Minute)
// 做一些其他事情
<- wait		// 阻塞直至文本被发布
```

source code: [channel1.go](../src/channel1.go)、[channel1.go](../src/channel2.go)
