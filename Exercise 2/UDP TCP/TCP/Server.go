package main

import (
    "fmt"
    "net"
  	"time"
)

func main() {
    // Listen for incoming connections on port 8080
    ln, err := net.Listen("tcp", ":20000")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Accept incoming connections and handle them
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        // Handle the connection in a new goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    // Close the connection when we're done
    defer conn.Close()

    // Read incoming data
    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the incoming data
    fmt.Printf("Received: %s", buf)
  	time.Sleep(1 * time.Second)
}
