### 字符串

注意，内置的 len 函数可以返回一个字符串中的字节数目（不是 rune 字符数目），索引操作`s[i]`返回第 i 个字节的字节值，i 必须满足 0 ≤ i < len(s)
条件约束。第 i 个`字节`并不一定是字符串的第 i 个`字符`。因为对于非 ASCII 字符的 UTF8 编码会要两个或多个字节。

字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变，当然也可以给一个字符串变量分配一个新字符串值，如下：

```
s := "left foot"
t := s
s += ", right foot"
```

这并不会导致原始的字符串值被改变，但是变量 s 将因为 += 语句持有一个新的字符串值，但是 t 依然是包含原先的字符串值。

因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的。

`不变性意味如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的。同样，一个字符
串 s 和对应的子字符串切片 s[7:] 的操作也可以安全地共享相同的内存，因此字符串切片操作也是低廉的。在这两种情况下都
没有必要分配新的内存。`

#### Unicode

Unicode 为每个符号都分配一个唯一的 Unicode 码点，Unicode 码点对应 Go 语言中的 rune 整数类型，
rune 是 int32 等价类型。

可以将一个符文序列表示为一个 int32 序列，这种编码方式叫 UTF-32 或 UCS-4，每个 Unicode 码点都使用同样
大小的 32big 来表示。这种方式比较简单统一，但是它会浪费很多存储空间，因为大多数计算机可读的文本是 ASCII
字符，本来每个 ASCII 字符只需要 8bit 或 1 字节就能表示，而且即使是常用的字符也远少于 65,536 个，也就是
说用 16bit 编码方式就能表达常用字符。

#### UTF-8

UTF8 是一个将 Unicode 码点编码为字节序列的变长编码，现在 UTF8 编码已经是 Unicode 的标准。UTF8 编码
使用 1 到 4 个字节来表示每个 Unicode 码点，ASCII 部分字符只使用 1 个字节，常用字符部分使用 2 或 3 个
字节表示。每个符号编码后第一个字节的高端 bit 位用于表示总共有多少个编码字节。

如果第一个字节的高端 bit 为 0，则表示对应 7bit 的 ASCII 字符，ASCII 字符每个字符依然是一个字节，和传
统的 ASCII 编码兼容。如果第一个字节的高端 bit 是 110，则说明需要 2 个字节；后续的每个高端 bit 都以 10
开头。更大的 Unicode 码点也是采用类似的策略处理。

```
0xxxxxxx                            runes 0-127     (ASCII)
110xxxxx 10xxxxxx                   128-2047        (values < 128 unused)
1110xxxx 10xxxxxx 10xxxxxx          2048-65535      (values < 2028 unused)
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx 65536-0x10ffff  (other values unused)
```

变长的编码无法直接通过索引来访问第 n 个字符，但是 UTF8 编码获得了很多额外的优点。首先，UTF8 编码比较紧凑，
完全兼容 ASCII 码，并且可以自动同步：它可以通过向前回朔最多 2 个字节就能确定当前字符编码的开始字节的位置。
它也是一个前缀编码，所以当从左向右解码时不会有任何歧义也并不需要向前查看。没有任何字符的编码是其他字符编码
的子串，或是其它编码序列的子串，因此搜索一个字符时只要搜索它的字节编码序列即可，不用担心前后的上下文会对搜索
产生干扰。同时 UTF8 编码的顺序和 Unicode 码点的顺序一致，因此可以直接排序 UTF8 编码序列。同时因为没有嵌入
的 NUL(0) 字节，可以很好地兼容那些使用 NUL 作为字符串结尾的编程语言。

Go 语言字符串面值中的 Unicode 转义字符让我们可以通过 Unicode 码点输入特殊字符。有两种形式：`\uhhhh`对
应 16bit 的码点值，`\Uhhhhhhhh`对应 32bit 的码点值，其中 h 是一个十六进制数字；一般很少需要使用 32bit
的形式。

得益于 UTF8 编码优良的设计，诸多字符串操作都不需要解码操作。我们可以不用解码直接测试一个字符串是否是另一个
字符串的前缀：

```go
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
```

或者是后缀测试：

```go
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}
```

或者是包含子串测试：

```go
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i ++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
```

### 常量

常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean，string 或
数字。

常量的值不可修改，这样可以防止在运行期被意外或恶意的修改。

所有常量的运处都可以在编译期完成，这样可以减少运行时的工作，也方便其他编译优化。当操作数是常量时，
一些运行时的错误也可以在编译时被发现，例如整数除零，字符串索引越界，任何导致无效浮点数的操作等。

常量间的所有算术运算，逻辑运算和比较运算的结果也是常量，对常量的类型转换操作或以下函数调用都是
返回常量结果：len, cap, real, imag, complex 和 unsafe.Sizeof。

因为它们的值是在编译期就确定的，因此常量可以是构成类型的一部分，例如用于指定数组类型的长度：

```
const IPv4Len = 4

// parseIPv4 parses an IPv4 address (d.d.d.d).
func parseIPv4(s string) IP {
    var p [IPv4Len]byte
    // ...
}
```

一个常量的声明也可以包含一个类型和一个值，但是如果没有显式指明类型，那么将从右边的表达式推断
类型。在下面的代码中，time.Duration 是一个命名类型，底层类型是 int64，time.Minute 是对
应类型的常量。下面声明的两个常量都是 time.Duration 类型。

```
const noDelay time.Duration = 0
const timeout = 5 * time.Minute
fmt.Printf("%T %[1]v\n", noDelay)       // "time.Duration 0"
fmt.Printf("%T %[1]v\n", timeout)       // "time.Duration 5m0s"
fmt.Printf("%T %[1]v\n", time.Minute)   // "time.Duration 1m0s"
```

如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式
则表示使用前面常量的初始化表达式写法，对应的常量类型也一样的。例如：

```
const (
    a = 1
    b
    c = 2
    d
)

fmt.Println(a, b, c, d) // "1 1 2 2"
```

#### iota 常量生成器

iota 常量生成器用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。在一个
const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的
行加一。以下是来自 time 包的例子，首先定义了一个 Weekday 命名类型，然后为一周的每天定义了
一个常量，从周日 0 开始。在其它编程语言中，这种类型一般被称为枚举类型。

```
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

周日将对应 0，周一为 1，如此等等。

也可以在复杂的常量表达式中使用 iota，下面是来自 net 包的例子，用于给一个元符号整数的最低 5
bit 的每个 bit 指定一个名字：

```
type Flags int

const (
    FlagUp Flags = 1 << iota    // 
    FlagBroadcast
    FlagLoopback
    FlagPointToPoint
    FlagMulticast
)
```

#### 无类型常量

Go 语言的常量有个不同寻常之处。虽然一个常量可以有任意一个确定的基础类型，但是许多常量并没有一个
明确的基础类型。编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算：可以
认为至少有 256bit 的运算精度。这里有六种未明确类型的常量类型，分别是无类型的布尔型，无类型的
整数，无类型的字符，无类型的浮点数，无类型的复数，无类型的字符串。

通过延迟明确常量的具体类型，无类型的常量不仅可以提供更高的运算精度，而且可以直接用于更多的表达
式而不需要显式的类型转换。