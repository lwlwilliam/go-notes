### 如何终止 goroutine

> 原文: [https://yourbasic.org/golang/stop-goroutine/](https://yourbasic.org/golang/stop-goroutine/)

**一个 goroutine 不能强制终止另一个**。

为了令 goroutine 可以被停止，需要让它监听 channel 上的停止信号。

```
quit := make(chan struct {})
go func() {
	for {
		select {
		case <- quit:
			return
		default:
			// ...
		}
	}
}()
// ...
close(quit)
```

有时候只用一个 channel 既传递数据也发送信号会很方便。

```
// Generator 返回一个整数类型的 channel，传递 1, 2, 3 等整数
// 为了令以下 goroutine 停止，要关闭 channel
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n ++
			case <- ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<- number)
	fmt.Println(<- number)
	close(number)
	// ...
}
```


source code: [generator.go](../src/generator.go)
