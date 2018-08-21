### 常见的陷阱和错误

*	永远不要使用`var p*a`的形式声明变量，这会混淆指针声明和乘法运算；
*	永远不要在 for 循环中改变计数器变量；
*	永远不要在 for-range 循环中使用一个值去改变自身的值；
*	永远不要将 goto 和前置标签一起使用；
*	永远不要忘记在函数名后加括号，尤其调用一个对象的方法或者使用匿名函数启动一个协程时；
*	永远不要使用 new 创建 map，而要使用 make；
*	当为一个类型定义 String 方法时，不要使用`fmt.Print`或者类似的代码；
*	永远不要忘记当终止缓存写入时，使用 Flush 函数；
*	永远不要忽略错误提示，忽略错误会导致程序崩溃；
*	不要使用全局变量或者共享内存，这会使并发执行的代码变量不安全；
*	println 函数仅用于调试目的；

**最佳实践**：

*	使用正确的方式初始化一个键为 slice 的 map，如 map[string]slice；
*	使用`comma, ok`的形式作为类型断言；
*	使用工厂函数创建并初始化自定义类型；
*	仅当一个结构体的方法需要改变结构体时，使用结构体指针作为方法的接受者，否则使用一个结构体值类型；

#### 误用短声明导致变量覆盖

```
var remember bool = false
if something {
	remember := true
}
// use remember
```

以上代码中，if 结构外的 remember 变量永远不会变为 true。如果 something 为 true，在 if 结构内会声明一个新的 remember 并覆盖外部声明的 remember，
并且它的值为 true。在 if 结构执行结束后，remember 会获取外部的值 false。因此要这样写：

```
if something {
	remember = true
}
```

这种情况也会在 for 循环中出现，尤其在函数中返回全名变量时难以察觉：

```
func shadow() (err error) {
	x, err := check1()  // x 被创建，err 被赋值
	if err != nil {
		return  // err 正确返回
	}
	if y, err := check2(); err != nil {  // y 和 内部的 err 被创建
		return  // 内部错误覆盖了外部错误，因此 nil 被错误地返回
	} else {
		fmt.Println(y)
	}
	return
}
```

#### 误用字符串

当需要对字符串频繁进行操作时，谨记字符串在 Go 中是不可变的（类似 Java 和 C#）。使用诸如`a += b`形式连接字符串效率低下，尤其在一个循环内部使用
这种形式。这会导致大量的内存开销和拷贝。**应该使用一个字符数组代替字符串，将字符串内容写入一个缓存中**。如下例：

```
var b bytes.Buffer
...
for condition {
	b.WriteString(str)  // appends string str to the buffer
}
return b.String()
```

**注意：由于编译优化以及依赖于使用缓存操作的字符串大小，当循环次数大于 15 时，效率才会更佳。**

#### 发生错误时使用 defer 关闭一个文件

当在一个 for 循环内部处理一系列文件，需要使用 defer 确保文件在处理完毕后被关闭，如：

```
for _, file := range files {
	if f, err = os.Open(file); err != nil {
		return
	}
	// 这是错误的方式，当循环结束时文件没有关闭
	defer f.Close()
	// 对文件进行操作
	f.Process(data)
}
```

但是在循环结尾处的 defer 没有执行，所以文件一起没有关闭！垃圾回收机制可能会自动关闭文件，但是这会产生一个错误，更好的做法是：

```
for _, file := range files {
	if f, err = os.Open(file); err != nil {
		return
	}
	// 对文件进行操作
	f.Process(data)
	// 关闭文件
	f.Close()
}
```

**defer 仅在函数返回时才会执行，在循环的结尾或其他一些有限范围的代码内不会执行。**

#### 不需要将一个指向切片的指针传递给函数

切片实际是一个指向数组的指针。常常需要把切片作为一个参数传递给函数是因为：实际就是传递一个指向变量的指针，在函数内可以改变这个变量，而不是传递
数据的拷贝。

因此应该这样做：

```
func findBiggest (listOfNumbers []int) int {}
```

而不是：

```
func findBiggest (listOfNumbers *[]int) int {}
```

当切片作为参数传递时，切记不要解引用切片。

