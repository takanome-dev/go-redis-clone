package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close()
	// scanner := bufio.NewScanner(conn)
	// for scanner.Scan() {
	// 	line := strings.ToLower(scanner.Text())
	// 	log.Printf("Received: %s\n", line)

	// 	switch line {
	// 	case "ping":
	// 		conn.Write([]byte("+PONG\r\n"))
	// 	case "echo":
	// 		conn.Write([]byte("+Client sent an echo cmd\r\n"))
	// 	default:
	// 		conn.Write([]byte("+Default case hit\r\n"))
	// 	}
	// }
	for {
		buff := make([]byte, 512)
		n, err := conn.Read(buff)
	
		if errors.Is(err, io.EOF) {
			log.Printf("client %v closed a connection\n", conn.RemoteAddr())
			return
		}
	
		if err != nil {
			log.Printf("error reading from connection: %v\n", err.Error())
			return
		}

		log.Printf("capacity of the buffer is: %d", cap(buff))
		// text := strings.ToUpper(string(buff[:n]))
		// var response string
		log.Printf("Received: %s\n", buff[:n])

		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			log.Fatalf("error writing to connection: %v", err.Error())
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Fatalf("Failed to bind to port 6379: %v", err.Error())
	}

	defer l.Close()
	log.Printf("server listening at localhost %s\n", l.Addr())
	
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection: %v", err.Error())
		}
		
		go handler(conn)
	}
}
