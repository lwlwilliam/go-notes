// 这个示例程序展示如何使用最基本的 log 包
// log 包有一个很方便的地方就是，这些日志记录器是多 goroutine 安全的。这意味着在多个 goroutine 可以同时调用来自同一个日志记录器
// 的这些函数中，而不会有彼此间的写冲突。标准日志记录器具有这一性质，用户定制的日志记录器也应该满足这一性质。
package main

import (
	"log"
)

// init 函数会在运行 main 之前作为程序初始化的一部分执行。通常程序会在 init 里配置日志参数，这个程序一开始就能使用 log 包进行正确的输出
func init() {
	// 设置了一个字符串，作为每个日志项的前缀，这个字符串应该是能让用户从一般的程序输出中分辨出日志的字符串。传统上这个字符串的字符会全部大写
	log.SetPrefix("TRACE: ")

	// 有几个和 log 包相关联的标志，这些标志用来控制可以写到每个日志项的其他信息
	/*
	const {
		// 将下面的位使用或运算符连接在一起，可以控制要输出的信息。
		// 没有办法控制这些信息出现的顺序或者打印的格式。这些项后面会有一个冒号

		// 日期：2009/01/23
		// itoa 在常量声明区里有特殊作用。这个关键字让编译器为每个常量复制相同的表达式，直到声明区结束，或者遇到一个新的赋值语句
		// 它的另一个功能是，iota 的初始值为 0，之后 iota 的值在每次处理为常量后，都会自增 1
		Ldate = 1 << itoa			// 1 << 0 = 000000001 = 1

		// 时间：01:23:23
		Ltime						// 1 << 1 = 000000010 = 2

		// 毫秒级时间：01:23:23.123123。该设置会覆盖 Ltime 标志
		Lmicroseconds				// 1 << 2 = 000000100 = 4

		// 完整路径的文件名和行号：/a/b/c/d.go:23
		Llongfile					// 1 << 3 = 000001000 = 8

		// 最终的文件名元素和行号：d.go:23。覆盖 Llongfile
		Lshortfile					// 1 << 4 = 000010000 = 16

		// 标准日志记录器的初始值
		// 这个常量展示了如何使用以上标志
		LstdFlags = Ldate | Ltime
	}
	 */

	 // 设置日志标志
	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	// Println 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用 Println() 之后会接着调用 os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln 在调用 Println() 之后会接着调用 panic()
	// 除非程序执行 recover 函数，否则会导致程序打印调用栈后终止。Print 系列函数是写日志消息的标准方法
	log.Panicln("panic message")
}
