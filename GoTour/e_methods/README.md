> 方法

[methods.go](methods.go)

`Go`没有类。然而，仍然可以在结构体类型上定义方法。

`方法接收者`出现在`func`关键字和方法名之间的参数中。

> 方法(续)

[methods-continued.go](methods-continued.go)

你可以对包中的`任意`类型定义任意方法，而不仅仅是针对结构体。

但是，不能对来自其他包的类型或基础类型定义方法。

> 接收者为指针的方法

[methods-with-pointer-receivers.go](methods-with-pointer-receivers.go)

方法可以与命名类型或命名类型的指针关联。

刚刚看到两个`Abs`方法，一个是在`*Vertex`指针类型上，而另一个在`MyFloat`值类型上。有两个原因需要使用指针接收者。
首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话更有效率）。其次，方法可以修改接收者指向的值。

尝试修改`Abs`的定义，同时`Scale`方法使用`Vertex`代替`*Vertex`作为接收者。

当`v`是`Vertex`的时候`Scale`方法没有任何作用。`Scale`修改`v`。当`v`是一个值（非指针），方法看到的是`Vertex`的副本，
并且无法修改原始值。

`Abs`的工作方式是一样的。只不过，仅仅读取`v`。所以读取的是原始值（通过指针）还是那个值的副本并没有关系。

> 接口

[intefaces.go](interfaces.go)

接口类型是由一组方法定义的集合。

接口类型的值可以存放实现这些方法的任何值。

注意：例子代码的 22 行存在一个错误。由于`Abs`只定义在`*Vertex`（指针类型）上，所以`Vertex`（值类型）不满足`Abser`。

> 隐式接口

[interfaces-are-satisfied-implicitly.go](interfaces-are-satisfied-implicitly.go)

类型通过实现那些方法来实现接口。没有显式声明的必要；所以也就没有关键字`implements`。

隐式接口解耦了实现接口的包和定义接口的包：互不依赖。

此，也就无需要每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

[包io](http://golang.org/pkg/io/)定义了`Reader`和`Writer`；其实不一定要这么做。

> Stringers

[stringer.go](stringer.go)

一个普遍存在的接口是`fmt`包中定义的`Stringer`。

```go
type Stringer interface {
	String() string
}
```

`Stringer`是一个可以用字符串描述自己的类型。`fmt`包（还有其他包）使用这个来进行输出。

> 练习：Stringers

[exercise-stringer.go](exercise-stringer.go)

让`IPAddr`类型实现`fmt.Stringer`以便用点分格式输出地址。

例如，`IPAddr{1, 2, 3, 4}`应当输出`1.2.3.4`。

> 错误

[errors.go](errors.go)

`Go`程序使用`error`值来表示错误状态。

与`fmt.Stringer`类似，`error`类型是一个内建接口：

```go
type error interface {
	Error() string
}
```

(与`fmt.Stringer`类似，`fmt`包在输出时也会试图匹配`error`。)

通常函数会返回一个`error`值，调用的它的代码应当判断这个错误是否等于`nil`，来进行错误处理。

```go
i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't convert number: %v\n", err)
	return
}
fmt.Println("Converted integer: ", i)
```

`error`为`nil`时表示成功：非`nil`的`error`表示错误。

> 练习：错误(未完待续：TODO)

[exercise-errors.go](exercise-errors.go)

> Readeres

[reader.go](reader.go)

`io`包指定了`io.Reader`接口，它表示从数据流结尾读取。

`Go`标准库包含了这个接口的许多实现，包括文件、网络连接、压缩、加密等等。

`io.Reader`接口有一个`Read`方法：

```go
func (T) Read(b []byte) (n int, err error)
```

`Read`用数据填充指定的字节`slice`，并且返回填充的字节数和错误信息。在遇到数据流结尾时，返回`io.EOF`错误。

例子代码创建了一个`strings.Reader`。并且以每次 8 字节的速度读取它的输出。
