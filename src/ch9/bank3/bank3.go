package bank3

import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

// Deposit func
func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

// Balance func
func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
