// 这个示例程序展示如何创建定制的日志记录器
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace	*log.Logger	// 记录所有日志
	Info	*log.Logger	// 重要的消息
	Warning	*log.Logger	// 需要注意的信息
	Error	*log.Logger	// 非常严重的问题
)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}


	/*
	log 包的 New 函数，创建并正确初始化一个 Logger 类型的值。New 会返回新创建的值的地址。
	第一个参数设置日志数据将被写入的目的地，第二个参数会在生成的每行日志的最开始出现
	第三个参数定义日志记录包含哪些属性
	*/

	// ioutil 包里的 Discard 变量作为目的地，Discard 为 io.Writer 接口类型，
	// 并被给定了一个 devNULL 类型的值 0，因此会忽略所有写入这一变量的数据
	Trace = log.New(ioutil.Discard,
		"TRACE: ",
			log.Ldate | log.Ltime | log.Lshortfile)

	// Info 和 Warning 都使用 stdout 作为日志输出。
	// Stdin、Stdout 和 Stderr 都被声明为 File 类型指针
	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate | log.Ltime | log.Lshortfile)

	// 第一个参数来自一个特殊的函数，这个函数调用会返回一个 io.Writer 接口类型的值，
	// 这个值包含之前打开的文件 file，以及 stderr。MultiWriter 是一个变参函数，
	// 可以接受任意个实现了 io.Writer 接口的值。返回的 io.Writer 接口类型的值会把
	// 所有传入的 io.Writer 的值绑在一起。当对这个返回值进行写入时，会向所有绑在一起
	// 的 io.Writer 值做写入。这让类似 log.New 这样的函数可以同时向多个 Writer 做输出。
	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
