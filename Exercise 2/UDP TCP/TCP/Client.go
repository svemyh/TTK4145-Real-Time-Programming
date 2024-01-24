package main

import (
	"fmt"
	"net"
)

// Socket -> Connect -> Send <-> Recv (from server) -> Send -> Close

func main() {
	port := 10000
	ip := "10.100.23.23"

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("The connection failed. Error: ", err)
		return
	} else {
		fmt.Printf("The connection was established to: %s \n", conn.RemoteAddr()) //wtf
	}

	// _, err =
	conn.Write([]byte("Gruppe 13!"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// msg := fmt.Sprintf("Connected to: %s in client\n", conn.LocalAddr())
	// conn.Write([]byte(msg))

	// ls, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// if err != nil {
	// 	fmt.Println("The connection failed. Error:", err)
	// 	return
	// }
	// defer ls.Close()

	buffer := make([]byte, 1024)
	bytes, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	fmt.Println(string(buffer[:bytes]))

	defer conn.Close()
}
