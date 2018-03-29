/*
常常需要程序在数据集上执行操作，比如选择满足给定条件的所有项，或者将所有的项通过一个自定义函数映射到一个新的集合上。

在某些语言中，会习惯使用泛型。Go 不支持泛型，在 Go 中，当程序或数据类型需要时，通常是通过组合的方式来提供操作函数。

以下是一些 strings 切片的组合函数示例。可以使用这些例子来构建自己的函数。

注意，有时候，直接使用内联组合操作代码会更清晰，而不是创建并调用一个帮助函数。
 */
package main

import "strings"
import "fmt"

// 查找字符串 t 在 字符串切片 vs 中的索引，如果 t 不是 vs 中，返回 -1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 判断 vs 字符串切片中是否包含 t 元素
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// vs 字符串切片是否存在符合 f 函数的元素
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// vs 字符串切片中的元素是否所有都符合 f 函数的条件
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 使用 f 函数对 vs 字符串切片中的字符串进行过滤，以字符串切片的形式返回符合条件的字符串
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 把 vs 切片中的元素转换为 f 函数处理后的结果并返回新的切片
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	// pear 在 strs 中的索引
	fmt.Println(Index(strs, "pear"))  // 2

	// strs 是否包含 grape 元素
	fmt.Println(Include(strs, "grape"))  // false

	// 是否存在 strs 的元素前缀是 p
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))  // true

	// 是否 strs 所有元素的前缀都是 p
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))  // false

	// 获取 strs 中存在 e 的元素
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))  // [peach apple pear]

	// 获取 strs 所有元素转为大写后的值
	fmt.Println(Map(strs, strings.ToUpper))  // [PEACH APPLE PEAR PLUM]
}
