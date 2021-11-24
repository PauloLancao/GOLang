package main

import "fmt"

func singleDigits() (int, int, int) {
	return 9, 3, 5
}

func doubleDigits() (int, int, int) {
	return 23, 66, 99
}

func tripleDigits() (int, int, int) {
	return 928, 373, 123
}

func sumAll() int {
	a, b, c := singleDigits()
	d, e, f := doubleDigits()
	g, h, j := tripleDigits()

	return a + b + c + d + e + f + g + h + j
}

func main() {
	res := sumAll()

	fmt.Println(res)
}
