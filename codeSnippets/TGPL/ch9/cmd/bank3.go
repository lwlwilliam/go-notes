package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch9/bank3"
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		bank3.Deposit(200)
		fmt.Println(bank3.Balance())
		w.Done()
	}()

	go func() {
		bank3.Deposit(100)
	}()

	//bank3.Withdraw(100)
	//fmt.Println(bank3.Balance())

	w.Wait()
}
