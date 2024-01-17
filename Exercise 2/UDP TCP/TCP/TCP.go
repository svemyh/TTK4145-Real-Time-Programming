package main

import (
	"fmt"
	"net"
)

//```C
//addr = new InternetAddress(serverIP, serverPort)
//sock = new Socket(tcp) // TCP, aka SOCK_STREAM
//|sock.connect(addr)
// use sock.recv() and sock.send(), just like with UDP
//```

func client(conn net.Conn) {
	// Connect to the server

	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
		fmt.Println(err)
	}
}

//```C
// Send a message to the server:  "Connect to: " <your IP> ":" <your port> "\0"

// do not need IP, because we will set it to listening state
//addr = new InternetAddress(localPort)
//acceptSock = new Socket(tcp)

// You may not be able to use the same port twice when you restart the program, unless you set this option
//acceptSock.setOption(REUSEADDR, true)
//acceptSock.bind(addr)

//loop {
// backlog = Max number of pending connections waiting to connect()
//    newSock = acceptSock.listen(backlog)

// Spawn new thread to handle recv()/send() on newSock
//}
//```

func server(conn net.Conn) {
	ln, err := net.Listen("tcp", ":20000")
	if err != nil {
		fmt.Println("Error while listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server listening on ", ln)

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
	conn, err := net.Dial("tcp", "10.24.37.52:20000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Start the client and server
	go server(conn)
	go client(conn)
}
