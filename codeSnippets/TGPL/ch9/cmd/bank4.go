package main

import (
	"fmt"

	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch9/bank4"
	"time"
)

func main() {
	go func() {
		bank4.Deposit(200)
		fmt.Println(bank4.Balance())
	}()

	go func() {
		bank4.Deposit(200)
	}()

	time.Sleep(1 * time.Second)

	bank4.Withdraw(100)
	fmt.Println(bank4.Balance(), "...")
}
