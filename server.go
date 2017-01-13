package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// Establish socket listenr

	netListener, err := net.Listen("tcp", "localhost:1025")
	// Ensure everyhting going well
	CheckError(err)
	// close it later
	defer netListener.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListener.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), "tcp connect success")
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			Log(conn.RemoteAddr().String(), "connection error")
			return
		}

		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
