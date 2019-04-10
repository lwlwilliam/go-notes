// 深度优先遍历
package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/toposort"
)

func main() {
	for i, course := range toposort.TopoSort(toposort.Prereqs) {
		fmt.Println(i, course)
	}
}