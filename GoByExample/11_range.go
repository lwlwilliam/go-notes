/*
range 迭代各种各样的数据结构
 */
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// 使用 range 对 slice 中的元素求和。对于数组也可采用这种方法
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range 在数组和 slice 中提供对每项的索引和值的访问。上面不需要索引，所以使用 _ 空白标识符来忽略它
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// range 在 map 中迭代键值对
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 也可以只迭代键
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 在字符串中迭代 unicode 码点(code point)。第一个返回值是字符的起始字节位置，第二个是字符本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
