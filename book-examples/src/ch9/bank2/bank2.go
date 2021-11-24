package bank2

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

// Deposit func
func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

// Balance func
func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
