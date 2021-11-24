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

func TestTCPClient(t *testing.T) {
	// arrange
	logging.CreateLogger([]string{})
	defer logging.Close()

	domain := "localhost"
	port := "60402"

	l := tcp.Listen(domain, port)
	go tcp.Accept(l, storage.Start())

	// Connect to server
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%s", domain, port))
	if err != nil {
		t.Errorf("TestTCPClient::TCP Dial failed %+v", err)
	}

	var commands = [...]string{
		"cmd=get | key=1",
		"cmd=create | key=1 | body=testk1",
		"cmd=update | key=1 | body=updatetestk1",
		"cmd=delete | key=1"}

	iterations := 1000

	for i := 0; i < iterations; i++ {
		cmdID := i % len(commands)

		fmt.Fprintf(c, commands[cmdID]+"\n")

		message, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Printf("TestTCPClient::Cmd:: %s Error:: %+v", commands[cmdID], err)
		} else {
			fmt.Printf("TestTCPClient::Cmd:: %s Resp:: %s", commands[cmdID], message)
		}

		if strings.Contains(commands[cmdID], "cmd=stop") {
			fmt.Println("TestTCPClient::TCP client exiting...")
		}
	}

	l.Close()
	c.Close()
}

func TestTCPClientWithGoRoutines(t *testing.T) {
	var wg sync.WaitGroup
	var mux sync.Mutex

	logging.CreateLogger([]string{})
	defer logging.Close()

	domain := "localhost"
	port := "60401"

	l := tcp.Listen(domain, port)
	go tcp.Accept(l, storage.Start())

	// Connect to server
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%s", domain, port))
	if err != nil {
		t.Errorf("TestTCPClientWithGoRoutines::TCP Dial failed %+v", err)
	}

	var commands = [...]string{
		"cmd=get | key=1",
		"cmd=create | key=1 | body=testk1",
		"cmd=update | key=1 | body=updatetestk1",
		"cmd=delete | key=1"}

	iterations := 1000

	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func(i int) {
			defer wg.Done()

			cmdID := i % len(commands)

			mux.Lock()
			fmt.Fprintf(c, commands[cmdID]+"\n")

			message, err := bufio.NewReader(c).ReadString('\n')

			if err != nil {
				fmt.Printf("TestTCPClientWithGoRoutines::Cmd:: %s Error:: %+v", commands[cmdID], err)
			} else {
				fmt.Printf("TestTCPClientWithGoRoutines::Cmd:: %s Resp:: %s", commands[cmdID], message)
			}

			mux.Unlock()

			if strings.Contains(commands[cmdID], "cmd=stop") {
				fmt.Println("TestTCPClientWithGoRoutines::TCP client exiting...")
			}
		}(i)
	}

	wg.Wait()

	l.Close()
	c.Close()
}
