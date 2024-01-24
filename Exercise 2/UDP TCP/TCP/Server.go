package main

import (
	"fmt"
	"net"
	"time"
)

// Socket -> Bind -> Listen -> Accept -> Recv (from client) <-> Send -> Recv -> Close

func main() {

	port := 10000
	//Listen
	//conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("The connection failed. Error:", err)
		return
	}
	defer ls.Close()

	fmt.Println("Connected to port:", port)
	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		go handleConnection(conn)
		time.Sleep(1 * time.Second)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	//Accept
	fmt.Printf("Accepted connection from %s\n", conn.LocalAddr())

	//Send
	msg := fmt.Sprintf("Connect to: %s\n", conn.LocalAddr())
	conn.Write([]byte(msg))
}

// func main() {
// 	// Listen for incoming connections on port 8080
// 	ln, err := net.Listen("tcp", ":2001	// ls, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
// if err != nil {
// 	fmt.Println("The connection failed. Error:", err)
// 	return
// }
// defer ls.Close()
//3")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Accept incoming connections and handle them
// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 		// Handle the connection in a new goroutine
// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	// Close the connection when we're done
// 	defer conn.Close()

// 	// Read incoming data
// 	buf := make([]byte, 1024)
// 	_, err := conn.Read(buf)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// Print the incoming data
// 	fmt.Printf("Received: %s", buf)
// 	// To not get spammed
// 	time.Sleep(1 * time.Second)
// }
