package main

import (
	"fmt"

	"ch6/geometry"
)

// A Path is a journey connecting the points with straight lines.
type Path []geometry.Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}
	fmt.Println(geometry.Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))           // "5", method call

	perim := Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	fmt.Println(perim.Distance()) // "12"
}
