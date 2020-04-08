package main

import "fmt"

// Create program that initialises an array with the integer values 1 to 10.
// Display the array content in sequential order 1 to 10 and then from 10 to 1.
// Count even numbers and then odds numbers in increasing and decreasing sequential order.
// Display the count sequences to screen. [Arrays][Slices][For loops][Unit Testing]
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("\nArray order:")

	for number := range arr {
		fmt.Println("Numbers are:", arr[number])
	}

	fmt.Println("\nArray reverse order:")

	for number := range arr {
		fmt.Println("Numbers are:", arr[len(arr)-1-number])
	}

}
