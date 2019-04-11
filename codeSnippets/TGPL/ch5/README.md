### 函数声明

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。

```go
func name(parameter-list) (result-list) {
	body
}
```

如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。

返回值可以像形式参数一样被命名。在这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为 0。

如果一组形参或返回值有相同的类型，就不必为每个形参都写出参数类型。下面两个声明是等价的：

```go
func f(i, j, k int, s, t string)                { /* ... */ }
func f(i int, j int, k int, s string, t string) { /* ... */ }
```

下面用 4 种方法声明拥有 2 个 int 型参数和 1 个 int 型返回值的函数。blank identifier(即`_`符号)可以强调某个参数未被使用。

```go
func add(x int, y int) int   {return x + y}
func sub(x, y int) (z int)   {z = x - y; return}
func first(x int, _ int) int {return x}
func zero(int, int) int      {return 0}

fmt.Printf("%T\n", add)     // "func(int, int) int"
fmt.Printf("%T\n", sub)     // "func(int, int) int"
fmt.Printf("%T\n", first)   // "func(int, int) int"
fmt.Printf("%T\n", zero)    // "func(int, int) int"
```

函数的类型被称为函数的`标识符`。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型和
标识符。形参和返回值的变量名不影响函数标识符也不影响它们是否可以以省略参数类型的形式表示。

在函数调用时，Go 语言没有默认参数值，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。

在函数体中，函数的形参作为局部变量，被初始化为调用者提供的值。函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。

实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针、slice、map、
function、channel 等类型，实参可能会由于函数的间接引用被修改。

可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以 Go 实现的。这样的声明定义了函数标识符。

```go
package math

func Sin(x float64) float64 // implemented in assembly language
```

### 错误

在 Go 中有一部分函数总是能成功的运行。比如 strings.Contains 和 strconv.FormatBool 函数，
对各种可能的输入都做了良好的处理，使得运行时几乎不会失败，除非遇到灾难性的，不可预料的情况，比如
运行时的内存溢出。导致这种错误的原因很复杂，难以处理，从错误中恢复的可能性也很低。

还有一部分函数只要输入的参数满足一定条件，也能保证运行成功。比如 time.Date 函数，该函数将年月日
等参数构造成 time.Time 对象，除非最后一个参数（时区）是 nil。这种情况下会引发 panic 异常。panic
是来自被调函数的信号，表示发生了某个已知的 bug。一个良好的程序永远不应该发生 panic 异常。

对于大部分函数而言，永远无法确保能否成功运行。这是因为错误的原因超出了程序员的控制。举个例子，任何
进行 I/O 操作的函数都会面临出现错误的可能，只有没有经验的程序员才会相信读写操作不会失败，即使是简
单的读写。因此，当本该可信的操作出乎意料的失败后，我们必须弄清楚导致失败的原因。

在 Go 的错误处理中，错误是软件包 API 和应用程序用户界面的一个重要组成部分，程序运行失败仅被认为是
几个预期的结果之一。

对于那些将运行失败看作是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个，来传递错误信息。
如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为 ok。例如下例：

```go
value, ok := cache.Lookup(key)
if !ok {
	// ...
}
```

通常，导致失败的原因不止一种，尤其是对 I/O 操作而言，用户需要了解更多的错误信息，因此，额外的返回
值不再是简单的布尔类型，而是 error 类型。

内置的 error 是接口类型，它的值可能是 nil 或 non-nil。nil 意味着函数运行成功，non-nil 表示失
败。对于 non-nil 的 error 类型，可以通过调用 error 的 Error 函数或者输出函数获得字符串类型的
错误信息。

通常，当函数返回 non-nil 的 error 时，其他的返回值是未定义的，这些未定义的返回值应该被忽略。然而，
有少部分函数在发生错误时，仍然会返回一些有用的返回值。比如，当读取文件发生错误时，Read 函数会返回
读取的字节数以及错误信息。对于这种情况，正确的处理方式应该是先处理这些不完整的数据，再处理错误。

在 Go 中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常(exception)，这
使得 Go 有别于那些将函数运行失败看作是异常的语言。虽然 Go 有各种异常机制，但这些机制仅被使用在处理
那些未被预料到的错误，即 bug，而不是那些在健壮程序中应该被避免的程序错误。

Go 这样设计的原因是由于对于某个应该在控制流程中处理的错误而言，将这个错误以异常的形式抛出会混乱对
错误的描述，这通常会导致一些糟糕的后果。当某个程序错误被当作异常处理后，这个错误会将堆栈的信息返回
给终端用户，这些信息复杂且无用，无法帮助定位错误。

正因此，Go 使用控制流机制（如 if 和 return）处理异常，这使得编码人员能更多地关注错误处理。

#### 错误处理策略

根据情况的不同，有很多错误的处理方式，常用的五种方式如下。

