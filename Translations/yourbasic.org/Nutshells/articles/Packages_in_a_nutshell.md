### 包概述

#### 基础

每个 Go 程序都由包组成，每个包都有一个 import 路径（译注：也不一定都有，package 声明才是必须的），例如：

*   "fmt"
*   "math/rand"
*   "github.com/yourbasic/graph"

标准库里的包都有一个短的 import 路径，例如 "fmt" 和 "math/rand"。第三方包，例如 "github.com/yourbasic/graph"，import 路径都包含了
主机服务(github.com)和一个组织名(yourbasic)。

按照约定，包名是 import 路径的最后一个元素（译注：例如上面 import 的路径的包名如下）：

*   fmt
*   rand
*   graph

要引用其它包的定义必须以包名为前缀，而且其它包中只有以首字母为大写的名称都能被访问。

[packages.go](../src/packages.go)

#### 声明一个包

每个 Go 文件都以包声明开头，包声明只包含包名。

例如，文件`src/math/rand/exp.go`是`math/rand`包的一部分，包含了以下代码。

```
package rand

import "math"

const re = 7.69711747013104972

...
```

不需要担心包名冲突，只有 import 路径必须是唯一的。[如何编写 Go 代码](https://golang.org/doc/code.html)展示了在文件结构中如何组织代码
以及它的包。

#### 包名冲突

可以自定义引用已导入包的名称。

[packagesConflict.go](../src/packagesConflict.go)

#### . 导入

如果一段时间`.`代替了 import 语句中的名称，该包里的所有可导出标识符可以在不使用限定符的情况下访问。

[packagesDotImport.go](../src/packagesDotImport.go)

`.`导入方式会使程序变得难以阅读，通常应该避免这样做。

#### 包下载

`go get`命令下载由 import 路径命名的包以及它们的依赖，然后进行安装。

```
$ go get github.com/yourbasic/graph
```

import 路径对应保存代码的仓库。这减少了未来可能的名称冲突。

[Go Wiki](https://github.com/golang/go/wiki/Projects) 和 [Awesome Go](https://github.com/avelino/awesome-go) 提供了一些高
质量的 Go 名和资源。

更多关于用 Go 工具使用远程仓库的信息，阅读 [Command go: Remote import paths](https://golang.org/cmd/go/#hdr-Remote_import_paths)。

#### 包文档

[GoDoc](https://godoc.org/) 网站为 Bitbucket、GitHub、Google 项目托管和发布平台的所有公共 Go 包提供文档：

*   [https://godoc.org/fmt](https://godoc.org/fmt)
*   [https://godoc.org/math/rand](https://godoc.org/math/rand)
*   [https://godoc.org/github.com/yourbasic/graph](https://godoc.org/github.com/yourbasic/graph)

godoc 命令可以为所有安装在本地的 Go 程序提取并生成文件。以下命令开启了一个 web 服务器`http://localhost:6060/`显示文档。

```
$ godoc -http=:6060 &
```

关于更多如何访问和创建文档的内容，可以查阅文章 [https://yourbasic.org/golang/package-documentation/](https://yourbasic.org/golang/package-documentation/)
