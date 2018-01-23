// 修改 echo 程序，使其打印每个参数的索引和值，每个一行

package main

import (
	"fmt"
	"os"
	// "strings"
	"strconv"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], "\n"))
	var s, sep string
	for index, value := range os.Args[1:] {
        sep = strconv.Itoa(index)
        s += sep + " : " + value + "\n"
	}
	fmt.Println(s)
}
