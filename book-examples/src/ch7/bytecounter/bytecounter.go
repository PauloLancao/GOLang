package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// ByteCounter type
type ByteCounter int

// WordCounter type
type WordCounter int

// LineCounter type
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanWords)
	*w += WordCounter(count)
	return count, nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanLines)
	*l += LineCounter(count)
	return count, nil
}

func retCount(p []byte, fn bufio.SplitFunc) (count int) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(fn)
	count = 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return
}

var counter = struct {
	count uint
}{
	count: 0,
}

// CountingWriter function
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := bufio.NewWriter(w)
	newWLen := (*newWriter).Available()
	counter.count += uint(newWLen)
	newBytesInt64 := int64(newWLen)

	return newWriter, &newBytesInt64
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "17" == "5": len("hello") + "12": len("hello, Dolly")

	var w WordCounter
	w.Write([]byte("Hello This is a line"))
	fmt.Println("Word Counter ", w)

	var l LineCounter
	l.Write([]byte("Hello \nThis \n is \na line\n.\n.\n"))
	fmt.Println("Length ", l)

	var buf bytes.Buffer
	buf.Write([]byte("bla bla bla"))

	CountingWriter(&buf)

	// CountingWriter(&c)
	// CountingWriter(&w)
	// a, b := CountingWriter(&l)
	// fmt.Println(*b)
	// a.Write([]byte(""))
	fmt.Println(counter.count)
}
