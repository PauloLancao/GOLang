package main

import (
	"fmt"
	"os"
)

// go run .\osargs.go paulo lancao
func main() {
	fmt.Println(os.Args[0:])
	fmt.Println(os.Args[1:])

	for idx, arg := range os.Args[0:] {
		fmt.Printf("\nIdx: %d, Value: %s", idx, arg)
	}
}
