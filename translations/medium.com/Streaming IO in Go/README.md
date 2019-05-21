### Go 语言的 IO 流

在 Go 中，IO 操作是通过使用原语来把数据模拟为可读写的字节流来实现的。为了这个实现，Go 的 io 包提供了`io.Reader`和`io.Writer` 接口分别进行
数据的输入和输出操作，如下图所示：

![figure1.png](./images/figure1.png)

Go 附带了许多支持从内存结构、文件、网络连接等资源的 IO 流操作的 API。这篇文章着重于创建 Go 程序，这些程序能够使用`io.Reader`和`io.Writer`接口的自定义实现和标准库的实现来传输数据。

#### io.Reader

reader 由`io.Reader`接口表示，它可以从一些数据源读取数据到一个可以流式传输以及消费的传输缓冲区中，如下图所示：

![figure2.png](./images/figure2.png)

要使一个类型起到 reader 的作用，必须要使其实现`io.Reader`接口中的`Read(p []byte)`方法，如下所示：

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

`Read()`方法的实现应该返回读取的字节数，如果出错了就返回它的错误。如果数据源已经耗尽，`Read()`方法应该返回`io.EOF`。

##### 读取规则（附加）

在 Reddit 反馈后，我决定添加关于读取的这一小节，希望（对大家）有所帮助。reader 的行为取决于它的实现，但有一些规则来自`io.Reader`文档，
当你在直接从 reader 消费时应该知道这些规则：

1.  当可能时，`Read()`会读取`len(p)`字节到 p 中；

2.  `Read()`调用后，n 可能会小于`len(p)`；

3.  出错时，`Read()`可能仍然返回 n 个字节到缓冲区 p 中。例如，从 TCP socket 中读取时突然连接关闭了。这取决于你的需要，你可以选择保存这些字节到 p 中或者重试；
    
4.  当`Read()`耗尽可用的数据时，reader 可能返回非零的 n 和 err = `io.EOF`。然而，这取决于实现，流结束时，reader 可能选择返回非零的 n 和 err = nil。在这种情况下，任何后续的读取都必须返回 n = 0，err = `io.EOF`；
    
5.  **最后，当一个`Read()`的调用返回 n = 0 和 err = nil 时并不意味着 EOF，因为下一次调用`Read()`可能返回更多数据；**

如你所见，可能直接从 reader 正确地读取流会比较棘手。幸运的是，标准库的 reader 遵循了明智的方法，使其易于流式处理。不过，在使用 reader 之前，最好还是查阅其文档。

##### 从 reader 流式传输数据

直接从 reader 流式传输数据是容易的。`Read()`方法被设计用于在循环体中调用，每次循环都从数据源中读取一块数据到 p 缓冲区中。循环会一直持续到方法返回`io.EOF`错误。

下面是一个简单的例子，该例使用了一个用`strings.NewReader(string)`创建的字符串 reader，用来流式传输字符串源中的字节：[string_reader.go](./src/string_reader.go)。

以上源码用`make([]byte, 4)`创建了一个 4 字节的传输缓冲区 p。该缓冲区用来保存小于字符串源长度的数据。这是为了演示怎样正确地从大于缓冲区的
数据源流式传输数据块。

**更新**：有人在 Reddit 中指出上例有一个 bug。代码永远不会捕获非 nil 错误且不为`io.EOF`的实例。下面修复了代码[string_reader2.go](./src/string_reader2.go)。

##### 实现自定义 io.Reader

在上一节中使用了标准库实现的 IO reader。现在，让我们看看怎样自己实现。下面是一个`io.Reader`的简单实现，它会从流中过滤掉非字母字符[alpha_reader.go](./src/alpha_reader.go)。

当程序执行时，它会打印：

```bash
$ go run alpha_reader.go
HelloItSamwhereisthesum
```

##### 链式 Reader

标准库已经实现了很多 reader。用一个 reader 作为另一个 reader 的数据源是很普遍的。链式 reader 允许一个 reader 重用另一个的逻辑，如以下源码片段中，对 alphaReader 进行了修改，使得其可以接受一个`io.Reader`作为它的数据源。通过把流传到根 reader 可以降低代码复杂性。代码如下[alpha_reader2.go](./src/alpha_reader2.go)

这种方式的另一个优点是 alphaReader 现在可以从任何 reader 实现那里读取数据。例如，以下代码片段展示了 alphaReader 是怎样与`os.File`组合起来过滤文件中的非字母字符的。

#### io.Writer

