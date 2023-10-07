package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
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

	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("error reading from connection: %v", err.Error())
			return
		}

		fmt.Printf("Received: %s", string(buff[:n]))
		conn.Write([]byte("+PONG\r\n"))
	}
}