> 首先，也是最常用的方式是传播错误。这意味着函数中某个子程序的失败，会变成该函数的失败。

```go
resp, err := http.Get(url)
if err != nil {
	return nil, err
}
```

对 http.Get 调用失败，直接返回 HTTP 错误给调用者。

```go
doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
	rturn nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```

当对 html.Parse 调用失败时，不会直接返回 html.Parse 的错误，因为缺少两条重要信息：1. 错误发生
在解析器；2. url 已经被解析。这些信息有助于错误的处理。编写错误信息时，要确保错误信息对问题细节的
描述是详尽的。尤其是要注意错误信息表达的一致性，即相同的函数或同包内的同一组函数返回的错误在构成
和处理方式上是相似的。一般而言，被调函数会将调用信息和参数信息作为发生错误时的上下文放在错误信息中
并返回给调用者，调用者需要添加一些错误信息中不包含的信息，比如添加 url 到 html.Parse 返回的错误
中。

> 处理错误的第二种策略。如果错误的发生是偶然性的，或由不可预知的问题导致的。一个明智的选择是重新尝
试失败的操作。在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试。

[wait.go](./cmd/wait.go)

> 如果错误发生后，程序无法继续运行，就可以采用第三种策略：输出错语信息并结束程序。需要注意的是，这
种策略只应在 main 中执行。对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，
即遇到了 bug，才能在库函数中结束程序。

同见[wait.go](./cmd/wait.go)

> 第四种策略：有时，只需要输出错误信息就足够了，不需要中断程序的运行。可以通过 log 包提供函数。

```go
if err := Ping(); err != nil {
	log.Printf("ping failed: %v; networking disabled", err)
}
```

或者标准错误流输出错误信息。

```go
if err := Ping(); err != nil {
	fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
}
```

> 第五种，也是最后一种策略：可以直接忽略掉错误。

```go
die, err := ioutil.TempDir("", "scratch")
if err != nil {
	rturn fmt.Errorf("failed to create temp dir: %v\n", err)
}
// ...use temp dir
os.RemoveAll(dir)   // ignore errors: $TMPDIR is cleaned periodically
```

尽管 os.RemoveAll 会失败，但上面的例子并没有做错误处理。这是因为操作系统会定期的清理临时目录。
正因如此，虽然程序没有处理错误，但程序的逻辑不会因此受到影响。我们应该在每次函数调用后，都养成考
虚错误处理的习惯，当决定忽略某个错误时，应该清晰地记录下自己的意图。

#### 文件结尾错误(EOF)

从文件中读取 n 个字节。如果 n 等于文件的长度，读取过程的任何错误都表示失败。如果 n 小于文件的长度，调用者会重复的读取固定大小
的数据直到文件结束。这会导致调用者必须分别处理由文件结束引起的各种错误。基于这样的原因，io 包保证任何由文件结束引起的读取失败
都返回同一个错误——io.EOF，该错误在 io 包中定义：

```go
package io

import "errors"

// EOF is the error returned by Read when no more input is available.
var EOF = errors.New("EOF")
```

### 函数值

在 Go 中，函数被看作第一类值(first-class values)：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。

```go
func square(n int) int {return n * n}
func negative(n int) int {return -n}
func product(m, n int) int {return m * n}

f := square
fmt.Println(f(3))   // "9"

f = negative
fmt.Println(f(3))   // "-3"
fmt.Printf("%T\n", f)   // "func(int) int"

f = product // compile error: can't assign func(int, int) int to func(int) int
```

函数类型的零值是 nil。调用值为 nil 的函数值会引起 panic 错误：

```go
var f func(int) int
f(3)    // 此处 f 的值为 nil，会引起 panic 错误
```

函数值可以与 nil 比较：

```go
var f func(int) int
if f != nil {
	f(3)
}
```

但是函数值之间是不可比较的，也不能用函数值作为 map 的 key。

函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为。标准库中包含许多这样的例子。以下的代码展示了如何使用这个技巧。strings.Map 对
字符串中的每个字符调用 add1 函数，并将每个 add1 函数的返回值组成一个新的字符串返回给调用者。

```go
func add1(r rune) rune {return r + 1}

fmt.Println(strings.Map(add1, "HAL-9000"))  // "IBM.:111"
fmt.Println(strings.Map(add1, "VMS"))  // "WNT"
fmt.Println(strings.Map(add1, "Admix"))  // "Benjy"
```

### 匿名函数

拥有函数名的函数只能在包级语法块中被声明，通过函数字面量(function literal)，可以绕过这一限制，在任何表达式中表示一个函数值。函数字面量
的语法和函数声明相似，区别在于 func 关键字后没有函数名。函数值字面量是一种表达式，它的值被称为`匿名函数(anonymous function)`。

函数字面量允许我们在使用函数时，再定义它。

