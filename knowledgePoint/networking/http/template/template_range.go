package main

import (
	"fmt"
	"text/template"
	"os"
)

func main() {
	s := []int{1, 2, 3, 4}
	t := template.New("test")
	// range 用于迭代集合：pipeline 的值必须是 array, slice 或 map 类型。
	// {{range pipeline}} T1 {{else}} T0 {{end}}
	// 如果 pipeline 的长度为 0，dot 不可访问，T0 执行；否则 dot 被设置为 array, slice 或 map 的元素，T1 被执行。
	t, _ = t.Parse("{{range .}}{{.}}\n{{end}}")
	if err := t.Execute(os.Stdout, s); err != nil {
		fmt.Println("There is an error:", err.Error())
	}
}
