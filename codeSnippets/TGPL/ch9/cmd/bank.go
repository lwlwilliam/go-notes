package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch9/bank"
	"fmt"
	"sync"
)

func main()  {
	// Alice:
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		bank.Deposit(200)
		fmt.Println(bank.Balance())

		w.Done()
	}()

	// Bob
	go bank.Deposit(100)

	w.Wait()
}
