package main

import (
	"fmt"
	"time"
)

// Employee exported
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	fmt.Println(dilbert)
	dilbert.Salary -= 5000
	fmt.Println(dilbert)

	dilbert := Employee{11, "Dilbert", "Dilbert address", time.Now(), "Senior Manager", 99999, 555}
	fmt.Println(dilbert)
}
