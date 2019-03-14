package main

import "fmt"

func main() {
	// 基本例子
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 可以不用 else
	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 在条件语句之前可以有一个声明语句；在这里声明的变量可以在所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	/*
	Go 中条件语句不需要圆括号，但是需要加上花括号
	Go 中没有三目运算符
	 */
}
