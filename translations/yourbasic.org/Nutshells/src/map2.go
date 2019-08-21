package main

import "fmt"

func main() {
	m := map[string]float64{
		"e": 2.71828,
		"pi": 3.1416,
	}

	for key, value := range m {  // 无序
		fmt.Println(key, value)
	}
}
