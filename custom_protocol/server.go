package main

import (
	"fmt"
	"net"
	"os"
	"protocol"
)

func main() {
	netListener, err := net.Listen("tcp", "localhost:6060")
	CheckError(err)

	defer netListener.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListener.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	tmpBuffer := make([]byte, 0)

	go reader(Channel)
}
