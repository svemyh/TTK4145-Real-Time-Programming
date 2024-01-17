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
	raddr, err := net.ResolveUDPAddr("udp", "10.100.23.129:30000") //addressString to actual address(server/)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	//create a connection to raddr through a socket object
	sockConn, err := net.DialUDP("udp", laddr, raddr)

	//Create an empty buffer to to filled
	buffer = make([]byte, 1024)

	sockConn.WriteMsgUDP()

}
