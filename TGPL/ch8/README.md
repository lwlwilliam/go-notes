### Goroutines 和 Channels

并发程序指同时进行多个任务的程序，随着硬件的发展，并发程序变得越来越重要。Web 服务器会一次处理成千上万的请求。平板电脑和手机 app 在渲染
用户画面同时还会后台执行各种计算和网络请求。即使是传统的批处理问题——读取数据，计算，写输出——现在也会用并发来隐藏掉 I/O 的操作延迟以充分
利用现代计算机设备的多个核心。计算机的性能每年都在以非纯线性的速度增长。

Go 语言中的并发程序可以用两种手段来实现。本章所讲的 goroutine 和 channel，其支持"顺序通信进程"(communicating sequential processes, 简称为 CSP)。CSP
是一种现代的并发编程模型，在这种编程模型中值会在不同的运行实例(goroutine)中传递，尽管大多数情况下仍然是被限制在单一实例中。

尽管 Go 对并发的支持是众多强力特性之一，但跟踪调试并发程序还是很困难，在线性程序中形成的直觉往往还会使我们误入歧途。

> Goroutines

在 Go 语言中，每一个并发的执行单元叫做一个 goroutine。设想这里的一个程序有两个函数，一个函数做计算，另一个输出结果，假设两个函数没有相
互之间的调用关系。一个线性的程序会先调用其中的一个函数，然后再调用另一个。如果程序中包含多个 goroutine，对两个函数的调用则可能发生在同
一时刻。

如果使用过操作系统或者其它语言提供的线程，可以简单地把 goroutine 类比作一个线程，当然 goroutine 并不等于线程。

当一个程序启动时，其主函数即在一个单独的 goroutine 中运行，我们叫它 main goroutine。新的 goroutine 会用 go 语句来创建。在语法上，go 语句是
一个普通的函数或方法调用前加上关键字 go。go 语句会使其语句中的函数在一个新创建的 goroutine 中运行。而 go 语句本身会迅速地完成。

```go
f()		// call f(); wait for it to return
go f()	// create a new goroutine that calls f((); don't wait
```

在下面的例子，main goroutine 将计算菲波那契数列的第 45 个元素值。由于计算函数使用低效的递归，所以会运行相当长时间，在此期间我们想让
用户看到一个可见的标识来明程序依然在正常运行，所以来做一个动画的小图标：

[spinner.go](spinner.go)

动画显示了几秒之后，fib(45) 的调用成功地返回，并且打印结果：

```go
Fibonacci(45) = 1134903170
```

然后主函数返回。主函数返回时，所有的 goroutine 都会被直接打断，程序退出。除了从主函数退出或者直接终止程序之外，没有其他的编程方法能够
让一个 goroutine 来打断另一个的执行，但是之后可以看到一种方式来实现这个目的，通过 goroutine 之间的通信来让一个 goroutine 请求其它
的 goroutine，并让被请求的 goroutine 自行结束执行。

这里的两个独立的单元，spinning 和菲波那契的计算。分别在独立的函数中，但两个函数会同时执行。

> 并发的 Clock 服务

网络编程是并发大显身手的一个领域。以下是一个顺序执行的时钟服务器，它会每隔一秒钟将当前时间写到客户端：

[clock1.go](clock1.go)

Listen 函数创建了一个 net.Listener 对象，这个对象会监听一个网络端口上到来的连接，在这个例子中用的是 TCP 的 localhost:8080 端口。listener 对
象的 Accept 方法会直接阻塞，直到一个新的连接被创建，然后会返回一个 net.Conn 对象来表示这个连接。handleConn 函数会处理一个完整的客户端连接。

为了连接例子里的服务器，需要一个客户端程序，比如 netcat 这个工具（nc 命令），这个工具可以用来执行网络连接操作。

```
$ go build clock1
$ ./clock1 &
$ nc localhost 8080
09:36:00
09:36:01
09:36:02
09:36:03
^C
```

如果系统没有装 nc 工具，可以用 telnet 来实现同样的效果，或者可以用以下用 go 写的简单的 telnet 程序，用 net.Dial 就可以简单地创建一个 TCP 
连接：

[netcat1.go](netcat1.go)

这个程序会从连接中读取数据，并将读到的内容写到标准输出中，直到遇到 end of file 的条件或者发生错误。不过这里的服务器程序同一时间只能处理一个
客户端连接。所以需要做一点小改动，使其支持并发。

[clock.go](clock.go)

> 并发的 Echo 服务

clock 服务器每一个连接都会起一个 goroutine。本节中会创建一个 echo 服务器，这个服务在每个连接中会有多个 goroutine。大多数 echo 服务仅仅会返
回读取到的内容，如下：

```go
func handleConn(c net.Conn) {
	io.Copy(c, c)  // NOTE: ignoringn errors
	c.Close()
}
```

一个更有意思的 echo 服务应该模拟一个实际的 echo 的"回响"，并且一开始要用大写 HELLO 来表示"声音很大"，之后经过一小段延迟返回一个有所缓和的 
Hello，然后一个全小写字母的 hello 表示声音渐渐变小直至消失，如下：

[reverb1.go](reverb1.go)


