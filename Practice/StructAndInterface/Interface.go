package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type Nokia struct {}

func (nokia Nokia) call() {
	fmt.Println("I am Nokia, I can call you.")
}

type IPhone struct {}

func (iphone IPhone) call() {
	fmt.Println("I am IPhone, I can call you.")
}

func main() {
	var phone Phone

	phone = new(Nokia)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

