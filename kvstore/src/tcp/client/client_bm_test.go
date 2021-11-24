package main

import (
	"bufio"
	"fmt"
	"logging"
	"net"
	"storage"
	"strings"
	"sync"
	"tcp"
	"testing"
)

func BenchmarkTCPClientWithGoRoutines(b *testing.B) {
	var wg sync.WaitGroup
	var mux sync.Mutex

	logging.CreateLogger([]string{})

	domain := "localhost"
	port := "9003"

	l := tcp.Listen(domain, port)
	go tcp.Accept(l, storage.Start())

	// Connect to server
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%s", domain, port))
	if err != nil {
		b.Errorf("TCP Dial failed %+v", err)
	}

	var commands = [...]string{
		"cmd=get | key=1",
		"cmd=create | key=1 | body=testk1",
		"cmd=update | key=1 | body=updatetestk1",
		"cmd=delete | key=1"}

	wg.Add(b.N)

	// act
	for i := 0; i < b.N; i++ {
		go func(i int) {
			defer wg.Done()

			cmdID := i % 5

			mux.Lock()
			fmt.Fprintf(c, commands[cmdID]+"\n")

			message, err := bufio.NewReader(c).ReadString('\n')

			if err != nil {
				fmt.Printf("BenchmarkTCPClientWithGoRoutines::Cmd:: %s Error:: %+v", commands[cmdID], err)
			} else {
				fmt.Printf("BenchmarkTCPClientWithGoRoutines::Cmd:: %s Resp:: %s", commands[cmdID], message)
			}

			mux.Unlock()

			if strings.Contains(commands[cmdID], "cmd=stop") {
				fmt.Println("BenchmarkTCPClientWithGoRoutines::TCP client exiting...")
			}
		}(i)
	}

	wg.Wait()

	l.Close()
	c.Close()
}
