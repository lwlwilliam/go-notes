### 如何检测数据竞争

> 原文：[https://yourbasic.org/golang/detect-data-races/](https://yourbasic.org/golang/detect-data-races/)

**数据竞争很容易发生而且很难 debug。幸运地是，go 的 runtime 常常可以提供帮助**。

使用`-race`可以启用内置的数据竞争检测器。

```
$ go test -race [package]
$ go run -race [package]
```

#### 示例

这里有一个数据竞争的程序：[detectDataRace.go](../src/detectDataRace.go)

加上`-race`选项来运行该程序会显示在第 8 行的写入和第 10 行的读取存在竞争。

```
$ go run -race detectDataRace.go
0
==================
WARNING: DATA RACE
Write at 0x00c420086008 by goroutine 6:
  main.main.func1()
	/Applications/MAMP/htdocs/notes/web/Golang/Translations/src/detectDataRace.go:8 +0x54

Previous read at 0x00c420086008 by main goroutine:
  main.main()
	/Applications/MAMP/htdocs/notes/web/Golang/Translations/src/detectDataRace.go:10 +0x8e

Goroutine 6 (running) created at:
  main.main()
	/Applications/MAMP/htdocs/notes/web/Golang/Translations/src/detectDataRace.go:7 +0x7d
==================
Found 1 data race(s)
exit status 66
```

#### 详细

数据竞争检测器不执行任何静态分析。它检查运行时的内存访问，只检查实际执行的代码路径。

它在 darwin/amd64, freebsd/amd64, linux/amd64 和 windows/amd64 上运行。

开销会不同，但通常内存使用量增加 5-10 倍，执行时间增加 2-20 倍。
