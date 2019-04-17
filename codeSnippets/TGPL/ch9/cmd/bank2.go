package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch9/bank2"
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		bank2.Deposit(200)
		fmt.Println(bank2.Balance())
		w.Done()
	}()

	go func() {
		bank2.Deposit(100)
	}()

	w.Wait()
}