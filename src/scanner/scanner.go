package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter FirstName: ")
	firstName, _ := reader.ReadString('\n')
	fmt.Print("Enter MiddleName: ")
	middleName, _ := reader.ReadString('\n')
	fmt.Print("Enter LastName: ")
	lastName, _ := reader.ReadString('\n')

	firstName = strings.Replace(firstName, "\r\n", "", -1)
	middleName = strings.Replace(middleName, "\r\n", "", -1)
	lastName = strings.Replace(lastName, "\r\n", "", -1)

	fmt.Printf("%s %s %s", firstName, middleName, lastName)

	// To create dynamic array
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nEnter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}
	}
	// Use collected inputs
	fmt.Println(arr)
}
