# 基于共享变量的并发

上一章使用了 goroutine 和 channel 这样直接而自然的方式来实现并发。然而这样做实际上回避了在写并发代码时必须处理的一些重要而且细微的问题。

本章中，我们会细致地了解并发机制。尤其是在多 goroutine 之间的共享变量，并发问题的分析手段，以及解决这些问题的基本模式。最后还有 goroutine
和操作系统线程之间的技术上的一些区别。

### 竞争条件

在一个线性（就是说只有一个 goroutine）的程序中，程序的执行顺序只由程序的逻辑来决定。在有两个或更多 goroutine 的程序中，每一个 goroutine
内的语句也是按照既定的顺序去执行的，但是一般情况下我们没法去知道分别位于两个 goroutine 的事件 x 和 y 的执行顺序，x 是在 y 之前还是之后还
是同时发生是没法判断的。当我们没有办法自信地确认一个事件是在另一个事件前面或者后面发生的话，就说明 x 和 y 这两个事件是并发的。

考虑一下，一个函数在线性程序中可以正确地工作。如果在并发的情况下，这个函数依然可以正确地工作的话，那么我们就说这个函数是并发安全的。可以把
这个概念概括为一个特定类型的一些方法和操作函数，如果这个类型是迸发安全的话，那么它的所有访问方法和操作都是并发安全的。

在一个程序中有非并发安全的类型的情况下，依然可以使这个程序并发安全。只有当文档中明确地说明了该类型是并发安全的情况下，才可以并发地去访问它。
我们会避免并发访问大多数的类型，无论是将变量局限在单一的一个 goroutine 内不是用互斥条件维持更高级别的不变性都是为了这个目的。

导出包级别的函数一般情况下都是并发安全的。由于包级别的变量没法被限制在单一的 goroutine，所以修改这些变量"必须"使用互斥条件。

一个函数在并发调用时没法工作的原因太多了，比如死锁(deadlock)、活锁(livelock)和饿死(resource starvation)。但这里只聚焦在竞争条件上。

`竞争条件`指的是程序在多个 goroutine 交叉执行操作时，没有给出正确的结果。竞争条件是很恶劣的一种场景，因为这种问题会一起潜伏在程序里，然后
在非常海风的时候蹦出来，或许只是会在很大的负载时才会发生，又得会在使用了某一种编译器、某一种平台或者某一种架构的时候才会出现。这些使得竞争
条件带来的问题非常难以复现而且难以分析诊断。

传统上经常用经济损失来为竞争条件做比喻，看一个简单的银行账户程序[bank.go](./cmd/bank.go)。当并发地调用存款和查余额的函数时，程序无法保证
结果正确。这个程序包含了一个特定的竞争条件，叫做数据竞争。无论什么时候，只要有两个 goroutine 并发访问同一变量，且至少其中的一个是写操作的时候
就会发生`数据竞争`。

```go
var x []int
go func() { x = make([]int, 10) }()
go func() { x = make([]int, 1000000) }()
x[999999] = 1 // NOTE: undefined behavior; memory corruption possible!
```

最后一个语句中的 x 值是未定义的，有可能是 nil，或者是长度为 10 的 slice，也可能是长度为 1000000 的 slice。

一个好的经验法则是根本就没有什么所谓的良性数据竞争，所以一定要避免数据竞争，那么在程序中要如何做到呢？根据数据竞争的定义，有三种方式可以避免
数据竞争：

1.  不要去写变量，不过这种方法对需要更新的数据来说并不适用；
2.  避免从多个 goroutine 访问变量，"不要使用共享数据来通信；使用通信来共享数据"[bank1.go](./cmd/bank1.go)；
3.  允许很多 goroutine 去访问变量，但是在同一时刻最多只有一个 goroutine 在访问。这种方式被称为"互斥"；

### sync.Mutex 互斥锁

以下用一个容量只有 1 的 channel 来保证最多只有一个 goroutine 在同一时刻访问一个共享变量。一个只能为 1 和 0 的信号量叫做二元信号量
(binary semaphore)，[bank2.go](./cmd/bank2.go)。

