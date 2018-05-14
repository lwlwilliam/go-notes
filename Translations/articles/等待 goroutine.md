### 等待 goroutine

> 原文：[https://yourbasic.org/golang/wait-for-goroutines-waitgroup/](https://yourbasic.org/golang/wait-for-goroutines-waitgroup/)

sync.WaitGroup 用来等待一组 goroutine 的完成。

```
var wg sync.WaitGroup
wg.Add(2)
go func() {
	// Do work.
	wg.Done()
}()
go func() {
	// Do work.
	wg.Done()
}()
wg.Wait()
```

*	首先，主 goroutine 调用 Add 函数设置要等待的 goroutine 数量；
*	然后，两个新的 goroutine 开始执行，执行完调用 Done 函数；

与此同时，Wait 函数是用来一直阻塞直到两个 goroutine 执行完成。

**注意，一个 WaitGroup 绝对不能在第一次使用之后进行复制**。

source code: [waitGoroutine.go](../src/waitGoroutine.go)
