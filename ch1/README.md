### 入门

> Hello, world!

**helloworld.go**

```
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

Go 语言的代码通过包(package)组织，包类似于其它语言里的库(libraries)或者模块(modules)。一个包由位于单个目录下的一个或多个 .go 源代码文件组成，
目录定义包的作用。每个源文件都以一条`package`声明语句开始，这个例子里就是`package main`，表示该文件属于哪个包，紧跟着一系列导入(import)的包，
之后是存储在这个文件里的程序语句。

`main`包比较特殊。它定义了一个独立可执行的程序，而不是一个库。在`main`里的`main`函数也很特殊，它是整个程序执行时的入口（译注：C 系语言差不多
都这样）。`main`函数所做的事情就是程序做的。

必须告诉编译器源文件需要哪些包，这就是跟随在`package`声明后面的`import`声明扮演的角色。

`import`声明必须跟在文件的`package`声明之后。随后，则是组成程序的函数、变量、常量、类型的声明语句（分别由关键字`func`，`var`，`const`，`type`定义）。

Go 语言不需要在语句或者声明的末尾添加分号，除非一行上有多条语句。实际上，编译器会主动把特定符号后的换行符转换为分号，因此换行符添加的位置会影响 Go 代
码的正确解析（译注：比如行末是标识符、整数、浮点数、虚数、字符或字符串文字、关键字`break`、`continue`、`fallthrough`或`return`中的一个、运算符和分隔
符`++`、`--`、`)`、`]`或`}`中的一个）。举个例子，函数的左括号（必须和`func`函数声明在同一行上，且位于末尾，不能独占一行，而在表达式`x+y`中，可以`+`后换行，
不能在`+`前换行（译注：以`+`结尾的话不会被插入分号分隔符，但是以`x`结尾的话则会被分号分隔，从而导致编译错误）。

Go 语言在代码格式上采取了很强硬的态度。`gofmt`工具把代码格式化为标准格式，并且`go`工具中的`fmt`子命令会对指定包（如无指定默认为当前目录）中所有`.go`源文件
应用`gofmt`命令。


> 命令行参数

大多数的程序都是处理输入，产生输出；这也正是"计算"的定义。

`os`包以跨平台的方式 ，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从`os`包的 Args 变量获取；`os`包外部使用 os.Args 访问该变量。

下面是 Unix 里 echo 命令的一份实现，echo 把它的命令行参数打印成一行。程导入了两个包，用括号把它们括起来写成列表形式，而没有分开写成独立的 
import 声明。两种形式都合法，列表形式习惯上用得多。包导入顺序并不重要；gofmt 工具格式化时按照字母顺序对包名排序。

**echo1.go**

```
// Echo1 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i ++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}
```

注释语句以 // 开头。

var 声明定义了两个 string 类型的变量 s 和 sep。变量会在声明时直接初始化。如果变量没有显式初始化，则被隐式地赋予其类型的零值(zero value)，
数值类型是 0，字符串类型是空字符中""。这个例子里，声明把 s 和 sep 隐式地初始化成空字符串。对数值类型，Go 语言提供了常规的数值和逻辑运算符。
而对 string 类型，+ 运算符连接字符串。

循环索引变量 i 在 for 循环的第一部分中定义。符号 := 是短变量声明(short variable declaration)的一部分，这是定义一个或多个变量并根据它们
的初始值为这些变量赋予适当类型的语句。自增语句 i++ 给 i 加 1。

Go 语言只有 for 循环这一种循环语句。for 循环有多种形式，其中一种如下所示：

```
for initialization; condition; post {
    // zero or more statements
}
```

for 循环三个部分不需括号包围。大括号强制要求，左大括号必须和 post 语句在同一行。for 循环的这三个部分每个都可以省略，如果省略 initialization 
和 post，分号也可以省略：

```
// a traditional "while" loop
for condition {
    // ...
}
```

如果连 condition 也省略了，像下面这样：

```
// a traditional infinite loop
for {
    // ...
}
```

这就变成一个无限循环，尽管如此，还可以用其他方式终止循环，如一条 break 或 return 语句。

for 循环的另一种形式，在某种数据类型的区间(range)上遍历，如字符串或切片。

**echo2.go**

```
// Echo2 prints its command-line arguments
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
```

每次循环迭代，range 产生一对值；索引以及在该索引处的元素值。这个例子不需要索引，但 range 的语法要求，要处理元素，必须处理索引。一种思路是把索
引赋值给一个临时变量，如 temp，然后忽略它的值，但 Go 语言不允许使用无用的局部变量(local variables)，因为这会导致编译错误。

Go 语言中这种情况的解决方法是用 **空标识符(blank identifier)**，即`_`（也就是下划线）。空标识符可用于任何语法需要变量名但程序逻辑不需要的
时候，例如，在循环里，丢弃不需要的循环索引，保留元素值。大多数 Go 程序员都会像上面这样使用 range 和 _ 写 echo 程序，因为隐式地而非显式地索
引 os.Args，容易写对。

echo 的这个版本使用一条短变量声明来声明并安始化 s 和 seps，也可以将这两个变量分开声明，声明一个变量有好几种方式，下面这些都等价：

```
s := ""
var s string
var s = ""
var s string = ""
```

用哪种不用哪种，为什么呢？第一种形式，是一条短变量声明，最简洁，但只能用在函数内部，而不能用于包变量。第二种形式依赖于字符中的默认初始化零值机
制，被初始化为 ""。第三种形式用得很少，除非同时声明多个变量。第四种形式显式地标明变量的类型，当变量类型与初始值类型相同时，类型冗余，但如果两
者类型不同，变量类型就必须了。实践中一般使用前两种形式中的某个，初始值重要的话就显式地指定变量的类型，否则使用隐式初始化。

如前文所述，每次循环失代字符串 s 的内容都会更新。+= 连接原字符串、空格和下个参数，产生新字符串，并把它赋值给 s。s 原来的内容已经不再使用，将
在适当时机对它进行垃圾回收。

如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用 strings 包的 Join 函数：

**echo3.go**

```
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
```

>> 查找重复的行

对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：一个处理输入
的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出。

**dup1.go**

```
// Dup1 prints the text of each line that appears more than 
// once int the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```

