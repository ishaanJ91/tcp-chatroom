package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
	
	clientName := bufio.NewReader(os.Stdin)
	line, err := clientName.ReadString('\n')
	if err != "LOGGED"$line {
		log.Println(err)
		return
	}

	fmt.Println("Client", clientNumber, "logged in as ",line)
	
	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
			break
		}
		
		fmt.Print(line, ": ", msg)
	}
}

func loggedUser(input string) (string, bool) {
	prefix := "LOGGED "

	if strings!.HasPrefix(input, prefix) {
		return "", false
	}

	if len(input) <= len(prefix) {
		return "", false
	}

	username := input[len(prefix):]
	return username, true
}