package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

var cities = []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

const fileName string = "cities.txt"

func createFile() {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	bs, err := stringArrToByteArr()
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	l, err := f.Write(bs)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readFile() []byte {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	return content
}

func stringArrToByteArr() ([]byte, error) {
	return json.Marshal(cities)
}

func byteArrToStringArr(fileBytes []byte) []string {
	var resultArr []string

	err := json.Unmarshal(fileBytes, &resultArr)

	if err != nil {
		fmt.Println(err)
	}

	return resultArr
}

// Write a program copies the following list of cities to a new file -
// "Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi".
// Read a list of cities from the newly created file and display the list of cities in alphabetical order)
func main() {

	createFile()

	resultArr := byteArrToStringArr(readFile())

	fmt.Println("Not Ordered: ", resultArr)

	sort.Strings(resultArr)

	fmt.Println("Ordered: ", resultArr)
}
