# Goroutines 和 Channels

并发程序指同时进行多个任务的程序，随着硬件的发展，并发程序变得越来越重要。Web 服务器会一次处理成千上万的请求。
平板电脑和手机 app 在渲染用户画面同时还会后台执行各种计算任务和网络请求。即使是传统的批处理问题——读取数据，
计算，写输出——现在也会用并发来隐藏掉 I/O 的操作延迟以充分利用现代计算机设备的多个核心。计算机的性能每年都在
以非线性的速度增长。

Go 语言中的并发程序可以用两种手段来实现。本章的 goroutine 和 channel，其支持“顺序通信进程”(communicating 
sequential processes)或被简称为 CSP。CSP 是一种现代的并发编程模型，在这种编程模型中值会在不同的运行实例
(goroutine)中传递，尽管大多数情况下仍然是被限制在单一实例中。下一章覆盖更为传统的并发模型：多线程共享内存。

尽管 Go 对并发的支持是从多强力特性之一，但跟踪调试并发程序还是很困难，在线性程序中形成的直觉往往还会使我们
误入歧途。

### Goroutines

在 Go 语言中，每一个并发的执行单元叫做一个 goroutine。当一个程序启动时，其主函数即在一个单独的 goroutine
中运行，我们叫它 main goroutine。新的 goroutine 会用 go 语句来创建。在语法上，go 语句是一个普通的函数
或方法调用前加上关键字 go。go 语句会使其语句中的函数在一个新创建的 goroutine 中运行。而 go 语句本身会迅
速地完成。

```go
f()     // call f(); wait for it to return
go f()  // create a new goroutine that calls f(); don't wait
```

使用低效的递归计算菲波那契数列：[spinner.go](./cmd/spinner.go)。

以上代码主函数返回时，所有的 goroutine 都会被直接打断，程序退出。除了从主函数退出或者直接终止程序之外，
没有其它的编程方法能够让一个 goroutine 来打断另一个的执行，但是之后可以看到一种方式来实现这个目的，
通过 goroutine 之间的通信来让一个 goroutine 请求其它的 goroutine，并让被请求的 goroutine 自行结束
执行。

### 并发的 Clock 服务

网络编程是并发大显身手的一个领域，由于服务器是最典型的需要同时处理很多连接的程序，这些连接一般来自于彼此独
立的客户端。

顺序执行的时钟服务器，每隔一秒钟将当前时间写到客户端：[clock1.go](./cmd/clock1.go)。与之配合使用的客户
端：[netcat1.go](./cmd/netcat1.go)。第二个客户端必须等待第一个客户端完成工作，这样服务端才能继续向后
执行，因为这个时钟服务器同一时间只能处理一个客户端连接。

现在对 clock1 进行小小的修改，在调用 handleConn 函数前加上 go 关键字，让每一次 handleConn 的调用都进
入一个独立的 goroutine：[clock2.go](./cmd/clock2.go)。现在多个客户端可以同时接收到时间了。
