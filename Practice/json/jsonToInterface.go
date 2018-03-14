package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	b := []byte(`{"Name":"Wednesday", "Age":6, "Parents":["Gomes", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("err")
	}

/*
	f = map[string]interface{}{
		"Name": "Wednesday",
		"Age": 6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}
*/
	
	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
