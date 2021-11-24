package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Write (q)uit to end scanner")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		input := input.Text()
		if input == "quit" || input == "q" {
			break
		}
		counts[input]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
