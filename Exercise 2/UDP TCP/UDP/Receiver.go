// Run the command 'go run Receiver.go' to test the UDP receiver
// NOT TESTED SINCE RDMD WONT WORK
package main

import (
	"fmt"
	"net"
)

func main() {

<<<<<<< HEAD
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:30000") //addressString to actual address(server/)
=======
	addr, err := net.ResolveUDPAddr("udp", ":20013") //addressString to actual address(server/)
>>>>>>> 19f6597c142f3d54a27a3ff5d2bc34a6a34112e1
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// TODO: recvSock = new Socket(udp). Bind address we want to use to the socket
	recvSock, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer recvSock.Close() // Close recvSock AFTER surrounding main function completes

	buffer := make([]byte, 1024) // a buffer where the received network data is stored byte[1024] buffer

	// var fromWho *net.UDPAddr // Variable empty address about who sent the data

	for {
		buffer = make([]byte, 1024)

		// TODO: Return number of bytes received
		numBytesReceived, fromWho, err := recvSock.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error readFromUDP:", err)
			return
		}

		message := string(buffer[:numBytesReceived])

		localIP, err := net.ResolveUDPAddr("udp", "10.100.23.23:10000") // localIP
		if err != nil {
			fmt.Println("Error resolving UDP address:", err)
			return
		}

		if string(fromWho.IP) != string(localIP.IP) {
			fmt.Println(message)
			//fmt.PrintIn("Filtered out: ", string(buffer[0:numBytesReceived]))
		} else {
<<<<<<< HEAD
			fmt.Println("rand message is: ", message)
=======
			fmt.Println("random message is: ", message)
>>>>>>> 19f6597c142f3d54a27a3ff5d2bc34a6a34112e1
		}
	}
}
