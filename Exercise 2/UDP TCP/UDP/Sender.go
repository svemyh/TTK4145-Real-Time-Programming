package main

import (
	"fmt"
	"net"
)

func main() {
	//def our local address
	laddr, err := net.ResolveUDPAddr("udp", "10.100.23.23:10000") // localIP
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	//def remote addr
	raddr, err := net.ResolveUDPAddr("udp", "10.100.23.129:20013") //addressString to actual address(server/)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	//create a connection to raddr through a socket object
	sockConn, err := net.DialUDP("udp", laddr, raddr)

	defer sockConn.Close()

	//Create an empty buffer to to filled
	buffer := make([]byte, 1024)
	//oob := make([]byte, 1024)

	n := copy(buffer, []byte("Hello World!"))

	fmt.Printf("copied %d bytes to the buffer: %s\n", n, buffer[:n])

	n, err = sockConn.Write(buffer)

	if err != nil {
		fmt.Println("Error resolving WriteMsgUDP", err)
		return
	}

}
