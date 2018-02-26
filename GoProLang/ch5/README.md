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


