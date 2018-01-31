> goroutine

[goroutines.go](goroutines.go)

`goroutine`是由`Go`运行时环境管理的轻量级线程。

```go
go f(x, y, z)
```

开启一个新的`goroutine`执行

```go
f(x, y, z)
```

`f`，`x`，`y`和`z`是当前`goroutine`中定义的，但是在新的`goroutine`中运行`f`。

`goroutine`在相同的地址空间中运行，因此访问共享内存必须进行同步。`sync`提供了这种可能，
不过在`Go`中并不经常用到，因为有其他的办法。

> channel

[channels.go](channels.go)

`channel`是有类型的管道，可以用`channel`操作符`<-`对其发送或者接收值。

```go
ch <- v  // 将 v 送入 channel ch。
v := <-ch  // 从 ch 接收，并且赋值给 v。
```

(“箭头”就是数据流的方向。)

和`map`与`slice`一样，`channel`使用前必须创建：

```go
ch := make(chan int)
```

默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得`goroutine`可以在没有明确的锁或竞态变量的情况下进行同步。

> 缓冲 channel

[buffered-channels.go](buffered-channels.go)

`channel`可以是`带缓冲`的。为`make`提供第二个参数作为缓冲长度来初始化一个缓冲`channel`：

```go
ch := make(chan int, 100)
```

向带缓冲的`channel`发送数据的时候，只有在缓冲区满的时候才会阻塞。而当缓冲区为空的时候接收操作会阻塞。

修改例子使得缓冲区被填满，然后看看会发生什么。

> range 和 close

[range-and-close.go](range-and-close.go)

发送者可以`close`一个`channel`来表示再没有值会被发送了。接收者可以通过赋值语句的第二参数来测试`channel`是
否被关闭，那么经过

```go
v, ok := <-ch
```

之后`ok`会被设置为`false`。

循环`fot i := range c`会不断从`channel`接收值，直到它被关闭。

**注意：** 只有发送者才能关闭`channel`，而不是接收者。向一个已经关闭的`channel`发送
数据会引起`panic`。**还要注意：**`channel`与文件不同；通常情况下无需关闭它们。只有在需要告诉
接收者没有更多的数据的时候才有必要进行关闭，例如中断一个`range`。

> select

[select.go](select.go)

`select`语句使得一个`goroutine`在多个通讯操作上等待。

`select`会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。
当多个都准备好的时候，会随机选择一个。

> 默认选择

[default-selection.go](default-selection.go)

当`select`中的其他条件分支都没有准备好的时候，`default`分支会被执行。

为了非阻塞的发送或者接收，可使用`default`分支：

```go
select {
case i := <-c:
	// 使用 i
default:
    // 从 c 读取会阻塞
}
```

> 练习：等价二叉树

```go
// TODO
```

> sync.Mutex

[mutex-counter.go](mutex-counter.go)

`channel`用来在各个`goroutine`间进行通信是非常合适的。

但是如果我们不需要通信呢？比如说，如果我们只是想保证在每个时刻，只有一个`goroutine`能访问一个共享的变量从而避免冲突？

这里涉及的概念叫做`互斥`，通常使用`互斥锁(mutex)`来提供这个限制。

`Go`标准库中提供了`sync.Mutex`类型及其两个方法：

```go
Lock
Unlock
```

我们可以通过在代码前调用`Lock`方法，在代码后调用`Unlock`方法来保证一段代码的互斥执行。参见`Inc`方法。

也可以用`defer`语句来保证互斥锁一定会被解锁。参见`Value`方法。

> 练习：Web 爬虫

```go
// TODO
```
