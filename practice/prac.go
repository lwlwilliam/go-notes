package practice

import (
	"fmt"
)

// 1
func c() (i int) {
	i = 0
	defer func() {
		i++
	}()
	return
}

// 0
func c1() int {
	i := 0
	defer func() {
		i++
	}()
	return i
}

func c2() int {
	i := 0
	defer func(i int) {
		i++
	}(i)
	return i
}

func c3() (i int) {
	i = 0
	defer func(i int) {
		i++
	}(i)
	return
}

func main() {
	fmt.Println("func c=", c())
	fmt.Println("func c1=", c1())
	fmt.Println("func c2=", c2())
	fmt.Println("func c3=", c3())
}