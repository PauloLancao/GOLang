package main

import (
	"bufio"
	"ch8/portflag"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	var clPortNumber = portflag.PortCommandLine("port", 8000, "TCP port to run clock")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "localhost", *clPortNumber))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
