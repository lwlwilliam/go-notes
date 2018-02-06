### 结构体和接口

> 结构体定义

Go 提供的结构体就是把使用各种数据类型定义的不同变量组合起来的高级数据类型。

```go
type Rect struct {
	width float64
	length float64
}
```

上面定义了一个矩形结构体，首先是关键：type 表示要定义一个新的数据类型，Rect 是数据类型名称，最后是 struct 表示这个
高级数据类型是结构体。

由于 width 和 length 的数据类型相同，还可以写成如下格式：

```go
type Rect struct {
	width, length float64
}
```

[Rect.go](Rect.go)

结构体可以通过`.`来访问内部成员。包括给内部成员赋值和读取内部成员值。


> 结构体参数传递方式

Go 函数的参数传递方式是值传递。

[Params.go](Params.go)

可以看出，虽然在 double_area 函数里将结构体的宽度和长度都加倍，但仍没影响 main 函数里的 rect 变量。


> 结构体组合函数

上面我们在 main 函数中计算了矩形的面积，但是矩形的面积如果作为矩形结构体的“内部函数”提供会更好。这样就可以直接知道
这个矩形面积是多少，而不用另外去取宽度和长度去计算。

[StructFunction.go](StructFunction.go)

以上代码看起来不太像定义了“内部方法”，因为根本没有定义在 Rect 数据类型的内部。这点跟 C 语方有很大不同。Go 使用组合
函数的方式来为结构体定义结构体方法。下面来看看方法的定义。

首先关键字 func 表示这是一个函数，第二个参数是结构体类型和实例变量，第三个是函数名称，第四个是函数返回值。可以看出，
结构体组合函数跟普通函数的区别就是结构体组合函数多了一个结构体类型限定。这样一来 Go 就知道了这是一个为结构体定义的方法。

注意，定义在结构体上面的函数(function)一般叫做方法(method)。


> 结构体和指针

[StructPointer.go](StructPointer.go)


> 结构体内嵌类型

我们也可以在结构体内部定义另外一个结构体类型的成员。

[StructEmbedStruct.go](StructEmbedStruct.go)


> 接口

[NoInterface.go](NoInterface.go)

以上代码并没有使用接口，NokiaPhone 和 IPhone 都有各自的方法 call()。但是所有手机都应该有 call()。

接口定义：关键字 type，然后是接口名称，最后是关键字 interface 表示这个类型是接口类型。在接口类型里面，我们定义了一组方法。

Go 语言提供了一种接口功能，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口，不一定非要
显式地声明要去实现哪些接口。

[Interface.go](Interface.go)

在 Go 语言，一个类型 A 只要实现了接口 X 所定义的全部方法，那么 A 类型的变量也是 X 类型的变量。
