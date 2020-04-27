package bank

var balance int

// Deposit func
func Deposit(amount int) { balance = balance + amount }

// Balance func
func Balance() int { return balance }