```go
fmt.Println(strings.Map(func(r rune) rune {return r + 1}, "HAL-9000"))
```

更为重要的是，通过这种方式定义的函数可以访问完整的词法环境(lexical environment)，这意味着在函数中定义的内部函数可以引用该函数的
变量，[squares.go](./cmd/squares.go)。

函数 squares 返回另一个类型为 func() int 的函数。对 squares 的一次调用会生成一个局部变量 x 并返回一个匿名函数。每次调用匿名函数时，
该函数都会先使 x 的值加 1，再返回 x 的平方。第二次调用 squares 时，会生成第二个 x 变量，并返回一个新的匿名函数。新匿名函数操作的是第
二个 x 变量。

squares 的例子证明，函数值不仅仅是一串代码，还记录了状态。在 squares 中定义的匿名内部函数可以访问和更新 squares 中的局部变量，这意味
着匿名函数和 squares 中，存在变量引用。这就是函数值属于引用类型和函数值不可比较的原因。Go 使用闭包(closures)技术实现函数值，Go 程序员也
把函数值叫做闭包。

由此可见，变量的生命周期不由它的作用域决定：squares 返回后，变量 x 仍然隐式地存在于 f 中。

接下来，讨论一个有点学术性的例子，考虑这样一个问题：给定一些计算机课程，每个课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；
目标是选择出一组课程，这组课程必须确保按顺序学习时，能全部被完成。每个课程的前置课程如[toposort](./toposort)。

这类问题被称为拓扑排序。从概念上说，前置条件可以构成有向图。图中的顶点表示课程，边表示课程间的依赖关系。显然，图中应该无环，这也就是说从
某点出发的边，最终不会回到该点。下面的代码用深度优先搜索了整张图，获得了符合要求的课程序列[toposort](./cmd/toposort.go)。

#### 警告：捕获迭代变量

这里介绍 Go 词法作用域的一个陷阱。

考虑这样一个问题：你被要求首先创建一些目录，再将目录删除。为了使代码简单，忽略了所有的异常处理。

```go
var rmdirs []func()
for _, d := range tempDirs() {
	dir := d  // NOTE: necessary!
	os.MkdirAll(dir, 0755)  // creates parent directories too
	rmdirs = append(rmdirs, func() {
		os.RemoveAll(dir)
	})
}

// ... do some work...
for _, rmdir := range rmdirs {
	rmdir()  // clean up
}
```

为什么要在循环体中用循环变量 d 赋值一个新的局部变量，而不是像下面的代码一样直接使用循环变量 dir。需要注意，
下面的代码是错误的。

```go
var rmdirs []func()
for _, dir := range tempDirs() {
	os.MkdirAll(dir, 0755)
	rmdirs = append(rmdirs, func() {
		os.RemoveAll(dir)  // NOTE: incorrect!
	})
}
```

问题的原因在于循环变量的作用域。在上面的程序中，for 循环语句引入了新的词法块，循环变量 dir 在这个词法块中
被声明。在该循环中生成的所有函数值都共享相同的循环变量。需要注意，函数值中记录的是循环变量的内存地址，而不
是循环变量某一时刻的值。以 dir 为例，后续的迭代会不断更新 dir 的值，当删除操作执行时，for 循环已完成，dir
中存储的值等于最后一次迭代的值。这意味着，每次对 os.RemoveAll 调用删除的都是相同的目录。

通常，为了解决这个问题，会引入一个与循环变量同名的局部变量，作为循环变量的副本。比如下面的变量 dir，虽然这看
起来很奇怪，但却很有用。

```go
for _, dir := range tempDirs() {
	dir := dir  // declares inner dir, initialized to outer dir
	// ...
}
```

这个问题不仅存在基于 range 的循环，在下例中，对循环变量 i 的使用也存在同样的问题：

```go
var rmdirs []func()
dirs := tempDirs()
for i := 0; i < len(dirs); i ++ {
	os.MkdirAll(dirs[i], 0755)  // OK
	rmdirs = append(rmdirs, func() {
		os.RemoveAll(dirs[i])  // NOTE: incorrect!
	})
}
```

如果使用 go 语句或者 defer 语句会经同到此类问题。这不是 go 或 defer 本身导致的，而是因为它们都会等待循
环结束后，再执行函数值。

### 可变参数

参数数量可变的函数称为可变参数函数。在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号"..."，
这表示该函数会接收任意数量的该类型参数。

在[sum.go](./cmd/sum.go)中，在第 14 行，调用者隐式地创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传给
被调函数。如果原始参数已经是切片类型，该如何传递给 sum？只需在最后一个参数后加上省略符，如第 15 行。

虽然在可变参数函数内部，...int 型参数的行为看起来很像切片类型，但实际上，可变参数函数和以切片作为参数的函数是不同的，如上例的
第 16、17 行。

### Deferred 函数

