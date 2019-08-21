### 死锁检测

> 原文：[https://yourbasic.org/golang/detect-deadlock/](https://yourbasic.org/golang/detect-deadlock/)

**当一组 goroutine 彼此等待，没有一个可以继续执行时就会发生死锁**。

来看一下这个简单的示例。

```
func main() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<- ch)
}
```

这个程序发送操作时会卡在 channel 中一直等待（其它 goroutine）读取该值。Go 可以检测这种运行时出现的状况。这里是程序的输出：

```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Applications/MAMP/htdocs/notes/web/Golang/Translations/articles/test.go:6 +0x59
exit status 2
```

#### Debug 技巧

goroutine 遇到以下情况可能会卡住：

*	等待 channel；
*	等待 sync 库里的其中一个 locks

通常原因如下：

*	没有其他 goroutine 访问 channel 或者 lock；
*	一组 goroutine 在相互等待，他们中没有一个可以继续运行；

当前 Go 只有当程序卡住时才检测到，而不是 goroutine 中的一个卡住时。

有了 channel ，常常很容易找到导致死锁的原因。另一方面，大量使用互斥的程序很难 debug。
