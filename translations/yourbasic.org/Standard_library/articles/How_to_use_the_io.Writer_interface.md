### 怎样使用 io.Writer 接口

> 原文：[https://yourbasic.org/golang/io-writer-interface-explained/](https://yourbasic.org/golang/io-writer-interface-explained/)

#### 基础

io.Writer 接口表示可以往其中写入字节流的实体。

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

Write 可以从 p 中向底层数据流中写入最多 len(p) 个字节——它返回写入的字节数以及导致写入提前停止的任何错误。

标准库提供了许多 Writer 的实现，许多实用程序都接受 Writer 作为输入。

#### 怎样使用内置的 writer（3 个示例）

作为第一个示例，你可以使用 fmt.Fprintf 函数直接写入 bytes.Buffer。之所以可以这么做，是由于以下原因：

*   bytes.Buffer 有 Write 方法；
*   fmt.Fprintf 的第一个参数为 Writer；

[buffer.go](../src/buffer.go)

类似地，你可以直接写入文件或其它流，例如 HTTP 连接。完整的代码示例，请参阅 [HTTP 服务器示例文章](https://yourbasic.org/golang/http-server-example/)。

这是 Go 中非常常见的模式。作为另一个示例，你可以通过将文件复制到合适的 hash.Hash 对象的 io.Writer 函数中来计算文件的哈希值。
请参阅[哈希检验和的相关代码](https://yourbasic.org/golang/hash-md5-sha256-string-file/#file)。

#### 优化字符串写入

标准库里有些 Writer 有额外的 WriteString 方法。这个方法比标准的 Write 方法效率更高，因为它是直接写入字符串，并没有分配 byte slice。

你可以直接使用 io.WriteString() 函数来实现优化。

```go
func WriteString(w Writer, s string) (n int, err error)
```

如果 w 实现了 WriteString 方法，则直接被调用。否则，w.Write 其实只会调用一次。