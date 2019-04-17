package bank5

import (
	"sync"
)

var (
	mu      sync.Mutex // guards balance
	murw	sync.RWMutex
	balance int
)

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	murw.RLock()
	defer murw.RUnlock()
	return balance
}

// This function requires that the lock be held.
func deposit(amount int) {
	balance += amount
}