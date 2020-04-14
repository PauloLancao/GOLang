package main

import (
	"fmt"
	"log"
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

func rotate(s []int, idxSlice int, left bool) []int {
	var size = len(s)
	var res = make([]int, size)

	if idxSlice > size {
		log.Fatal("IdxSlice out of bounds")
	} else {
		if left {
			// rotate left
			res = append(s[idxSlice:], s[:idxSlice]...)
		} else {
			// rotate right
			res = append(s[size-idxSlice:], s[:size-idxSlice]...)
		}
	}
	return res
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

	r := []int{0, 1, 2, 3, 4, 5, 8, 9}
	left := false
	idx := 2
	fmt.Printf("Rotate array original \nLEFT -> %t: \nIdx -> %d \n%v \nresult: \n%v",
		left, idx, r, rotate(r, idx, left))

}