#### 使用指针指向接口类型

查看如下程序：nexter 是一个接口类型，并且定义了一个 next() 方法读取下一节。函数 nextFew 将 nexter 接口作为参数并读取接下来的 num 个字节，并
返回一个切片，这是正确的做法。但是 nextFew2 使用一个指向 nexter 接口类型的指针作为参数传递给函数：当使用 next() 函数时，系统会给出一个编译
错误：`*n.next undefined(type nexter has no field or method next)。

```
package main
import (
	"fmt"
)

type nexter interface {
	next() byte
}

func nextFew1(n nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i ++ {
		b[i] = n.next()
	}
	return b
}

func nextFew2(n *nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i ++ {
		b[i] = n.next()  // 编译错误：n.next 未定义（*nexter 类型没有 next 成员或 next 方法）
	}
	return b
}

func main() {
	fmt.Println("Hello world!")
}
```

**永远不要使用指针指向 interface 类型，因为它本身已经是一个指针。**

#### 使用值类型时误用指针

将一个值作为参数传递给函数或者作为一个方法的接收者，似乎是对内存的滥用，因为值类型一起是传递拷贝。但是另一方面，值类型的内存是在栈上分配，内
存分配快速且开销不大。如果传递一个指针，而不是一个值类型，go 编译器大多数情况下会认为需要创建一个对象，并将对象移动到堆上，所以会导致额外的
内存分配：因此当使用指针代替值类型作为参数传递时，没有任何收获。

#### 闭包和协程的使用

[closures_goroutines.go](./code/closures_goroutines.go)

以上代码中，输出结果：

```
0 1 2 3 4
4 4 4 4 4
1 0 3 4 2
10 11 12 13 14
```

版本 A 调用闭包 5 次打印每个索引值，版本 B 也做相同的事，但是通过协程调用每个闭包。按理说这将执行得更快，因为闭包是并发执行的。如果阻塞足够
多的时间，让所有协程执行完毕，版本 B 的输出是：4 4 4 4 4。为什么会这样？在版本 B 的循环中，ix 变量实际是一个单变量，表示每个数组元素的索引
值。因为这些闭包都只绑定到一个变量，这是一个比较好的方式，当运行这段代码时，将看见每次循环都打印最后一个索引值 4，而不是每个元素的索引值。因
为协程可能在循环结束后还没有开始执行，而此时 ix 值是 4。

版本 C 的循环写法才是正确的：调用每个闭包时将 ix 作为参数传递给闭包。ix 在每次循环时都被重新赋值，并将每个协程的 ix 放置在栈中，所以当协程
最终被执行时，每个索引值对协程都是可用的。输出的顺序取决于每个协程何时被执行。

在版本 D 中，变量声明是在循环体内部，所以在每次循环时，这些变量相互之间是不共享的，所以这些变量可以单独的被每个闭包使用。

#### 糟糕的错误处理

##### 不要使用布尔值

创建一个布尔类型的变量测试错误条件是多余的：

```
var good bool
// test for an error, good becoms true or false
if ! good {
	return errors.New("things aren't good")
}
```

应该立即检测一个错误：

```
.. err1 := api.Func1()
if err1 != nil { ... }
```

##### 不要让错误检测令代码变得混乱

避免写这样的代码：

```
... err1 := api.Func1()
if err1 != nil {
	fmt.Println("err: " + err.Error())
	return
}

err2 := api.Func2()
if err2 != nil {
	...
	return
}
```

首先要在 if 语句包含对函数的调用。

但即使这样，通过 if 语句打印的错误遍布代码。使用这种模式，很难区分哪些是逻辑代码和哪些是错误检测。也能发现大部分代码都用来处理错误了。通常比
较好的解决方法是尽可能地把错误条件封装在闭包中，如下：

```
func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	err := func() error {
		if req.Method != "GET" {
			return errors.New("expected GET")
		}

		if input := parseInput(req); input != "command" {
			return errors.New("malformed command")
		}
		// other error conditions can be tested here
	}()

	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, err)
		return
	}
	doSomething()...
}
```

这种方法可以很容易分辨出错误检测、错误通知和正常的程序逻辑。
