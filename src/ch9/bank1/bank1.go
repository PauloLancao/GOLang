package main

import (
	"fmt"
	"time"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

// Deposit func
func Deposit(amount int) { deposits <- amount }

// Balance func
func Balance() int { return <-balances }

// Withdraw func
func Withdraw(amount int) bool {
	balance := <-balances
	if balance-amount < 0 {
		fmt.Printf("Insufficient funds available%d :: amount%d\n", balance, amount)
		return false
	}

	Deposit(-amount)

	return true
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func main() {
	go teller() // start the monitor goroutine

	go Deposit(100)
	go Deposit(200)
	fmt.Printf("Balance:: %d\n", Balance())
	go Deposit(100)
	fmt.Printf("Balance:: %d\n", Balance())
	fmt.Printf("Balance:: %d\n", Balance())
	go Withdraw(100)
	go Withdraw(100)
	fmt.Printf("Balance:: %d\n", Balance())
	fmt.Printf("Balance:: %d\n", Balance())
	fmt.Printf("Balance:: %d\n", Balance())
	go Withdraw(100)
	go Withdraw(100)
	go Withdraw(100)
	go Withdraw(100)
	fmt.Printf("Balance:: %d\n", Balance())
	fmt.Printf("Balance:: %d\n", Balance())
	go Withdraw(100)
	go Withdraw(100)
	go Withdraw(100)
	go Withdraw(100)
	go Withdraw(100)
	fmt.Printf("Balance:: %d\n", Balance())
	fmt.Printf("Balance:: %d\n", Balance())
	go Withdraw(100)

	time.Sleep(10 * time.Second)
}