正如 for 循环一样，if 语句条件两边也不加括号，但是主体部分需要加。

map 存储了键/值(key/value)的集。键值可以是任意类型。该例中的键是字符串，值是整数。内置函数 make 创建 map。（译注：从功能和实现上说，
Go 的 map 类似于 Java 中的 HashMap，Python 中的 dict，Lua 中的 table，通常使用 hash 实现。

每次 dup 读取一行输入，该行被当做 map，其对应的值递增。`counts[input.Text()]++`语句等价于下面两句：

```
line := input.Text()
counts[line] = counts[line] + 1
```

为了打印结果，使用了基于 range 的循环，并在 counts 这个 map 上迭代。map 的迭代顺序并不确定，从实践来看，该顺序随机，每次运行都会变化。
这种设计是有意为之的，因为能防止程序依赖特定遍历顺序，而这是无法保证的（译注：具体可以参见这里 
[https://stackoverflow.com/questions/11853396/google-go-lang-assignment-order](https://stackoverflow.com/questions/11853396/google-go-lang-assignment-order)）

继续来看 bufio 包，它使处理输入和输出方便又高效。Scanner 类型是该包最有用的特性之一，它读取输入并将其拆成行或单词；通常是处理行形式的
输入最简单的方法。

程序使用短变量声明创建 bufio.Scanner 类型的变量 input。

```
input := bufio.NewScanner(os.Stdin)
```

该变量从程序的标准输入中读取内容。每次调用 input.Scan()，即读入下一行，并移除行末的换行符；读取的内容可以调用 input.Text() 得到。
Scan 函数在读到一行时返回 true，不再有输入时返回 false。

类似于 C 或其它语言里的 printf 函数，fmt.Printf 函数对一些表达式产生格式化输出。

**dup2.go**

```
// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%s\t", line)
            // fmt.Printf("%d\t%s\n", n, line)
        } else {
            fmt.Println("Hello world!")
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        fmt.Printf("%s\n", input.Text())
        counts[input.Text()] ++
    }
}
```

os.Open 函数返回两个值。第一个值是被打开的文件(*os.File)，其后被 Scanner 读取。第二个值是内置 error 类型的值。如果 err 等于内置值 nil 
（译注：相当于其它语言里的 NULL），那么文件被成功打开；相反的话，如果 err 的值不是 nil，说明打开文件时出错了。这种情况下，错误值描述了所遇到
的问题。该例的错误处理非常简单，只是使用 Fprintf 与表示任意类型默认格式值的动词 %v，向标准错误流打印一条信息。

注意 countLines 函数在其声明前被调用。函数和包级别的变量(package-level entities)可以任意顺序声明，并不影响其被调用。（译注：最后还是遵循
一定的规范）

map 是一个由 make 函数创建的数据结构的引用。map 作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），被调用函数对 map 
底层数据结构的任何修改，调用者函数都可以通过持有的 map 引用看到。在该例中，countLines 函数向 counts 插入的值，也会被 main 函数看到。（译
注：类似于 C++ 里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）

dup 的前两个版本以“流”模式读取输入，并根据需要拆分成多个行。理论上，这些程序可以处理任意数量的输入数据。还有另一个方法，就是一口气把全部输入
数据读到内存中，一次分割为多行，然后处理它们。下面这个版本，dup3，就是这么操作的。这个例子引入了 ReadFile 函数（来自于 io/ioutil 包），其
读取指定文件的全部内容，strings.Split 函数把字符串分割成子串的切片。（Split 的作用与前文提到的 strings.Join 相反。）

**dup3**

```
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line] ++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
		}
	}
}
```

ReadFile 函数返回一个字节切片(byte slice)，必须把它转换为 string，才能用 strings.Split 分割。

实现上，bufio.Scanner、ioutil.ReadFile 和 ioutil.WriteFile 都使用 *os.File 的 Read 和 Write 方法，但是，一般很少需要直接调用那些
低级(lower-level)函数。高级(higher-level)函数，像 bufio 和 io/ioutil 包中所提供的那些用起来要容易点。


> GIF 动画

下面的程序会演示 Go 语言标准库里的 image 这个 package 的用法，将用这个包来生成一系列的 bit-mapped 图，然后将这些图片编码为一个 GIF 动画。
这个图形名字叫利萨如图形(Lissajous figures)。这段代码使用了一些新的结构，包括 const 声明，struct 结构体类型，复合声明。

**lissajous**

```
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	WhiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thank to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5      // number of complete x oscillator revolutions
		res     = 0.001  // angular resolution
		size    = 100    // image canvas covers [-size..+size]
		nframes = 64     // number of animation frames
		delay   = 8      // delay between frames in 10 ms units
	)

	freq := rand.Float64() * 3.0  // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0  // phase difference
	for i := 0; i < nframes; i ++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)  // NOTE: ignoring encoding errors
}
```