writer 由接口`io.Writer`表示，（它可以）使数据从缓冲区流入并将其写入目标源，如下图所示：

![writer](./images/writer.png)

所有的流式 writer 必须实现`io.Writer`接口的`Write(p []byte)`方法，设计这个方法的目的是为了从缓冲区 p 中读取数据并将其写入指定的目标源中。

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

`Write()`方法的实现应该返回写入的字节数，如果出错的话就返回错误。

##### 使用 writer

标准库附带很多预先实现的`io.Writer`类型。直接使用 writer 很简单，如下面的代码片段所示，它使用`bytes.Buffer`类型作为`io.Writer`将数据写入内存缓冲区，[using_writer](./src/using_writer.go)。

##### 实现自定义 io.Writer

本节中的代码演示了如何实现名为 chanWriter 的自定义`io.Writer`，它将内容以字节序列的形式写入 Go channel，[chan_writer](./src/chan_writer.go)。

要使用 writer，代码只需简单地在`main()`函数中调用`writer.Writer()`方法（在单独的 goroutine 中）。因为 chanWriter 也实现了接口`io.Closer`，所以调用`writer.Close()`方法来正确关闭 channel，以避免访问通道时出现死锁。

#### 有用的 IO 类型和包

如上所述，Go 标准库具有许多有用的功能和其他类型，使其易于使用流式 IO。

##### os.File

`os.File`类型表示本地系统中的一个文件。它表示`io.Reader`和`io.Writer`，因此，可以用于任何 IO 流的上下文中。例如，以下示例演示了怎样依次地直接往文件中写入字符串切片，[file_write](./src/file_write.go)。

与之相反，`io.File`类型也可以作为 reader 流式传输本地文件系统的文件内容。例如，以下源码片段读取文件内容并将其打印出来，[file_read](./src/file_read.go)

##### 标准输出、标准输入和标准错误

os 包公开了三个变量，`os.Stdout`，`os.Stdin`和`os.Stderr`，它们是`*os.File`类型，分别表示操作系统的标准输出，标准输入和标准错误的文件句柄。例如，以下源码片段直接打印到标准输出中，[stdout_write](./src/stdout_write.go)。

##### io.Copy()

函数`io.Copy()`使得从 reader 的数据源到 writer 的目标源的数据流式传输变得很简单。它将 for-loop 模式（我们已经见过的）进行了抽象，并正确处理`io.EOF`和字节计数。

以下是上一个程序的简单版本，它复制内存中的 reader 数据 proverbs 到 writer 文件中，[io_copy](./src/io_copy.go)。

类似地，我们可以使用`io.Copy()`函数重写以前从文件读取并打印到标准输出的程序，[io_copy2](./src/io_copy2.go)。

##### io.WriteString()

这个函数为写入字符串到指定 writer 提供了方便，[write_str](./src/write_str.go)。

##### 管道 writer 和 reader

`io.PipeWriter`和`io.PipeReader`类型将 IO 操作建模为内存管道。数据被写入管道 writer 写入端，并使用独立的 goroutine 在管道的读取端读取。下面使用`io.Pipe()`创建管道 reader/writer 对，然后用它从缓冲区 proverbs 复制数据到`io.Stdout`，[io_pipe](./src/io_pipe.go)。

##### 缓冲 IO

Go 通过 bufio 包来支持缓冲 IO，这使得文本内容更加容易处理。例如，以下程序逐行读取以`\n`分隔的文件内容，[bufread](./src/bufread.go)。

##### 通用包

ioutil 包是 io 包的子包，提供了几个方便的 IO 函数。例如，以下程序使用了`ReadFile()`函数来把一个文件内容加载到 []byte 中，[io_util](./src/io_util.go)。

#### 总结

这篇文章演示了如何使用`io.Reader`和`io.Writer`接口在程序中实现流式 IO。阅读本文后，你应该能够理解如何利用 io 包来创建 IO 流式传输数据的程序。有大量的例子以及文章向你演示了怎样为自定义功能创建自己的`io.Reader`和`io.Writer`类型。

这是一个介绍性的讨论，仅仅停留在表面以及支持流式 IO 的 Go 包的范围，例如，没有探究文件 IO、缓冲 IO、网络 IO 或格式化 IO。我希望这能让你了解 Go 中可能的流式 IO 习惯用法。

和往常一样，如果你觉得这篇文章对你有用，请点击拍手图标推荐本文。

另外，不要忘了看我的 Go 语言书，`Packt`出版的`Learning Go Programming`。


> 原文链接：[https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185](https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185)
