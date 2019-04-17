package bank3

import "sync"

var (
	mu sync.Mutex	// guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

// 错误演示，以下做法是不对的
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false // insufficient funds
	}
	return true
}