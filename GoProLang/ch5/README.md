### 函数

> 函数声明

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。

```go
func name(parameter-list) (result-list) {
    body
}
``` 

如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。如果一个函数声明不包括返回值列表，那么函数体执行完毕后，不会返回任
何值。

```go
func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(3, 4))  // "5"
```

上述 hypot 函数中，x 和 y 是形参名，3 和 4 是调用时传入的实数，函数返回了一个 float64 类型的值。返回值也可以像形式参数一样被命名。在
这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为 0。

函数的类型被称为函数的标识符。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型和标识符。

**每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值）。在函数调用时，Go 语言没有默认参数值，也没有任何方法可以通过参数名指定形
参，因此形参和返回值的变量名对于函数调用者而言没有意义。**

> 递归

```go
// TODO: 待研究代码
```

> 多返回值

在 Go 中，一个函数可以返回多个值。

> 错误

在 Go 中有一部分函数总是能成功的运行。比如 strings.Contains 和 strconv.FormatBool 函数，对各种可能的输入都做了良好的处理，使得运行几乎
不会失败，除非遇到灾难性的、不可预料的情况，比如运行时的内存溢出。导致这种错误的原因很复杂，难以处理，从错误中恢复的可能性也很低。

还有一部分函数只要输入的参数满足一定条件，也能保证运行成功。比如 time.Date 函数，该函数将年月日等参数构造成 time.Time 对象，除非最后一个参
数（时区）是 nil。这种情况下会引发 panic 异常。panic 是来自被调函数的信号，表示发生了某个已知的 bug。一个良好的程序永远不应该发生 panic 异
常。

对于大部分函数而言，永远无法确保能否成功运行。这是因为错误的原因超出了程序员的控制。举个例子，任何进行 I/O 操作的函数都会面临出现错误的可能，只
有没有经验的程序员才会相信读写操作不会失败，即使是简单的读写。因此，当本该可信的操作出乎意料的失败后，必须弄清楚导致失败的原因。

在 Go 的错误处理中，错误是软件包 API 和应用程序用户界面的一个重要组成部分，程序运行失败仅被认为是几个预期的结果之一。

对于那些将运行失败看作是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个，来传递错误信息。如果导致失败的原因只有一个，额外的返回值可以
是一个布尔值，通常被命名为 ok。比如，cache.Lookup 的失败的唯一原因是 key 不存在，那么代码可以按照下面的方式组织：

```go
value, ok := cache.Lookup(key)
if !ok {
	// ...cache[key] does not exist...
}
```

通常，导致失败的原因不止一种，尤其是对 I/O 操作而言，用户需要了解更多的错误信息。因此，额外的返回值不再是简单的布尔类型，而是 error 类型。

内置的 error 是接口类型。error 类型可能是 nil 或者 non-nil。nil 意味着函数运行成功，non-nil 表示失败。对于 non-nil 的 error 类型，可
以通过调用 error 的 Error 函数或者输出函数获得字符串类型的错误信息。

```go
fmt.Println(err)
fmt.Printf("%v", err)
```

通常，当函数返回 non-nil 的 error 时，其他的返回值是未定义的(undefined)，这些未定义的返回值应该被忽略。然而，有少部分函数在发生错误时，仍然
会返回一些有用的返回值。比如，当读取文件发生错误时，Read 函数会返回可以读取的字节数以及错误信息。对于这种情况，正确的处理方式应该是先处理这些不
完整的数据，再处理错误。因此对函数的返回值要有清晰的说明，以便于其他人使用。

在 Go 中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常(exception)，这使得 Go 有别于那些将函数运行失败看作是异常的
语言。虽然 Go 有各种异常机制，但这些机制仅被使用在处理那些未被预料到的错误，即 bug，而不是那些在健壮程序中应该被避免的程序错误。

Go 这样设计的原因是由于对于某个应该在控制流程中处理的错误而言，将这个错误以异常的形式抛出会混乱对错误的描述，这通常会致一些糟糕的后果。当某个
程序错误被当作异常处理后，这个错误会将堆栈根据信息返回给终端用户，这些信息复杂且无用，无法帮助定位错误。

正因此，Go 使用控制流机制（如 if 和 return）处理异常，这使得编码人员能更多的关注错误处理。

>> 错误处理策略

当一次函数调用返回错误时，调用者应该选择合适的方式处理错误。根据情况的不同，有很多处理方式，先来看看常用的五种方式。

