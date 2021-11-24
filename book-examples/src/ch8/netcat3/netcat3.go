package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("Err converting to TCPConn")
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tcpConn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(tcpConn, os.Stdin)
	// conn.Close()
	tcpConn.CloseWrite()

	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Println("NETCAT")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
