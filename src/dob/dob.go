package main

import (
	"fmt"
	"time"

	"github.com/bearbin/go-age"
)

func calcDob() int {
	return age.Age(time.Date(1976, 3, 3, 0, 0, 0, 0, time.UTC))
}

func main() {
	var dob = calcDob()
	fmt.Println(dob)
}
