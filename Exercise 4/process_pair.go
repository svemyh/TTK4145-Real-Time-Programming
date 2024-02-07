package main

import (
	//"errors"
	"fmt"
	"net"

	//"os"
	"os/exec"
	"time"
)

func primary(i int) {
	fmt.Println("----Primary----")
	exec.Command("gnome-terminal", "--", "go", "run", "process_pair.go").Run()
	time.Sleep(1 * time.Second)

	writeConn, err := net.Dial("udp", "10.100.23.23:30030")
	if err != nil {
		fmt.Println("Error in creating socket for writing: ", err)
	}
	defer writeConn.Close()

	buffer := make([]byte, 16)

	for {
		i = i + 1
		size := copy(buffer, []byte{byte(i)})

		_, err = writeConn.Write(buffer)
		if err != nil {
			fmt.Println("Error writing to connection: ", err)
		}

		fmt.Println(buffer[:size][0])
		time.Sleep(1 * time.Second)
	}
}

func backup(i int) {
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
	var size int

	for {
		size, _, err = recvConn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("failed to read: ", err)
			break
		}
		val := int(buffer[:size][0])
		if val == i+1 {
			i = val
		}
		recvConn.SetReadDeadline(time.Now().Add(3 * time.Second))
	}
	fmt.Println("timed out")
	recvConn.Close()
	primary(i)

}

func main() {
	i := 0
	backup(i)
}
