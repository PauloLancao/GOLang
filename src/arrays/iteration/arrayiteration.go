package main

import "fmt"

func naturalOrder(arr []int) {
	for number := range arr {
		fmt.Println("Numbers are:", arr[number])
	}
}

func reverseOrder(arr []int) {
	for number := range arr {
		fmt.Println("Numbers are:", arr[len(arr)-1-number])
	}
}

func naturalOrderSumOdd(arr []int) {
	var sum int = 0
	for number := range arr {
		if arr[number]%2 != 0 {
			sum += arr[number]
		}
	}

	fmt.Println("\nSum odd numbers: ", sum)
}

func reverseOrderSumOdd(arr []int) {
	var sum int = 0
	for number := range arr {
		if arr[len(arr)-1-number]%2 != 0 {
			sum += arr[len(arr)-1-number]
		}
	}

	fmt.Println("\nReverse Sum odd numbers: ", sum)
}

func naturalOrderSumEven(arr []int) {
	var sum int = 0
	for number := range arr {
		if arr[number]%2 == 0 {
			sum += arr[number]
		}
	}

	fmt.Println("\nSum even numbers: ", sum)
}

func reverseOrderSumEven(arr []int) {
	var sum int = 0
	for number := range arr {
		if arr[len(arr)-1-number]%2 == 0 {
			sum += arr[len(arr)-1-number]
		}
	}

	fmt.Println("\nReverse Sum even numbers: ", sum)
}

// Create program that initialises an array with the integer values 1 to 10.
// Display the array content in sequential order 1 to 10 and then from 10 to 1.
// Count even numbers and then odds numbers in increasing and decreasing sequential order.
// Display the count sequences to screen. [Arrays][Slices][For loops][Unit Testing]
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("\nArray order:")
	naturalOrder(arr)

	fmt.Println("\nArray reverse order:")
	reverseOrder(arr)

	naturalOrderSumOdd(arr)
	reverseOrderSumOdd(arr)

	naturalOrderSumEven(arr)
	reverseOrderSumEven(arr)
}
