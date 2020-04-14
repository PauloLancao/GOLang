package main

import (
	"fmt"
	"log"
)

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
	r := []int{0, 1, 2, 3, 4, 5, 8, 9}
	left := false
	idx := 4
	fmt.Printf("Rotate array original \nLEFT -> %t: \nIdx -> %d \n%v \nresult: \n%v",
		left, idx, r, rotate(r, idx, left))
}
