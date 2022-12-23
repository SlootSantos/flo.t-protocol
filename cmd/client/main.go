package main

import (
	"fmt"
	"net"

	protocol "github.com/SlootSantos/flo.t-server/pkg/protocol"
)

func main() {
	CONNECT := "localhost:3333"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := protocol.Message{
		Version:  "flo.t/1.0",
		Resource: "/test",
		Headers:  protocol.Headers{"Test-Heder": "test"},
	}

	fmt.Println(message.String())

	fmt.Fprintf(c, message.String())
}
