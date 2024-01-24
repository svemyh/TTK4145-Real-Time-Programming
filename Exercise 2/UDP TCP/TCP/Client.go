package main

import (
	"fmt"
	"net"
)

// Socket -> Connect -> Send <-> Recv (from server) -> Send -> Close

func main() {
	conn, err := net.Dial("tcp", "10.24.38.142:10000")
	if err != nil {
		fmt.Println("The connection failed. Error: ", err)
		return
	} else {
		fmt.Printf("The connection was established to: %s \n", conn.RemoteAddr()) //wtf
	}

	_, err = conn.Write([]byte("Gruppe 13!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
}