1.  首先，也是最常用的方式是传播错误。这意味着函数中某个子程序的失败，会变成该函数的失败。

    对 http.Get 的调用失败，会直接将这个 HTTP 错误返回给调用者。

    ```go
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    ```
    
    当对 html.Parse 的调用失败时，不会直接返回 html.Parse 的错误，因为缺少两条重要信息：1. 错误发生在解析器；2. url 已经被解析。这些信息
    有助于错误的处理，会构造新的错误信息返回给调用者：
    
    ```go
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    ```
    
    fmt.Errorf 函数使用 fmt.Sprintf 格式化错误信息并返回。当错误最终由 main 函数处理时，错误信息应提供清晰的从原因到后果的因果链。 
    
2.  处理错误的第二种策略。如果错误的发生是偶然性的，或由不可预知的问题导致的。一个明智的选择是重新尝试失败的操作。在重试时，需要限制重试的时间
    间隔或重试的次数，防止无限制的重试。
    
    ```go
    func WaitForServer(url string) error {
        const timeout = 1 * time.Minute
        deadline := time.Now().Add(timeout)
        for tries := 0; time.Now().Before(deadline); tries ++ {
            _, err := http.Head(url)
            if err == nil {
                return nil  // success	
            }
            log.Printf("server not responding (%s); retrying...", err)
            time.Sleep(time.Second << uint(tries))  // exponential back-off
        }
        return fmt.Errorf("server %s failed to respond after %s", url, timeout)
    }
    ```
    
3.  如果错误发生后，程序无法继续运行，就可以采用第三种策略：输出错误信息并结束程序。需要注意的是，这种策略只应在 main 中执行。对库函数而言，
    应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了 bug，才能在库函数中结束程序。
    
    ```go
    // (In function main.)
    if err := WaitForServer(url); err != nil {
        fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
        os.Exit(1)
    }
    ```
    
    调用 log.Fatalf 可以更简洁地代码达到与上文相同的效果。log 中的所有函数，都默认会在错误信息之前输出时间信息。
    
    ```go
    if err := WaitForServer(url); err != nil {
        log.Fatalf("Site is down: %v\n", err)
    }
    ```
    
4.  第四种策略：有时，只需要输出错误信息就足够了，不需要中断程序的运行。可以通过 log 包提供函数

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
    
5.  第五种，也是最后一种策略：可以直接忽略掉错误。

    ```go
    dir, err := ioutil.TempDir("", "scratch")
    if err != nil {
        return fmt.Errorf("failed to create temp dir: %v", err)
    }
    // ...use temp dir...
    os.RemoveAll(dir)  // ignore errors; $TMPDIR is cleaned periodically
    ```
    
    尽管 os.RemoveAll 会失败，但上面的例子并没有做错误处理。这是因为操作系统会定期的清理临时目录。正因如此，虽然程序没有处理错误，但程序的逻
    辑不会因此受到影响。我们应该在每次函数调用后，都养成考虑错误处理的习惯，当你决定忽略某个错误时，应该清晰地记录下你的意图。
    
>> 文件结尾错误(EOF)

函数经常会返回多种错误，程序必须根据错误类型，作出不同的响应。考虑这样一个例子：从文件中读取 n 个字节。如果 n 等于文件的长度，读取过程的
任何错误都表示失败。如果 n 小于文件的长度，调用者会重复地读取固定大小的数据直到文件结束。这导致调用者必须分别处理由文件结束引起的各种错
误。基于这样的原因，io 包保证任何由文件结束引起的读取失败都返回同一个错误——`io.EOF`。该错误在 io 包中定义：

```go
package io

import "errors"

// EOF is the error returned by Read when no more input is available.
var EOF = errors.New("EOF")
```

以下是一个 demo：

```go
in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break  // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed:%v", err) 
    }
    // ...use r... 
}
```

因为文件结束这种错误不需要更多的描述，所以 io.EOF 有固定的错误信息——"EOF"。对于其他错误，我们可能需要在错误信息中描述错误的类型和数量，
这使得我们不能像 io.EOF 一样采用固定的错误信息。

> 函数值

在 Go 中，函数被看作第一类值(first-class values)：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值
(function value)的调用类似函数调用。


```go
func square(n int) int { return n*n }
func negative(n int) int { return -n }
func product(m, n int) int { return m*n }

f := square
fmt.Println(f(3))  // "9"

f = negative
fmt.Println(f(3))  // "-3"
fmt.Println(%T\n", f)  // "func(int) int"

f = product  // compile error: can't assign func(int, int) int to func(int) int
```

函数类型的零值是 nil。调用值为 nil 的函数值会引起 panic 错误：

