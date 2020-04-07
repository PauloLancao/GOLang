package main

import (
	"fmt"
)

// RemoveIndexAlt version 1
func RemoveIndexAlt(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// RemoveIndex version 0
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Arr full ", arr)

	arr = RemoveIndex(arr, 3)

	fmt.Println("Arr remove index 3 ", arr)

	var arr1 = RemoveIndexAlt(arr, 1)

	fmt.Println("Arr1 remove index 1 ", arr1)
	fmt.Println("Arr full ", arr)

	sum := 0
	for _, num := range arr {
		sum += num
	}

	fmt.Println("sum:", sum)
}
