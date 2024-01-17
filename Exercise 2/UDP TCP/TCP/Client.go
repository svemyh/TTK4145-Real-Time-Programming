package main

import (
    "fmt"
    "net"
)

func main() {
    // Connect to the server
    conn, err := net.Dial("tcp", "10.24.37.52:20000")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Send some data to the server
    _, err = conn.Write([]byte("Hello, server!"))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Close the connection
    conn.Close()
}