只需要在调用普通函数或方法前加上关键字 defer，就完成了 defer 所需要的语法。当 defer 语句被执行时，跟在 defer 后面的函数会被延迟执行。
直到包含该 defer 语句的函数执行完毕时，defer 后的函数才会被执行，不论包含 defer 语句的函数是通过 return 正常结束，还是由于 panic
导致的异常结束。可以在一个函数中执行多条 defer 语句，它们的执行顺序与声明顺序相反。

defer 语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过 defer 机制，不论函数逻辑多复杂，都能保证在任何
执行路径下，资源被释放。

调试复杂程序时，defer 机制也常被用于记录何时进入和退出函数[trace.go](./cmd/trace.go)。

defer 语句中的函数会在 return 语句更新返回值变量后再执行，又因为在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量，
所以，对匿名函数采用 defer 机制，可以使其观察函数的返回值[deferReturn.go](./cmd/deferReturn.go)。

被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值[deferChange.go](./cmd/deferChange.go)。

在循环体中的 defer 语句需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行。下面的代码会导致系统的文件描述符耗尽，因为
在所有文件都被处理之前，没有文件会被关闭。

```go
for _, filename := range filenames {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close() // NOTE: risky; could run out of file descriptors
	// ...process f...
}
```

一种解决方法是将循环体中的 defer 语句移至另外一个函数。在每次循环时，调用这个函数。

```go
for _, filename := range filenames {
	if err := doFile(filename); err != nil {
		return err
	}
}

func doFile(filename string) err {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ..process f...
}
```

[fetch.go](./cmd/fetch.go)将 http 响应信息写入本地文件而不是从标准输出流输出。对 resp.Body.Close 延迟调用就不做解释了。通过 os.Create
打开文件进行写入，`在关闭文件时，没有对 f.Close 采用 defer 机制，因为这会产生一些微妙的错误`。许多文件系统，尤其是 NFS，写入文件时发生的错误
会被延迟到文件关闭时反馈。如果没有检查文件关闭时的反馈信息，可能会导致数据丢失，而我们还误以为写入操作成功。如果 io.Copy 和 f.Close 都失败
了，倾向于将 io.Copy 的错误信息反馈给调用者，因为它先于 f.Close 发生，更有可能接近问题的本质。

### Panic 异常

Go 的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起 panic 异常。

一般而言，当 panic 异常发生时，程序会中断运行，并立即执行在该 goroutine 中被延迟的函数（defer 机制）。随后，程序崩溃并输出日志信息。
日志信息包括 panic value 和函数调用的堆栈跟踪信息。

不是所有的 panic 异常都来自运行时，直接调用内置的 panic 函数也会引发 panic 异常；panic 函数接受任何值作为参数。当某些不应该发生的场景
发生时，就应该调用 panic。比如，当程序到达了某条逻辑上不可能到达的路径：

```go
switch s := suit(drawCard()); s {
case "Spades":      //...
case "Hearts":      // ...		
case "Diamonds":    // ...		
case "Clubs":       // ...		
default:
	panic(fmt.Sprintf("invalid suit %q", s))    // Joker?
}
```

虽然 Go 的 panic 机制类似于其他语言的异常，但 panic 的适用场景有一些不同。由于 panic 会引起程序的崩溃，因此一般用于严重错误，如程序内部
的逻辑不一致。

在健壮的程序中，任何可以预料到的错误，如不正确的输入、错误的配置或是失败的 I/O 操作都应该被优雅的处理，最好的处理方式，就是使用 Go 的错误机制。

为了方便诊断问题，runtime 包允许程序员输出堆栈信息[stack.go](./cmd/stack.go)。

### Recover 捕获异常

通常来说，不应该对 panic 异常做任何处理，但有时，也许我们可以从异常中恢复，至少可以在程序崩溃前，做一些操作。举个例子，当 web 服务器遇到不可
预料的严重问题时，在崩溃前应该将所有连接关闭；如果不做任何处理，会使得客户端一直处于等待状态。如果 web 服务器还在开发阶段，服务器甚至可以将
异常信息反馈到客户端，帮助调试。

如果在 deferred 孙炜中调用了内置函数 recover，并且定义该 defer 语句的函数发生了 panic 异常，recover 会使程序从 panic 中恢复，并返回
panic value。导致 panic 异常的函数不会继续运行，但能正常返回。在未发生 panic 时调用 recover，会返回 nil。

以语言解析器为例，说明 recover 的使用场景。考虑到语言解析器的复杂性，即使某个语言解析器目前工作正常，也无法肯定它没有漏洞。因此，当某个
异常出现时，我们不会选择让解析器崩溃，而是会将 panic 异常当作普通的解析错误，并附加额外信息提醒用户报告此错误。

```go
func Parse(input string) (s *Syntax, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	// ...parser...
}
```
