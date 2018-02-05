### 使用函数

> 什么是函数

[Definition.go](Definition.go)

函数，简单来讲就是一段将`输入数据`转换为`输出数据`的`公用代码块`。

看代码`Definition.go`，函数的定义：func(关键字) + 函数名称 + 参数列表 + 返回值列表。当然如果函数没有参数或者返回值，
这两项都是可选的。


> 命名返回值

Go 的函数很有趣，甚至可以为返回值预先定义一个名称，在函数结束的时候，直接一个 return 就可以返回所有的预定义返回值。
如`Return.go`。

[Return.go](Return.go)

不过，如果定义了命名返回值，那么在函数内部将不能重复定义一个同名变量。


> 实参数和虚参数

实参数就是函数调用时传入的参数，虚参数就是函数定义时表示函数需要传入哪些参数的占位参数。实参数和虚拟数名字可以一样，也可以不一样。


> 函数多返回值

在 java 或 c 里面需要返回多个值时还得定义一个对象或者结构体。而在 Go 里面，则不需要这么做。Go 函数支持返回多个值。
Go 自带的函数 range 就是多返回值的。

[MultiReturn.go](MultiReturn.go)


> 变长参数

Go 支持可变长参数列表。

[VariableLengthParam.go](VariableLengthParam.go)

在上例中，将原来的切片参数修改为可变长参数，然后使用 range 函数迭代这些参数，并求和。由此看出，可变长参数列表里的参数
类型都是相同的。另外要注意：**可变长参数定义只能是函数的最后一个参数**。


> 闭包函数

[Closure.go](Closure.go)

闭包函数对它外层的函数中的变量具有`访问`和`修改`的权限。


> 递归函数

谈到递归函数，绕不开阶乘和斐波拉契数列。

[Factorial.go](Factorial.go)

[Fibonacci.go](Fibonacci.go)


> 异常处理

[Defer.go](Defer.go)

Go 语言提供了 defer 来在函数运行结束的时候运行一段代码或调用一个清理函数。上例函数虽然 second() 函数写在 
first() 函数前面，但是由于使用了 defer 标注，所以它是在 main() 函数执行结束的时候才调用。

defer 用途最多的在于释放各种资源。比如读取一个文件，读完之后需要释放文件句柄。

[ReadFileWithDefer.go](ReadFileWithDefer.go)

`panic`和`recover`是 Go 语言提供的用以处理异常的关键字，panic 用来触发异常，recover 用来终止异常并且返回传递给 panic 的值。
注意：**recover 并不能处理异常，而且 recover 只能在 defer 里面使用，否则无效。**
