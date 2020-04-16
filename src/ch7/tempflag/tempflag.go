package main

import (
	"ch7/tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

// ./tempflag -temp -18C
func main() {
	flag.Parse()
	fmt.Println(*temp)
}
