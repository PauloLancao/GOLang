// Netcat1 is a read-only TCP client.
package main

import (
	"ch8/portflag"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	var clPortNumber = portflag.PortCommandLine("port", 8000, "TCP port to run clock")
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", "localhost", *clPortNumber))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Println("NETCAT")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
