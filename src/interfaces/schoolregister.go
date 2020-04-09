package main

import (
	"fmt"
	"time"

	"github.com/bearbin/go-age"
)

var pupils = []register{}

type register interface {
	record()
}

type schoolRegister struct {
	fullName string
	dob      time.Time
	age      int
}

func (sr schoolRegister) record() {
	fmt.Printf("\nFullName: %s, DOB: %s, Age: %d", sr.fullName, sr.dob, sr.age)
}

func calcDob(dob time.Time) int {
	return age.Age(dob)
}

func addPupilsToRegister(r register) {
	pupils = append(pupils, r)
}

func showAllPupils() {
	for info := range pupils {
		pupils[info].record()
	}
}

// Create a school register program listing 10 pupils - full name, date of birth and age.
// [Structures][Arrays][Interfaces]
func main() {
	dob := time.Date(1976, 3, 3, 0, 0, 0, 0, time.UTC)
	pupilInfo := schoolRegister{age: calcDob(dob), dob: dob, fullName: "Testing exercise 17_0"}
	addPupilsToRegister(pupilInfo)

	dob1 := time.Date(1971, 3, 8, 0, 0, 0, 0, time.UTC)
	pupilInfo1 := schoolRegister{age: calcDob(dob1), dob: dob1, fullName: "Testing exercise 17_1"}
	addPupilsToRegister(pupilInfo1)

	dob2 := time.Date(1972, 8, 3, 0, 0, 0, 0, time.UTC)
	pupilInfo2 := schoolRegister{age: calcDob(dob2), dob: dob2, fullName: "Testing exercise 17_2"}
	addPupilsToRegister(pupilInfo2)

	dob3 := time.Date(1979, 11, 3, 0, 0, 0, 0, time.UTC)
	pupilInfo3 := schoolRegister{age: calcDob(dob3), dob: dob3, fullName: "Testing exercise 17_3"}
	addPupilsToRegister(pupilInfo3)

	dob4 := time.Date(1970, 3, 13, 0, 0, 0, 0, time.UTC)
	pupilInfo4 := schoolRegister{age: calcDob(dob4), dob: dob4, fullName: "Testing exercise 17_4"}
	addPupilsToRegister(pupilInfo4)

	dob5 := time.Date(1973, 3, 3, 0, 0, 0, 0, time.UTC)
	pupilInfo5 := schoolRegister{age: calcDob(dob5), dob: dob5, fullName: "Testing exercise 17_5"}
	addPupilsToRegister(pupilInfo5)

	dob6 := time.Date(1977, 3, 3, 0, 0, 0, 0, time.UTC)
	pupilInfo6 := schoolRegister{age: calcDob(dob6), dob: dob6, fullName: "Testing exercise 17_6"}
	addPupilsToRegister(pupilInfo6)

	dob7 := time.Date(1978, 3, 3, 0, 0, 0, 0, time.UTC)
	pupilInfo7 := schoolRegister{age: calcDob(dob7), dob: dob7, fullName: "Testing exercise 17_7"}
	addPupilsToRegister(pupilInfo7)

	dob8 := time.Date(1975, 3, 3, 0, 0, 0, 0, time.UTC)
	pupilInfo8 := schoolRegister{age: calcDob(dob8), dob: dob8, fullName: "Testing exercise 17_8"}
	addPupilsToRegister(pupilInfo8)

	dob9 := time.Date(1974, 3, 3, 0, 0, 0, 0, time.UTC)
	pupilInfo9 := schoolRegister{age: calcDob(dob9), dob: dob9, fullName: "Testing exercise 17_9"}
	addPupilsToRegister(pupilInfo9)

	showAllPupils()
}
