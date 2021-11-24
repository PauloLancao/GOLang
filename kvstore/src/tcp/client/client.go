package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	connect := arguments[1]
	c, err := net.Dial("tcp", connect)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Available commands: cmd=getall|get|create|update|delete|stop")
	fmt.Println("Passing parameters: key=<?>|body=<?>")
	fmt.Println("E.g. cmd=get | key=1 '->'")
	fmt.Println("E.g. cmd=create | key=1 | body=<?> '->'")
	fmt.Println("Terminate character is '->'")
	fmt.Println("Commandline starting...")

	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Req::>> ")

		var lines string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasSuffix(line, "->") {
				lines += strings.Replace(line, "->", "", -1)
				break
			}
			lines += line
		}

		fmt.Fprintf(c, lines+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("Resp::--> " + string(message))

		if strings.Contains(lines, "cmd=stop") {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
