package practice

import (
	"encoding/json"
	"fmt"
)

type A struct {
	D1 int
	D2 int
}

type I struct {
	As interface{}
}

func main() {
	s := "128"
	n := 0
	for i := 0; i < len(s); i ++ {
		n *= 10
		n += (int(s[i]) - int('0'))
	}

	fmt.Printf("%d\t%T\n", n, n)

	// ###################################################################
	fmt.Printf("###################################################################\n")

	// 输出结果？
	fmt.Printf("func slice=%+v\n", slice())  // [1 -1 -2 -3 -1 -2 -3]
	fmt.Printf("func slice1=%+v\n", slice1())  // [1 -1 -2 -3 -4 2 3 4]

	// ###################################################################
	fmt.Printf("###################################################################\n")

	data := []byte(`{"As":[{"D1":1, "D2":2},{"D1":-1, "D2":-2}]}`)

	i := new(I)
	i.As = make([]A, 2)
	// i.As = new([]A)
	json.Unmarshal(data, i)

	//var vtype = reflect.TypeOf(i.As)
	//fmt.Println(vtype)

	test, ok := i.As.([]interface{})
	//test, ok := i.As.(*[]A)

	// 输出结果？
	fmt.Printf("test=%v\tok=%v\n", test, ok)
}

func slice() []int {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3}

	// 首先 s1 变为 [1 -1 -2 -3]，这时候容量还是足够的，但是 s1 底层数组的值被修改；
	// 接着整个 s1 添加元素时，容量不够，需要新建底层数组，返回的是新 slice [1 -1 -2 -3 -1 -2 -3]
	x := append(s1[:1], s2...)
	y := append(x, s1[1:]...)
	fmt.Println("slice")
	fmt.Println("len(x):", len(x), "; cap(x)", cap(x))
	fmt.Println("len(y):", len(y), "; cap(y)", cap(y))
	return y
	//return append(append(s1[:1], s2...), s1[1:]...)
}

func slice1() []int {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3, -4}

	// 首先 s1[:1] 容量不足以添加 s2 所有元素，这时候就需要新建底层数组 [1 -1 -2 -3 -4]；这时候容量为 8；而原来 s1 的底层数组没变
	// 接着把 s1[1:] 添加一下还是容量足够的 [1 -1 -2 -3 -4 2 3 4]
	x := append(s1[:1], s2...)
	y := append(x, s1[1:]...)
	fmt.Println("slice1")
	fmt.Println("len(x):", len(x), "; cap(x)", cap(x))
	fmt.Println("len(y):", len(y), "; cap(y)", cap(y))
	return y
	//return append(append(s1[:1], s2...), s1[1:]...)
}
