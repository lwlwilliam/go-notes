### 在 channel 广播信号

> 原文：[https://yourbasic.org/golang/broadcast-channel/](https://yourbasic.org/golang/broadcast-channel/)

**当关闭 channel 时，所有读取者都会接收到零值**。

在该例中，Publish 函数返回一个 channel，该 channel 用来在发布完消息后广播信号。

```
// 在给定时间打印文本，完成之后，等待中的 channel 会关闭
func Publish(text string, delay time.Duration) (wait <- chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch)
	}()
	return ch
}
```

注意，我们使用了一个空结构体类型的 channel `struct{}`。这清晰地表示该 channel 只用来传递信号，不会用来传递数据。
我们可能会这样使用该函数。

```
func main() {
	wait := Publish("Channels let goroutines communicate.", 5 * time.Second)
	fmt.Println("Waiting for news.")
	<- wait
	fmt.Println("Time to leave.")
}
```

```
Waiting for news...
BREAKING NEWS: Channels let goroutines communicate.
Time to leave.
```

source code: [signalOnChannel.go](../src/signalOnChannel.go)
