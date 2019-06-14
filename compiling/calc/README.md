### 简单的四则运算计算器

使用方法：

```go
package main

import (
	"fmt"
	"github.com/lwlwilliam/go-notes/compiling/calc"
)

func main() {
    fmt.Println(calc.Calc("3 * (3 + 3) / 2")) // 9
}
```