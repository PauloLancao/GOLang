package main

import (
	"fmt"

	"ch6/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}
	fmt.Println(geometry.Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))           // "5", method call
}
