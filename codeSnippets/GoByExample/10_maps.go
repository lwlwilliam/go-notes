/*
map 是 Go 内置关联数据类型（在一些其他的语言中被称为 哈希(hash)或者字典(dict)。
 */
package main

import "fmt"

func main() {
	// 要创建一个空 map，需要使用内建的 make：make(map[key-type]val-type)。
	m := make(map[string]int)

	// 使用 map[key] = val 语法来设置键值对
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// 使用 map[key] 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 调用内建的 len 获取键值对数目
	fmt.Println("len:", len(m))

	// 内建的 delete 移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))


	// 当从一个 map 中取值时，可选的第二返回值指示这个键是否在这个 map 中。这可以用来消除键不存在和键有零值，
	// 像 0 或者 "" 而产生的歧义。这里不需要值，所以用 _ 空白标识符忽略。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 也可以通过这种语法在同一行声明和初始化一个新的 map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
