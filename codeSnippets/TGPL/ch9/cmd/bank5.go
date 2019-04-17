package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch9/bank5"
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		bank5.Deposit(200)
		fmt.Println(bank5.Balance())
		w.Done()
	}()

	go func() {
		bank5.Deposit(200)
	}()

	w.Wait()
}
