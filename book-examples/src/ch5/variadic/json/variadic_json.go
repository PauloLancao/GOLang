package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	Name   string
	Author string
}

func main() {
	book := book{"C++ programming language", "Bjarne Stroutsrup"}
	res, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}
