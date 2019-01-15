> 方法

[methods.go](methods.go)

`Go`没有类。然而，仍然可以在结构体类型上定义方法。

`方法接收者`出现在`func`关键字和方法名之间的参数中。

> 方法(续)

[methods-continued.go](methods-continued.go)

方法仅仅是一个拥有接收者的函数而已。

你可以对包中的`任意`类型定义任意方法，而不仅仅是针对结构体。

只能为同一个包里的类型定义方法，也就是说不能对来自其他包的类型或基础类型定义方法。类型及其对应方法的定义必须在同一个包中。

> 接收者为指针的方法

[methods-with-pointer-receivers.go](methods-with-pointer-receivers.go)

方法可以与命名类型或命名类型的指针关联，也就是说接收者的类型可以为`*T`或`T`，其中`T`不能为指针类型。

刚刚看到两个`Abs`方法，一个是在`*Vertex`指针类型上，而另一个在`MyFloat`值类型上。有两个原因需要使用指针接收者：

1.  首先避免在每个方法调用中拷贝值（如果值类型是大型结构体的话更有效率）；
2.  其次，方法可以修改接收者指向的值。

尝试修改`Abs`的定义，同时`Scale`方法使用`Vertex`代替`*Vertex`作为接收者。

当`v`是`Vertex`的时候`Scale`方法没有任何作用。`Scale`修改`v`。当`v`是一个值（非指针），方法看到的是`Vertex`的副本，
并且无法修改原始值。

`Abs`的工作方式是一样的。只不过，仅仅读取`v`。所以读取的是原始值（通过指针）还是那个值的副本并没有关系。

> 接口

[intefaces.go](interfaces.go)

接口类型是由一组方法签名定义的集合。

接口类型的值可以存放实现这些方法的任何值。

注意：例子代码的 22 行存在一个错误。由于`Abs`只定义在`*Vertex`（指针类型）上，所以`Vertex`（值类型）不满足`Abser`。

接口跟接收者不一样，接收者可以根据调用的值自动转换为值类型或者指针类型，而接口不会自动转换，因此需要注意。

> 隐式接口

[interfaces-are-satisfied-implicitly.go](interfaces-are-satisfied-implicitly.go)

类型通过实现方法来实现接口，只要类型实现了所有的接口方法，就相当于实现了该接口，因此没有显式声明的必要；所以也就没有关键字`implements`。

隐式接口解耦了实现接口的包和定义接口的包：互不依赖。

因此，也就不需要每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

[包io](http://golang.org/pkg/io/)定义了`Reader`和`Writer`；其实不一定要这么做。

> interface 值

interface 值可以认为是一个值和具体类型的元组`(value, type)`。一个 interface 值保存了指定底层类型的值。调用 interface 值的方法时就会执行
该值的底层类型中的同名方法。

[interface-value.go](interface-values.go)

如果 interface 的具体值为 nil，调用方法时就会使用 nil 接收者。在其它语言中这会导致空指针异常，但在 Go 中优雅地处理 nil 接收者调用是很平常的事情。

注意，具体值为 nil 的 interface 值本身并不为 nil。

[interface-values-with-nil.go](interface-values-with-nil.go)

nil interface 值既没有保存值也没有保存具体类型。调用 nil interface 是一个运行时错误，因为 interface 元组中没有类型来说明
要调用哪个具体的方法。

[nil-interface-values.go](nil-interface-values.go)

没有指定方法的 interface 类型即为空 interface：`interface{}`。空 interface 可以保存任何类型的值（因为每种类型都至少实现了 0 个方法）。

空 interface 用来处理未知类型的值，例如，`fmt.Print`可以传入任何数量的`interface{}`类型的参数。

[empty-interface.go](empty-interface.go)

> 类型断言

类型断言提供了访问 interface 值的底层具体值的方式：`t := i.(T)`。

以上语句断言 interface 值 i 保存了具体的类型 T 并把 T 赋值给变量 t。如果 i 没有保存 T，该语句会导致 panic。为了测试 interface 值是否
保存了指定的类型，类型断言可以返回两个值：i 保存的底层值和判断断言是否成功的布尔值：`t, ok := i.(T)`。如果 i 保存了 T，t 就是 i 底层的值，ok 则是 true；
否则，ok 为 false，t 就是类型 T 的零值，并不会导致 panic。

[type-assertions.go](type-assertions.go)

> 类型转换

类型转换是一个允许多个连续类型断言的结构。

类型转换类似普通的 switch 语句，但类型转换中的 case 指定了类型（而不是值），那些值会和给定的 interface 值对应的具体类型进行比较。

```
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no math; here v has the same type as i
}
```

类型转换中的声明和类型断言`i.(T)`中的语法一样，但指定的类型 T 被关键词 type 所替代。以上类型转换语句测试 interface 值 i 是否保存了一个 T 或
S 类型的值。在 case T 和 case S 中，变量 v 分别为类型 T 或 S 并保存有 i 的值。在默认情况下（没有匹配项），变量 v 的类型和值与 i 一致。

[type-switches.go](type-switches.go)

> Stringers

[stringer.go](stringer.go)

最普遍的接口是`fmt`包中定义的`Stringer`。

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

通常函数会返回一个`error`值，调用它的代码时应当通过判断这个错误是否等于`nil`来进行错误处理。

```go
i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't convert number: %v\n", err)
	return
}
fmt.Println("Converted integer: ", i)
```

`error`为`nil`时表示成功；非`nil`的`error`表示错误。

> 练习：错误

[exercise-errors.go](exercise-errors.go)

注意：以上代码中如果在 Error 方法中调用`fmt.Sprint(e)`将会死循环。可以通过先转换 e 的类型来避免：`fmt.Sprint(floag64(e))`。

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

[exercise-reader.go](exercise-reader.go)