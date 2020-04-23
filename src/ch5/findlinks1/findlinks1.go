// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	ch5 "ch5/ch5base"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range ch5.Visit(nil, doc) {
		fmt.Println(link)
	}
}
