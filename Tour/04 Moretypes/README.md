> 指针

[pointers.go](pointers.go)

`Go`具有指针。指针保存了变量的内存地址。

类型`*T`是指向类型`T`的值的指针。其零值是`nil`。

```go
var p *int
```

`&`符号会生成一个指向其作用对象的指针。

```go
i := 42
p = &i
```

`*`符号表示指针指向的底层的值。

```go
fmt.Println(*p)  // 通过指针 p 读取 i
*p = 21  // 通过指针 p 设置 i
```

这也就是通常所说的"间接引用"或"非直接引用"。

与`C`不同，`Go`没有指针运算。

> 结构体

[structs.go](structs.go)

一个结构体(`struct`)就是字段的一个集合。（而`type`的含义跟其字面意思相符。）

> 结构体字段

[struct-fields.go](struct-fields.go)

结构体字段使用点号来访问。

> 结构体指针

[struct-pointers.go](struct-pointers.go)

结构体字段可以通过结构体指针来访问。

假设 p 是结构体的指针，可以使用`(*p).X`的方式来访问结构体的字段 X。由于这种方式比较繁琐，所以 Go 语言允许使用`p.X`来代替上述写法，免去了显式的解引用。

> 结构体字面量

[struct-literals.go](struct-literals.go)

结构体字面量通过枚举结构体字段值的方式来分配一个新结构体。

可以通过使用`Name:`语法列出字段的一个子集。（字段名的顺序无关。）

指定的前缀`&`返回一个指向结构体的指针。

> 数组

[array.go](array.go)

类型`[n]T`是有`n`个`T`类型值的数组。

表达式

```go
var a [10]int
```

定义变量`a`是一个有十个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是一个制约，但是请不要担心：`Go`提供了更加便利的方式来使用数组。

> slice

[slices.go](slices.go)

slice 指向一个底层数组，不保存任何数据，它只是对底层数组的一个片段的描述。修改 slice 的元素同时也会修改对应底层数组的元素。一旦对 slice 进行了修改，共享同一
底层数组的其他 slice 也会"看到"。

`[]T`是一个元素类型为`T`的`slice`。

slice 的零值是 nil。nil slice 的长度和容量都是 0，没有对应的底层数组。

> 对 slice 切片

[slicing-slices.go](slicing-slices.go)

`slice`可以重新切片，创建一个新的`slice`值指向相同的数组。

表达式

```go
s[lo:hi]
```

表示从`lo`到`hi-1`的`slice`元素，含两端。因此

```go
s[lo:lo]
```

是空的，而

```go
s[lo:lo+1]
```

有一个元素。

> 构造 slice

[making-slices.go](making-slices.go)

`slice`由函数`make`创建。这会分配一个零长度的数组并返回一个`slice`指向这个数组：

```go
a := make([]int, 5)  // len(a) = 5
```

为了指定容量，可传递第三个参数到`make`：

```go
b := make([]int, 0, 5)  // len(b) = 0, cap(b) = 5

b = b[:cap(b)]  // len(b) = 5, cap(b) = 5
b = b[1:]  // len(b) = 4, cap(b) = 4
```

> nil slice

[nil-slices.go](nil-slices.go)

`slice`的零值是`nil`。

一个`nil`的`slice`的长度和容量是 0。

> 向 slice 添加元素

[append.go](append.go)

向`slice`添加元素是一种常见的操作，因此`Go`提供了一个内建函数`append`。内建函数的[文档](http://golang.org/pkg/builtin/#append)
对`append`有详细介绍。

```go
func append(s []T, vs ...T) []T
```

`append`的第一个参数`s`是一个类型为`T`的数组，其余类型为`T`的值将会添加到`slice`。

`append`的结果是一个包含原`slice`所有元素加上新添加的元素的`slice`。

如果`s`的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。返回的`slice`会指向这个新分配的数组。

这里需要注意的是，当`s`的底层数组太小时，append 执行后，原 slice 并不会改变。

（了解更多关于`slice`的内容，参阅文章[slice: 使用和内幕](http://golang.org/doc/articles/slices_usage_and_internals.html)。）

> range

[range.go](range.go)

`for`循环的`range`格式可以对`slice`或者`map`进行迭代循环。

> range(续)

[range-continued.go](range-continued.go)

可以通过赋值给`_`来忽略序号和值。

如果只需要过索引值，去掉`, value`的部分即可。

> 练习：slice

[exercise-slices.go](exercise-slices.go)

实现`Pic`。它返回一个`slice`的长度`dy`，和`slice`中每个元素的长度的 8 位无符号整数`dx`。当执行这个程序，
它会将整数转为灰度（好吧，蓝度）图片进行展示。

图片的实现已经完成。可能用到的函数包括`(x+y)/1`、`x*y`和`x^y`（使用`math.Pow`计算最后的函数）。

（需要使用循环来分配`[][]uint8`中的每个`[]uint8`。）

（使用`uint8(intValue)`在类型之间进行转换。）

> map

[maps.go](maps.go)

`map`映射键到值。

`map`在使用之前必须用`make`而不是`new`来创建；值为`nil`的`map`是空的，并且不能赋值。

> map 的文法

[map-literals.go](map-literals.go)

`map`的文法跟结构体文法相似，不过必须有键名。

> map 的文法(续)

[map-literals-continued.go](map-literals-continued.go)

如果顶级的类型只有类型名的话，可以在文法的元不中省略键名。

> 修改 map

[mutating-maps.go](mutating-maps.go)

在`map m`中插入或修改一个元素：

```go
m[key] = elem
```

获得元素：

```go
elem = m[key]
```

删除元素：

```go
delete(m, key)
```

通过双赋值检测某个键存在：

```go
elem, ok = m[key]
```

如果`key`在`m`中，`ok`为`true。否则，`ok`为`false`，并且`elem`是`map`的元素类型的零值。

同样的，当从`map`中读取某个不存在的键时，结果是`map`的元素类型的零值。

> 练习：map

[exercise-maps.go](exercise-maps.go)

实现`WordCount`。它应当返回一个含有`s`中每个“词”个数的`map`。函数`wc.Test`针地这个函数执行一个测试用例，
并输出成功还是失败。

你会发现[strings.Fields](http://golang.org/pkg/strings/#Fields)很有帮助。

> 函数值

[function-values.go](function-values.go)

函数也是值。

> 函数的闭包

[function-closures.go](function-closures.go)

`Go`函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。函数可以对这个引用值进行访问和赋值；
换句话说这个函数被“绑定”在这个变量上。

例如，函数`adder`返回一个闭包。每个闭包都被绑定到其各自的`sum`变量上。

> 练习：斐波纳契闭包

[exercise-fibonacci-closure.go](exercise-fibonacci-closure.go)

现在来通过函数做些有趣的事情。

实现一个`fibonacci`函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
