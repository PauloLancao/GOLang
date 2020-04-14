package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reversePtr(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // [5 4 3 2 1 0]
	b := []int{0, 1, 2, 3, 4, 5, 7, 9}
	reversePtr(&b)
	fmt.Println(b) // [9 7 5 4 3 2 1 0]

	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("Rotate left", s)
	// Rotate s right by two positions.
	reverse(s)
	reverse(s[:2])
	reverse(s[2:])

	fmt.Println(s) // "[2 3 4 5 0 1]"

	fmt.Println("Rotate right", s)
	// Rotate s right by two positions.
	reverse(s)
	reverse(s[:2])
	reverse(s[2:])

	fmt.Println(s) // "[2 3 4 5 0 1]"
}
