package main

import (
	"fmt"
	"net"
)

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you.")
}

type IPhone struct {

}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you.")
}

func main() {
	var nokia NokiaPhone
	nokia.call()

	var iPhone IPhone
	iPhone.call()
}
