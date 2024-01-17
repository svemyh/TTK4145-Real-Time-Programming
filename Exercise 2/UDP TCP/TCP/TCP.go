package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func client(conn net.Conn) {
	for {
		// read input from stdin
		fmt.Print("Text to send: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		// send to socket
		fmt.Fprintf(conn, text+"\n")

		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

func handleConnection(newSock net.Conn) {
	defer newSock.Close()

	for {
		// Read the incoming message
		message, err := bufio.NewReader(newSock).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		fmt.Print("Message Received:", string(message))

		// Process the message (e.g., convert to uppercase)
		newmessage := strings.ToUpper(message)

		// Send the processed message back to the client
		newSock.Write([]byte(newmessage + "\n"))

		time.Sleep(1 * time.Second)
	}
}

func server(conn net.Conn) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error while listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server listening on :8080")

	for {
		// Accept a new connection
		newSock, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(newSock)
	}
}

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "10.100.23.23:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Start the client and server
	go client(conn)
	go server(conn)
}
