### Select 监控 channel 组

> 原文: [https://yourbasic.org/golang/select-explained/](https://yourbasic.org/golang/select-explained/)

**Select 语句同时监控多个发送或接收操作**。

```
// 阻塞直至 ch1 或 ch2 的数据可用
select {
case <- ch1:
	fmt.Println("Received from ch1")
case <- ch2:
	fmt.Println("Received from ch2")
}
```

*	该语句整体阻塞直到其中一个操作解除封锁；
*	如果多个 case 可以继续进行，会在它们中随机选中一个继续；

在 nil channel 中的发送和接收操作会永远阻塞。这可以被用禁用 select 语句中的一个 channel：

```
ch1 = nil	// 禁用该 channel
select {
case <- ch1:
	fmt.Println("Received from ch1")	// 不会执行
case <- ch2:
	fmt.Println("Received from ch2")
}
```

#### 默认 case

默认的 case 总是可以继续执行的，如果其他 case 阻塞它就会执行。

```
// 永不阻塞
select {
case x := <- ch:
	fmt.Println("Received", x)
default:
	fmt.Println("Nothing available")
}
```

#### 示例

一个无限的随机二进制序列

```
rand := make(chan int)
for {
	select {
	case rand <- 0:	// 没有语句
	case rand <- 1:
	}
}
```

阻塞操作超时

```
select {
case news := <- AFP:
	fmt.Println(news)
case <- time.After(time.Minute)
	fmt.Println("Time out: No news in one minute")
}
```

time.After 函数是标准库的一部分；它可以等待指定的时间然后向返回的 channel 中发送当前时间。

一个永远阻塞的语句

```
select {}
```

select 语句会一起阻塞走到它的其中一个 case 可以继续执行。没有 case 是永远不会发生的。（译者注：这句话翻译得怪怪的）

source code: [select1.go](../src/select1.go)、[select2.go](../src/select2.go)