```go
var f func(int) int
f(3)  // 此处 f 的值为 nil，会引起 panic 错误
```

函数值可以与 nil 比较，但是函数值之间是不可比较的，也不能用函数值作为 map 的 key。

```go
var f func(int) int
if f != nil {
    f(3)
}
```

> 匿名函数

拥有函数名的函数只能在包级语法块中被声明，通过函数字面量(function literal)，我们可绕过这一限制，在任何表达式中表示一个函数值。函数字面
量的语法和函数声明相似，区别在于 func 关键字后没有函数名。函数值字面量是一种表达式，它的值被称为匿名函数(anonymous function)。

函数字面量允许我们在使用函数时，再定义它。

更为重要的是，通过这种方式定义的函数可以访问完整的词法环境(lexical environment)，这意味着在函数中定义的内部函数可以引用该函数的变量。

```go
// squares 返回一个匿名函数
// 该匿名函数每次被调用时都会返回下一个数的平方
func squares() func() int {
    var x int
    return func() int {
        x ++
        return x * x 
    }
}

func main() {
    f := squares()
    fmt.Println(f())  // "1"
    fmt.Println(f())  // "4"
    fmt.Println(f())  // "9"
    fmt.Println(f())  // "16"
}
```

**函数 squares 返回另一个类型为 func() int 的函数。对 squares 的一次调用会生成一个局部变量 x 并返回一个匿名函数。每次调用匿名函数时，
该函数都会先使 x 的值加 1，再返回 x 的平方。第二次调用 squares 时，会生成第二个 x 变量，并返回一个新的匿名函数。新匿名函数操作的是第
二个 x 变量。**

squares 证明，函数值不仅仅是一串代码，还记录了状态。在 squares 中定义的匿名内部函数可以访问和更新 squares 中的局部变量，这意味着匿名
函数和 squares 中，存在变量引用。**这就是函数值属于引用类型和函数值不可比较的原因。** Go 使用闭包(closures)技术实现函数值，Go 程序员
也把函数值叫做闭包。

通过这个例子，可以看到变量的生命周期不由它的作用域决定：squares 返回后，变量 x 仍然隐式的存在于 f 中。

接下来，讨论一个有点学术性的例子，考虑这样一个问题：给定一些计算机课程，每个课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；
目标是选择出一组课程，这组课程必须确保按顺序学习时，能全部被完成。每个课程的前置课程如下：

```go
var prereqs = map[string][]string{
    "algorithms": {"data structures"},
    "calculus": {"linear algebra"},
    "compilers": {
        "data structures",
        "formal languages",
        "computer organization", 
    },
    "data structures": {"discrete math"},
    "databases": {"data structures"},
    "discrete math": {"intro to programming"},
    "formal languages": {"discrete math"},
    "networks": {"operating systems"},
    "operating systems": {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
}
```

这类问题被称作拓扑排序。从概念上说，前置条件可以构成有向图。

```go
func main() {
    for i, course := range topoSort(prereqs) {
        fmt.Printf("%d: \t%s\n", i + 1, course) 
    }
}

func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)
    var visitAll func(items []string)
    visitAll = func(items []string) {
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                visitAll(m[item])
                order = append(order, item) 
            } 
        } 
    }
    
    var keys []string
    for ken := range m {
        keys = append(keys, key) 
    }
    sort.Strings(keys)
    visitAll(keys)
    return order
}
```

[toposort.go](toposort.go)

>> 警告：捕获迭代变量

```go
var rmdirs []func()
for _, d := range tempDirs() {
    dir := d  // NOTE: necessary!
    os.MkdirAll(dir, 0755)  // creates parent directories too
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) 
    })
}

// ...do some work...
for _, rmdir := range rmdirs {
    rmdir()  // clean up
}
```

为什么要在循环体中用循环变量 d 赋值一个新的局部变量，而不是像下面的代码一样直接使用循环变量 dir。需要注意，下面的代码是错误的。

```go
var rmdirs []func()
for _, dir := ranges tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)  // NOTE: incorrect! 
    })
}
```

问题的原因在于循环变量的作用域。在上面的程序中，for 循环语句引入了新的词法块，循环变量 dir 在这个词法块中被声明。在该循环中生成的所有函
数值都共享相同的循环变量。需要注意，函数值中记录的是循环变量的的内存地址，而不是循环变量某一时刻的值。以 dir 为例，后续的迭代会不断更新
 dir 的值，当删除操作执行时，for 循环已完成，dir 中存储的值等于最后一次迭代的值。