这种互斥很实用，而且被 sync 包里的 Mutex 类型直接支持。它的 Lock 方法能够获取到 token（这里叫锁），并且 Unlock 方法会释放这个 token。

[bank3.go](./cmd/bank3.go)

每个 goroutine 访问 blance 变量时，都会调用 mutex 的 Lock 方法来获取一个互斥锁。如果其它的 goroutine 已经获得了这个锁的话，这个操作
会被阻塞直到其它 goroutine 调用了 Unlock 使该锁变回可用状态。mutex 会保护共享变量。惯例来说，被 mutex 所保护的变量是在 mutex 变量声明
之后立刻声明的。

在 Lock 和 Unlock 之间的代码段中的内容 goroutine 可以随便读取或者修改，这个代码段叫做`临界区`。goroutine 在结束后释放锁是必要的。

由于在存款和查询余额函数中的临界区代码这么短——只有一行，没有分支调用——在代码最后去调用 Unlock 就显得更为直截了当。在更复杂的临界区的应用中，
尤其是必须要尽早处理错误并返回的情况下，就很难靠人去判断对 Lock 和 Unlock 的调用是在所有路径中都能够严格配对的了。Go 语言里的 defer 简直
就是这种情况下的救星：我们用 defer 来调用 Unlock，临界区会隐式地延伸到函数作用域的最后。

```go
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
```

deferred Unlock 即使在临界区发生 panic 时依然会执行，这对于用 recover 来恢复的程序来说是很重要的。

考虑一下下面的代码。成功的时候，会正确地减掉余额并返回 true。但如果银行记录资金对交易来说不足，那么取款就会恢复余额，并返回 false。

```go
// NOTE: not atomic!
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false // insufficient funds
	}
	return true
}
```

但取款不是一个原子操作：它包含三个步骤，每一步都需要去获取并释放互斥锁，但任何一次锁都不会锁上整个取款流程。理想情况下，取款应该只在整个操作
中获得一次互斥锁。下面的尝试是错误的：

```go
// NOTE: incorrect!
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false // insufficient funds
	}
	return true
}
```

上例中，Deposit 会调用 mu.Lock() 第二次去获取互斥锁，但因为 mutex 已经锁上了，而无法被重入（go 里没有重入锁）——也就是说没法对一个已经
锁上的 mutex 来再次上锁，这会导致程序死锁，没法继续执行下去，Withdraw 会永远阻塞下去。

一个能用的解决方案是将一个函数分离为多个函数，比如把 Deposit 分离成两个：一个不导出的函数 deposit，这个函数假设锁总是会被保持并去做实际的
操作，另一个是导出的函数 Deposit，这函数会调用 deposit，但在调用前会先去获取锁。同理，也可将 Withdraw 也表示也这种形式：[bank4.go](./cmd/bank4.go)。

### sync.RWMutex 读写锁

上一节中的 Balance 函数只需要读变量的状态，所以同时让多个 Balance 调用并发运行事实上是安全的，只要在运行的时候没有存款或者取款操作就行。
这种场景下我们需要一种特殊类型的锁，其允许多个只读操作并行执行，但写操作会完全互斥。这种锁叫做"多读单写"锁(multiple readers, single writer lock)，
Go 语言提供的锁是 sync.RWMutex。

如[bank5.go](./cmd/bank5.go)中 Balance 函数现在调用了 RLock 和 RUnlock 方法来获取和释放一个读取或者共享锁。

RLock 只能在临界区共享变量没有任何写入操作时可用。一般来说，不应该假设逻辑上的只读函数/方法也不会去更新某一些变量。

RWMutex 只有当获得锁的大部分 goroutine 都是读操作，而锁在竞争条件下，也就是说，goroutine 必须等待才能获取到锁的时候，RWMutex 才是最能
带来好处的。RWMutex 需要更复杂的内部记录，所以会让它比一般的无竞争锁的 mutex 慢一些。
