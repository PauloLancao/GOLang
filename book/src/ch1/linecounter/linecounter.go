package main

import "fmt"

type linecounter int

func (lc *linecounter) Write(p []byte) (int, error) {
	counter := 1
	for _, v := range p {
		if v == '\n' {
			counter++
		}
	}
	*lc = linecounter(counter)
	return counter, nil
}

func main() {
	lc := new(linecounter)

	c, err := lc.Write([]byte("Here is a string....\n testing more breaklines\n testing more"))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(c)
	fmt.Println(*lc)

	lc = new(linecounter)

	c, err = lc.Write([]byte("Here is a string....\n testing more breaklines\n testing more\n"))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(c)
	fmt.Println(*lc)
}
