package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	fmt.Print("Start before counter")

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	fmt.Print("End after counter")
	fmt.Print("Start before squarer")

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	fmt.Print("End after squarer")
	fmt.Print("Start before printer")

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}

	fmt.Print("end after printer")
}
