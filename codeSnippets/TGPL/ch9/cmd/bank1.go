package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch9/bank1"
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		bank1.Deposit(200)
		fmt.Println(bank1.Balance())
		w.Done()
	}()

	go func() {
		bank1.Deposit(100)
	}()

	w.Wait()
}
