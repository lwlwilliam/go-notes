# 接口

接口类型是对其它类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我
们可以让函数更加灵活和更具有适应能力。

很多面向对象的语言都有相似的接口概念，但 Go 语言中接口类型的独特之处在于它是满足隐式实现的。也就是说，
我们没有必要对于给定的具体类型定义所有满足的接口类型；简单拥有一些必需的方法就足够了。这种设计可以让你
创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义。

### 接口约定

接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；
它只会展示出自己的方法。你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。

得益于使用接口，fmt.Printf 和 fmt.Sprintf 都使用了另一个函数 fmt.Fprintf 来进行封装。
fmt.Fprintf 这个函数对它的计算结果会被怎么使用是完全不知道的。

```go
package fmt

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
func Printf(format string, args ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, args...)
}
func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.String()
}
```

即使 Fprintf 函数中的第一个参数也不是文件类型。它是 io.Writer 类型，这是一个接口类型。

```go
package io

// Writer is the interface that wraps the basic Write method.
type Writer interface {
	// Write writes len(p) bytes from p to the underlying data stream.
	// It returns the number of bytes written from p (p <= n <= len(p))
	// and any error encountered that caused the write to stop early.
	// Write must return a non-nil error if it returns n < len(p).
	// Write must not modify the slice data, even temporarily.
	// 
	// Implementations must not retain p.
	Write(p []byte) (n int, err error)
}
```

io.Writer 类型定义了函数 Fprintf 和这个函数调用者之间的约定。一方面这个约定需要调用者提供具体类型
的值就像`*os.File`和`*bytes.Buffer`，这些类型都有一个特定签名和行为的 Write 的函数。另一方面这
个约定保证了 Fprintf 接受任何满足 io.Writer 接口的值都可以工作。Fprintf 函数可能没有假定写入的是
一个文件或是一段内存，而是写入一个可以调用 Write 函数的值。

因为 fmt.Fprintf 函数没有对具体操作的值作任何假设而是仅仅通过 io.Writer 接口的约定来保证行为，所
以第一个参数可以安全地传入一个任何具体类型的值，只需要满足 io.Writer 接口。一个类型可以自由地使用另
一个满足相同接口的类型来进行替换被称作`可替换性（LSP 里氏替换）`。这是一个面向对象的特征。

以下代码实现了 io.Writer 接口：[bytecounter.go](./cmd/bytecounter.go)

除了 io.Writer 接口，还有另一个对 fmt 包很重要的接口类型。Fprintf 和 Fprintln 函数向类型提供了
一种控制它们输出的途径。

```go
package fmt

// The String method is used to print values passed
// as an operand to any format that accepts a string
// or to an unformatted printer such as Print.
type Stringer interface {
	String() string
}
```

### 接口类型

接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。

io.Writer 类型是用得最广泛的接口之一，因为它提供了所有的类型写入 bytes 的抽象，包括文件类型，内存
缓冲区，网络连接，HTTP 客户端，压缩工个，哈希等等；Reader 可以代表任意可以读取 bytes 的类型；Closer
可以是任意可以关闭的值，例如一个文件或是网络连接。

```go
package io
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
```

有些新的接口类型通过组合已经有的接口来定义。

```go
type ReadWriter interface {
	Reader
	Writer
}
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

上面用到的语法和结构内嵌相似，可以用这种方式以一个简写命名另一个接口，而不是声明它所有的方法。这种
方式称为`接口内嵌`。

### 实现接口的条件

一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。接口指定的规则非常简单：表达一个
类型属于某个接口，只要这个类型实现这个接口。所以：

```go
var w io.Writer
w = os.Stdout           // OK: *os.File has Write method
w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
w = time.Second         // compile error: time.Duration lacks Write method

var rwc io.ReadWriteCloser
rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
```

对于每一个命名过的具体类型 T：它一些方法的接收者是类型 T 本身，然后另一些则是一个`*T`的指针。在 T
类型的参数上调用一个`*T`的方法是合法的，只要这个参数是一个变量；编译器隐式地获取了它的地址。但这仅
仅是一个语法糖：T 类型的值不拥有所有`*T`指针的方法，那这样它就可能只实现更少的接口。

如下例，IntSet 类型的 String 方法的接收者是一个指针类型，所以不能在一个不能寻址的 IntSet 值上
调用这个方法：

```go
type IntSet struct { /* ... */ }
func (*IntSet) String() string
var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
```

但是可以在一个 IntSet 值上调用这个方法：

```go
var s IntSet
var _ = s.String() // OK: s is a variable and &s has a String method
```

然而，由于只有`*IntSet`类型有 String 方法，所以也只有`*IntSet`类型实现了 fmt.Stringer 接口：

```go
var _ fmt.Stringer = &s // OK
var _ fmt.Stringer = s // compile error: IntSet lacks String method
```


