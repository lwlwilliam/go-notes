### 对比公有和私有

> 原文：[https://yourbasic.org/golang/public-private/](https://yourbasic.org/golang/public-private/)

**package 是 Go 私有封装的最小单元**。

*	所有在 package 中定义的标识符在整个 package 中都可见；
*	在导入 package 时只能访问 package 中可导出的标识符；
*	如果标识符以大写字母开头，它就是可导出的；

可导出和不可导出的标识符用来描述 package 中的公共接口以及防止某些程序错误。

**警告：不可导出的标识符不是安全的措施，它不能隐藏或者保护任何信息**。

#### 示例

在这个 package 中，只有 StopWatch 和 Start 标识符是可导出的。

```
package timer

import "time"

// StopWatch 是简易的时钟功能，它的零值与总时长为 0 的停止的时钟一样
type StopWatch struct {
	start time.Time
	total time.Duration
	running bool
}

// Start 用来开启时钟
func (s *StopWatch) Start() {
	if !s.running {
		s.start = time.Now()
		s.running = true
	}
}
```

StopWatch 和它可导出的方法可以在另一个 package 中被导入和使用。

```
package main

import "timer"

func main() {
	clock := new(timer.StopWatch)
	clock.Start()
	if clock.running {	// 非法
		// ...
	}
}
```

```
# command-line-arguments
../src/visibility.go:8:10: clock.running undefined (cannot refer to unexported field or method running)
```

source code: [visibility.go](../src/visibility.go)、[timer.go](../src/timer/timer.go)
