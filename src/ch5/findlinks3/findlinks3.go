package main

import (
	ch5 "ch5/ch5base"
	"os"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	ch5.BreadthFirst(ch5.Crawl, os.Args[1:])
}
