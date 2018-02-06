package main

import (
	"fmt"
)

type Phone struct {
	price int
	color string
}

func (phone Phone) ring() {
	fmt.Println("Phone is ringing...")
}

type IPhone struct {
	phone Phone
	model string
}

// 第二种写法
type Huawei struct {
	Phone
	model string
}

func main() {
	var p IPhone
	p.phone.price = 5000
	p.phone.color = "Black"
	p.model = "iPhone 5"

	fmt.Println("I have a iPhone: ")
	fmt.Println("Price: ", p.phone.price)
	fmt.Println("Color: ", p.phone.color)
	fmt.Println("Model: ", p.model)
	p.phone.ring()


	var h Huawei
	h.price = 5000
	h.color = "Black"
	h.model = "Rongyao V9"

	fmt.Println("I have a Harwei: ")
	fmt.Println("Price: ", h.price)
	fmt.Println("Color: ", h.color)
	fmt.Println("Model: ", h.model)
	h.ring()
}
