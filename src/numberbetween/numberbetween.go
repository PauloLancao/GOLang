package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nWrite number between 1 - 10: ")
		number, _ := reader.ReadString('\n')

		fmt.Println("Number:", number)

		i, err := strconv.Atoi(strings.Replace(number, "\r\n", "", -1))

		if err != nil {
			fmt.Printf("Incorrect number: %d, try again...", i)
		} else {
			if i > 0 && i < 11 {
				fmt.Println("Correct number")
				break
			} else {
				fmt.Printf("Incorrect number: %d, try again...", i)
			}
		}
	}
}
