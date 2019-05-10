### 怎样使用 io.Reader 接口

> 原文：[https://yourbasic.org/golang/io-reader-interface-explained/](https://yourbasic.org/golang/io-reader-interface-explained/)

#### 基础

io.Reader 接口表示一个可以读取字节流的实体。

```go
type Reader interface {
	Read(buf []byte) (n int, err error)
}
```

Read 可以向 buf 中读取最多 len(buf) 个字节，并返回已读取到的字节数，当流结束时，返回 io.EOF 错误。

标准库提供了许多 Reader 的实现（包括内存中的字节缓冲区，文件和网络连接），Reader 被许多实用程序接受作为输入（包括 HTTP 客户端和服务端实现）。

#### 使用内置的 reader

例如，你可以用 strings.Reader 函数从字符串创建一个 Reader，然后把 Reader 直接传到 net/http 包的 http.Post 函数。然后将 Reader 
作为 post 的数据源。

```go
r := strings.NewReader("my request")
resp, err := http.Post("http://foo.bar", "application/x-www-form-urlencoded", r)
```

由于 http.Post 用 Reader 代替了 []byte 传输，因此，使用文件的内容就比较简单了。完整代码[reader.go](../src/reader.go)。

#### 直接从字节流读取

你可以直接使用 Read 函数（这是最不常用的例子）：[direclyRead.go](../src/directlyRead.go)。

使用 io.ReadFull 读取 len(buf) 字节到 buf：[readFull.go](../src/readFull.go)。

使用 ioutil.ReadAll 读取所有内容：[readAll.go](../src/readAll.go)。

#### 缓冲读取及扫描

bufio.Reader 和 bufio.Scanner 类型封装了 Reader，创建了另一个同样实现了 Reader 接口的 Reader，为文件输入提供了缓冲及一些帮助。

下例使用了 bufio.Scanner 对文本中的单词进行计数[counter.go](../src/counter.go)。