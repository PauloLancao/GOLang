package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout

	w.Write([]byte("hello1")) // "hello1"
	fmt.Printf("%T\n", w)

	w = new(bytes.Buffer)

	w.Write([]byte("hello2")) // "hello2"
	fmt.Printf("%T\n", w)

	w = nil

	var x interface{} = time.Now()
	fmt.Println(x)

	var x1 interface{} = []int{1, 2, 3}
	fmt.Println(x1)
}
