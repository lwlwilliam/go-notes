### Timer 和 Ticker：未来的事件

> 原文：[https://yourbasic.org/golang/time-reset-wait-stop-timeout-cancel-interval/](https://yourbasic.org/golang/time-reset-wait-stop-timeout-cancel-interval/)

**Timers 和 Tickers 让你可以在未来的某个时间执行一次或重复执行代码**。

#### Timeout(Timer)

time.After 等待指定的时间然后把当前时间发送到返回的 channel 中：

```
select {
case news := <- AFP:
	fmt.Println(news)
case <- time.After(time.Hour):
	fmt.Println("No news in an hour.")
}
```

底层的 time.Timer 在启动时不会被 GC 回收。如果这是一个问题，那就使用 time.NewTimer 代替，当 timer 不再需要时调用它的 Stop 方法。

```
for alive := true; alive; {
	timer := time.NewTimer(time.Hour)
	select {
	case news := <- AFP:
		timer.Stop()
		fmt.Println(news)
	case <- timer.C:
		alive = false
		fmt.Println("No news in an hour. Service aborting.")
	}
}
```

#### Repeat(Ticker)

time.Tick 返回一个每隔一段时间就提供时钟的 channel：

```
go func() {
	for now := range time.Tick(time.Minute) {
		fmt.Println(now, statusUpdate())
	}
}()
```

底层的 time.Ticker 不会被 GC 回收。如果这是一个问题，就使用 time.NewTicker 来代替，当不再需要 ticker 时，调用它的 Stop 方法。

#### 等待，执行和关闭

time.AfterFunc 会等待指定的时间，然后在它自己的 goroutine 中调用一个函数。它会返回一个可以用来关闭调用的 time.Timer：

```
func Foo() {
	timer = time.AfterFunc(time.Minute, func() {
		log.Println("Foo run for more than a minute.")
	})
	defer timer.Stop()

	// Do heavy work
}
```


