package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("---Backup phase---")

	addr, err := net.ResolveUDPAddr("udp", "10.100.23.23:30030")
	if err != nil {
		fmt.Println("Error resolving UDP address: ", err)
		return
	}
	recvConn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening to UDP address: ", err)
		return
	}

	recvConn.SetReadDeadline(time.Now().Add(3 * time.Second))

	buffer := make([]byte, 16)

	for {
		size, _, err := recvConn.ReadFromUDP(buffer)
		if errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Println("Read Deadline reached: ", err)
			return
		} else if err != nil {
			fmt.Println("failed to read: ", err)
			return
		}
		recvConn.SetReadDeadline(time.Now().Add(3 * time.Second))

		i := (buffer[:size][0])

		fmt.Println(i)
	}
	fmt.Println("timed out")
	recvConn.Close()
}
