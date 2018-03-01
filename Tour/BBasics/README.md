> 包

[packages.go](packages.go)

每个`Go`程序都是由包组成的。

程序运行的入口是包`main`。

这个程使用并导入了包`fmt`和`math/rand`。

按照惯例，包名与导入路径的最后一个目录一致。例如，`math/rand`包由`package rand`语句开始。

注意：这个程序的运行环境是固定的，因此`rand.Intn`总是会返回相同的数字。（为了得到不同的数字，需要生成不同的种子数，参阅[rand.Seed](http://golang.org/pkg/math/rand/#Seed)。）

> 导入

[imports.go](imports.go)

这个代码用圆括号组合了导入，这是`打包`导入语句。

同样可以编写多个导入语句，例如：

```go
import "fmt"
import "math"
```

不过使用打包的导入语句是更好的形式。

> 导出名

[exported_names.go](exported_names.go)

在导入了一个包之后，就可以用其导出的名称来调用它。

在`Go`中，首字母大写的名称是被导出的。

`Foo`和`FOO`都是被导出的名称。名称`foo`是不会被导出的。

> 函数

[functions.go](functions.go)

函数可以没有参数或接受多个参数.

在这个例子中，`add`接受两个`int`类型的参数。

注意类型在变量名之后。

（参考[这篇关于 Go 语言定义的文章](http://golang.org/doc/articles/gos_declaration_syntax.html)了解类型以这种形式出现的原因。）

> 函数（续）

[functions-continued.go](functions-continued.go)

当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。

在这个例子中，

```go
x int, y int
```

被缩写为

```go
x, y int
```

> 多值返回

[multiple-results.go](multiple-results.go)

函数可以返回任意数量的返回值。

`swap`函数返回了两个字符串。

> 命名返回值

[named-results.go](named-results.go)

`Go`的返回值可以被命名，并且像变量那样使用。

返回值的名称应当具有一定的意义，可以作为文档使用。

没有参数的`return`语句返回结果的当前值。也就是`直接`返回。

直接返回语句仅应当用在像上面这样的短函数中。在长的函数中它们会影响代码的可读性。

> 变量

[variables.go](variables.go)

`var`语句定义了一个变量的列表：跟函数的参数列表一样，类型在后面。

就像在这个例子中看到的一样，`var`语句可以定义在包或函数级别。

> 初始化变量

[variables-with-initializers.go](variables-with-initializers.go)

变量定义可以包含初始值，每个变量对应一个。

如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。

> 短声明变量

[short-variable-declarations.go](short-variable-declarations.go)

在函数中，`:=`简洁赋值语句在明确类型的地方，可以用于替代`var`定义。

函数外的每个语句都必须以关键字开始(`var`、`func`等)，`:=`结构不能使用在函数外。

> 基本类型

[basic-types.go](basic-types.go)

`Go`的基本类型有`Basic types`

```
bool

string

int   int8   int16   int32   int64
uint  uint8  uint16  uint32  uint64  uintptr

byte  // uint8 的别名

rune  // int32 的别名
      // 代表一个 Unicode 码
      
float32  float64

complex64  comples128
```

这个例子演示了具有不同类型的变量。同时与导入语句一样，变量的定义`打包`在一个语法块中。`

> 零值

[zero.go](zero.go)

变量在定义时没有明确的初始化时会赋值为`零值`。

零值是：

```
数值类型为`0`
布尔类型为`false`
字符串为`""`（空字符串）
```

> 类型转换

[type-conversions.go](type-conversions.go)

表达式`T(v)`将值`v`转换为类型`T`。

一些关于数值的转换：

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

或者，更加简单的形式：

```go
i := 42
f := float64(i)
u := uint(f)
```

> 类型推导

[type-inference.go](type-inference.go)

在定义一个变量但不指定其类型时（使用没有类型的`var`或`:=`语句），变量的类型由右值推导得出。

当右值定义了类型时，新变量的类型与其相同：

```go
var i int
j := i  // j 也是一个 int
```

但是当右边包含了未指名类型的数字常量时，新的变量就可能是`int`、`float64`或`complex128`。这取决于常量的精度：

```go
i := 42  // int
f := 3.142  // float64
g := 0.867 + 0.5i  // complex128
```

> 常量

[constants.go](constants.go)

常量的定义与变量类似，只不过使用`const`关键字。

常量可以是字符、字符串、布尔或数字类型的值。

常量不能使用`:=`语法定义。

> 数值常量

[numeric-constants.go](numeric-constants.go)

数值常量是高精度的值。

一个未指定类型的常量由上下文来决定其类型。
