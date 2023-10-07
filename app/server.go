package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer l.Close()
	log.Printf("server listening at localhost %s", l.Addr())

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("Text received from client: %s", text)

		switch strings.ToUpper(text) {
		case "PING":
			fmt.Fprintf(conn, "%s\n", "+PONG\r\n")
		default:
			fmt.Fprintf(conn, "%s\n", "cmd not handled")
		}
	}
}
