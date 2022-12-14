package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	procotol "github.com/SlootSantos/flo.t-server/pkg/protocol"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	// Make a buffer to hold incoming data.

	message, _ := procotol.ParseMessage(reader)
	fmt.Printf("%+v", message)

	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
