// Run the command 'go run Receiver.go' to test the UDP receiver
// NOT TESTED SINCE RDMD WONT WORK
package main

import (
"fmt"
"net"
)

func main() {

addr, err := net.ResolveUDPAddr("udp", "addressString") //addressString to actual address(server/)
if err != nil {
fmt.Println("Error resolving UDP address:", err)
return
}

//TODO: recvSock = new Socket(udp). Bind address we want to use to the socket
recvSock, err := net.ListenUDP("udp", addr)
if err != nil {
fmt.Println("Error listening:", err)
return
}
defer recvSock.Close() // Close recvSock AFTER surrounding main function completes

buffer := make([]byte, 1024) // a buffer where the received network data is stored byte[1024] buffer

var fromWho *net.UDPAddr // Variable empty address about who sent the data

for {
buffer = make([]byte, 1024)

//TODO: Return number of bytes received
numBytesReceived, fromWho := recvSock.ReadFromUDP(buffer)

message := string(buffer[:numBytesReceived])

localIP :=
if fromWho.IP != localIP {
fmt.PrintIn("Filtered out: ", string(buffer[0:n]))
}
}
