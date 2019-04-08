// 读取多行输入，但是只打印第一次出现的行。
// dedup 程序通过 map 来表示所有的输入行所对应的 set 集合，以确保已经在集合存在的行不会被重复打印。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	seen := make(map[string]bool)	// a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if _, ok := seen[line]; !ok {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}