/// Nonempty 是 in-place 算法的一个示例
package main

import "fmt"

// nonempty 返回只保留非空字符串的 slice
// 底层数组会在调用时被修改
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i ++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main()  {
	strings := []string{
		"aaaaa",
		"bbbbb",
		"",
		"ccccc",
		"ddddd",
		"",
	}

	// 不要使用这种写法，不注意的话可能会导致意外出现
	//res := nonempty(strings)
	//res = nonempty(strings)
	//fmt.Printf("%v: %d\n", res, len(res))

	strings = nonempty(strings)
	strings = nonempty(strings)
	strings = nonempty(strings)
	fmt.Printf("%v: %d\n", strings, len(strings))
}
