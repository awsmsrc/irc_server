package main

import (
	"fmt"
	"net"

	"github.com/awsmsrc/irc_server/server"
)

func main() {
	fmt.Println("Starting IRC server on port: 6667")
	s := server.NewServer()
	listener, _ := net.Listen("tcp4", ":6667")
	fmt.Println("Server bound and awaiting connections")
	for {
		connection, _ := listener.Accept()
		s.AddConnection(connection)
	}
}
