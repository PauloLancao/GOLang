package main

import "fmt"

func main() {
	fmt.Println(fib(10))
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
