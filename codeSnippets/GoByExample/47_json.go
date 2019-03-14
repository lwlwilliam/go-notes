/*
Go 提供内置的 JSON 编解码支持，包括内置或者自定义类型与 JSON 数据之间的转化
*/
package main

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
	Page   int
	Fruits []string
	test   int
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	/*
		基本类型到 JSON 字符串的编码过程
	*/
	bolB, _ := json.Marshal(true)
	fmt.Println("1.", string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println("2.", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("3.", string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println("4.", string(strB))

	/*
		切片和 map 编码成 JSON 数组和对象的例子
	*/
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("5.", string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("6.", string(mapB))

	/*
		JSON 包可以自动编码自定义类型。编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的 key 值
	*/
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		test:   1,
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("7.", string(res1B))

	// 可以给结构体字段声明标签来自定义编码的 JSON 数据键名。
	// 在上面 Response2 的定义可以作为标签例子
	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("8.", string(res2B))


	/*
	解码 JSON 数据为 Go 值
	 */
	byt := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)

	// 提供一个 JSON 包可以存放解码数据的变量
	// 这里的 map[string]interface{} 将保存一个 string 为键，任意值为值的 map
	var dat map[string]interface{}

	// 实际解码和相关的错误检查
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("9.", dat)

	// 使用解码 map 中的值，对他们进行适当的类型转换。
	num, ok := dat["num"].(float64)
	fmt.Println("10.", num, ok)

	// 访问嵌套的值需要一系列转化
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("11.", str1)

	//var test []interface{}
	//test = make([]interface{}, 3)
	//fmt.Println(test)
	//
	//var haha []string
	//haha = make([]string, 3)
	//fmt.Println(haha)

	// 也可以解码 JSON 值到自定义类型。这个功能的好处就是可以为程序带来额外的类型安全加强，
	// 并且消除在访问数据时的类型断言
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), res)
	fmt.Println("12.", res)
	//fmt.Printf("%#x\n", res)
	fmt.Println("13.", res.Fruits[0])

	// 上例中，经常使用 byte 和 string 作为使用标准输出时数据和 JSON 表示之间的中间值。
	// 也可以和 os.Stdout 一样，直接将 JSON 编码直接输出至 os.Writer 流中，或者作为 HTTP 响应体
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	fmt.Printf("14. ")
	enc.Encode(d)

	//o := Response2{
	//	Page: 1,
	//	Fruits: []string{"apple", "banana"},
	//}
	//
	//enc.Encode(o)
}
