package main

import (
	"fmt"
)

func main() {
	// score 为 [1, 100] 之间的整数
	var score int = 69

	// 先用 if...else
	if score >= 90 && score <= 100 {
		fmt.Println("优秀")
	} else if score >= 80 && score < 90 {
		fmt.Println("良好")
	} else if score >= 70 && score < 80 {
		fmt.Println("一般")
	} else if score >= 60 && score < 70 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// 用 switch
	switch score / 10 {
	case 10:
	case 9:
		fmt.Println("优秀")
	case 8:
		fmt.Println("良好")
	case 7:
		fmt.Println("一般")
	case 6:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