> 可变参数

参数数量可变的函数称为可变参数函数。在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号`...`，这表示该函数会接收任意
数量该类型参数。

[sum.go](sum.go)

sum 函数返回任意个 int 类型参数的和。在函数中，vals 被看作是类型为 []int 的切片。调用者隐式地创建一个数组，并将原始参数复制到数组中，再
把数组的一个切片作为参数传给被调函数。如果原始参数已经是切片类型，该如何传递给 sum？只需要最后一个参数后加上省略符。

```go
values := []int{1, 2, 3, 4}
fmt.Println(sum(values...))  // "10"
```

虽然在可变参数函数内部，`...int`参数的行为看起来很像切片类型，但实际上，可变参数函数和以切片作为参数的函数是不同的。

```go
func f(...int) {}
func g([]int) {}
fmt.Println("%T\n", f)  // "func(...int)"
fmt.Println("%T\n", g)  // "func([]int)"
```

可变参数函数经常被用于格式化字符串。

```go
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

linenum, name := 12, "count"
errorf(linenum, "undefined: %s", name)  // "Line 12: undefined: count"
```

interface{} 表示函数的最后一个参数可以接收任意类型。

> Deferred 函数

只需要在调用普通函数或方法前加上关键字 defer，就完成了 defer 所需要的语法。当 defer 语句被执行时，跟在 defer 后面的函数会被延迟执行。直
到包含该 defer 语句的函数执行完毕时，defer 后的函数才会被执行，不论包含 defer 语句的函数是通过 return 正常结束，还是由于 panic 导致的异
常结束。可以在一个函数中执行多条 defer 语句，它们的执行顺序与声明顺序相反。

defer 语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过 defer 机制，不论函数逻辑多复杂，都能保证在任何执行
路径下，资源被释放。释放资源的 defer 应该直接跟在请求资源的语句后。

[title2.go](title2.go)

在处理其他资源时，也可以采用 defer 机制，比如对文件的操作：

[ioutil.go](ioutil.go)

或是处理互斥锁

```go
var mu sync.Mutex
var m = make(map[string]int)
func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
```

调试复杂程序时，defer 机制也常被用于记录何时进入和退出函数。下例中的 bigSlowOperation 函数，直接调用 trace 记录函数的被调情况。
bigSlowOperation 被调时，trace 会返回一个函数值，该函数值会在 bigSlowOperation 退出时被调用。通过这种方式，可以只通过一条语句控制函数的
入口和所有的出口，甚至可以记录函数的运行时间，如例子中的 start。需要注意一点：不要忘记 defer 语句后的圆括号，否则本该在进入时执行的操作
会在退出时执行，而本该在退出时执行的，永远不会被执行。

[trace.go](trace.go)

每一次 bigSlowOperation 被调用，程序都会记录函数的进入，退出，持续时间。

defer 语句中的函数会有 return 语句更新返回值变量后再执行，又因为在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量，所以，
对匿名函数采用 defer 机制，可以使其观察函数的返回值。

```go
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d)\n", x, result) }()
	return x + x
}

_ = double(4)  // "double(4) = 8"
```

在循环体中的 defer 语句需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行。下面的代码会导致系统的文件描述符耗尽，因为所有
文件都被处理之前，没有文件会被关闭。

```go
for _, filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
        return err 
    }
    defer f.Close()  // NOTE: risky; could run out of file descriptors
   // ...process f...
}
```

一种解决方法是将循环体中的 defer 语句移至另外一个函数。在每次循环时，调用这个函数。

```go
for _, filename = range filenames {
	if err := doFile(filename); err != nil {
		return err
	}
}

funcn doFile(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f...
}
```

[fetch.go](fetch.go)

对 resp.Body.Close 延迟调用不做解释。上例中，通过 os.Create 打开文件进行写入，在关闭文件时，没有对 f.Close 采用 defer 机制，因为这
会产生一些微妙的错误。许多文件系统，尤其是 NFS，写入文件时发生的错误会被延迟到文件关闭时反馈。如果没有检查文件关闭时的反馈信息，可能会导
致数据丢失，而我们还误以为写入操作成功。如果 io.Copy 和 f.Close 都失败了，我们倾向于将 io.Copy 的错误信息反馈给调用者，因为它先于 
f.Close 发生，更有可能接近问题的本质。

> Panic 异常

Go 的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起 panic 异常。

一般而言，当 panic 异常发生时，程序会中断运行，并立即执行在该 goroutine 中被延迟的函数（defer 机制）。