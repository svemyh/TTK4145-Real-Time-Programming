package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
	"errors"
)

func main() {
	i := 0
	exec.Command("gnome-terminal", "--", "go", "run", "A.go").Run()
	time.Sleep(1 *time.Second)

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
