package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// This is an echo-server skeleton. It listens on port 9000 and accepts
// one connection. It reads line from the client and sends back the same
// line. When this works, it means we can open the port, accept a TCP
// connection and move bytes in both directions.

func main() {
	ln, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 9000")

	// Runs when the main exits
	defer ln.Close()

	numConnections := 0
	for {
		conn, _ := ln.Accept()
		numConnections++
		go handleConn(conn, numConnections)
	}	
}

func handleConn(conn net.Conn, clientNumber int) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	fmt.Println("Client", clientNumber, "connected")

	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
			break
		}
		
		fmt.Print("Client ", clientNumber, ": ", msg)
	}
}