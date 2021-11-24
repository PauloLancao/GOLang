package main

import (
	"ch9/bank"
	"fmt"
)

func main() {
	// Alice:
	go func() {
		bank.Deposit(200)                // A1
		fmt.Println("=", bank.Balance()) // A2
	}()

	// Bob:
	go bank.Deposit(100) // B
}
