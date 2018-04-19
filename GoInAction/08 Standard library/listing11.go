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

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
			log.Ldate | log.Ltime | log.Lshortfile)
	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate | log.Ltime | log.Lshortfile)
	Trace = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate | log.Ltime | log.Lshortfile)
	Trace = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	//Warning.Println("There is something you need to know about")
	//Error.Println("Something has failed")
}
