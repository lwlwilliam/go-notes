package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	// 基本的 switch
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	// 在同一个 case 中，可以用逗号分隔多个表达式。default 分支可选
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
	}

	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式。
	// case 表达式也可以不使用常量
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("It's before noon")
		default:
			fmt.Println("It's after noon")
	}

	// 类型开关(type switch)比较类型而非值。可以用来发现接口值的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
			case bool:
				fmt.Println("I'm a bool")
			case int:
				fmt.Println("I'm an int")
			default:
				fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
