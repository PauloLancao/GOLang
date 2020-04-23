package main

import (
	ch5 "ch5/ch5base"
	"fmt"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := ch5.FindLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
