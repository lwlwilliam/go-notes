/*
有时，需要生成其他的非 Go 进程。例如，该网站的语法高亮是通过在 Go 程序中生成一个
pygmentize 来实现的。
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// 先从一个简单的命令开始，没有参数或者输入，仅打印一些信息到标准输出流
	// exec.Command 函数帮助创建一个表示这个外部进程的对象
	dateCmd := exec.Command("date")

	// Output 是另一个帮助处理运行一个命令的常见情况的函数，它等待命令运行完成，并收集命令的输出。
	// 如果没有出错，dateOut 将获取到日期信息的字节
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 从外部进程的 stdin 输入数据并从 stdout 收集结果
	grepCmd := exec.Command("grep", "hello")

	// 这里明确的获取输入/输出管道，运行这个进程，写入一些输入信息，读取输出结果，最后等待程序运行结束
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll((grepOut))
	grepCmd.Wait()

	// 以上例子中，忽略了错误检测，但是可以使用 if err != nil 的方式来进行错误检查
	// 注意，当需要提供一个明确的命令和参数数组来生成命令，和能够只需要提供一个命令行字符相比，
	// 想使用通过一个字符串生成一个完整的命令，那么可以使用 bash 命令的 -c 选项
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}