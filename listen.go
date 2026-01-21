package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// This is an echo-server skeleton. It listens on port 9000 and accepts
// one connection. It reads line from the client and sends back the same
// line. When this works, it means we can open the port, accept a TCP
// connection and move bytes in both directions.

func main() {
	ln, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Println("Error listening to our port 9000", err.Error())
	}
	fmt.Println("Listening on port 9000")

	// Runs when the main exits
	defer ln.Close()

	numConnections := 0
	m := make(map[string]int)
	for {
		conn, err := ln.Accept()
		
		if err != nil {
			log.Printf("Error while connecting to the addr: %s\n", err.Error(), conn.RemoteAddr())
		} else {
			fmt.Printf("Connection established to remote addr: %s\n", conn.RemoteAddr())
		}

		numConnections++
		go handleConnServer(conn, numConnections, m) 
	}	
}

func handleConnServer(conn net.Conn, clientNumber int, m map[string]int) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	
	if err != nil {
		log.Println(err)
		return
	}

	loggedUser, ok := loggedUser(line, m)
	if !ok {
		fmt.Fprintln(conn, "Invalid login message")
		return
	}

	messageServerTerm(clientNumber, loggedUser)
	messageClientTerm(loggedUser, conn)
	
	for {
		fmt.Fprintf(conn, "You: ")
		msg, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
			break
		}
		
		fmt.Print(loggedUser, ": ", strings.TrimSpace(msg), "\n")
	}
}

func loggedUser(input string, m map[string]int) (string, bool) {
	prefix := "LOGGED "

	if !strings.HasPrefix(input, prefix) {
		return "", false
	}

	if len(input) <= len(prefix) {
		return "", false
	}

	username := strings.TrimSpace(input[len(prefix):])

	if _, exists := m[username]; exists {
		return "", false
	}

	m[username]++
	return username, true
}

func messageClientTerm(username string, conn net.Conn) {
	fmt.Fprintf(conn, "-----------------------------\n")
	fmt.Fprintf(conn, "WELCOME %s\n", username)
	fmt.Fprintf(conn, "-----------------------------\n")
}

func messageServerTerm(clientNumber int, username string) {
	fmt.Println("-----------------------------")
	fmt.Println("Client", clientNumber, "logged in as", username)
	fmt.Println("-----------------------------")
}