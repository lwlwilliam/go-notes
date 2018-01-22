// 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
