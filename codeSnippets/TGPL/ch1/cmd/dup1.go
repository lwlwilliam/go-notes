// dup1 用来打印标准输入中出现次数超过 1 次的每行文本
// Scanner 类型是 bufio 最有用的特性之一，它读取输入并将其拆成行或单词；
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()] ++
		//fmt.Println(input.Bytes())
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
