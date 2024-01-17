// run the command 'go run TCP' to test both the TCP client and server
package main

import (
"fmt"
"net"
)

func receiver(conn net.Conn){

}

func sender(conn net.Conn){

}

func main() {
conn, err :=
if err != nil {
fmt.PrintIn("Error with TCP dial")
}

go sender(conn)
go receiver(conn)
}


