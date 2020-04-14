package main

import "fmt"

func main() {
	ages := make(map[string]int) // mapping from strings to ints
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)
	fmt.Println(ages["alice"]) // "31"

	delete(ages, "alice") // remove element ages["alice"]
	fmt.Println(ages)

	ages["bob"]++
	fmt.Println(ages)
}
