package server

import (
	"fmt"
	"io"
	"net"

	"github.com/sorcix/irc"
)

type Server struct {
	prefix   *irc.Prefix
	Channels map[string]*Channel
	Users    map[string]*User
}

func NewServer() *Server {
	s := new(Server)
	s.Channels = make(map[string]*Channel)
	s.Users = make(map[string]*User)
	s.prefix = &irc.Prefix{
		Name: "go-irc",
	}
	return s
}

func (s *Server) AddConnection(connection net.Conn) {
	fmt.Println("Connection received")
	go func(connection net.Conn) {
		// Create a decoder that reads from given io.Reader
		dec := irc.NewDecoder(connection)
		for {
			// Decode the next IRC message
			message, err := dec.Decode()
			if err != nil {
				if err == io.EOF {
					return
				} else {
					panic(err)
				}
			}
			s.handleMessage(message, connection)
		}
	}(connection)
}

func (s *Server) handleMessage(msg *irc.Message, connection net.Conn) {
	fmt.Printf("raw message: %s\n", msg)
	switch msg.Command {
	case "CAP":
		s.handleCap(msg, connection)
	case "NICK":
		s.handleNick(msg, connection)
	case "JOIN":
		s.handleJoin(msg, connection)
	case "USER":
		s.handleUser(msg, connection)
	default:
		s.handleUnknown(msg)
	}
}

func (s *Server) handleCap(msg *irc.Message, conneciton net.Conn) {
	fmt.Printf("CAP message: %s\n", msg)
}

func (s *Server) handleNick(msg *irc.Message, connection net.Conn) {
	fmt.Printf("NICK message: %s\n", msg)
}

func (s *Server) handleJoin(msg *irc.Message, connection net.Conn) {
	fmt.Printf("JOIN message: %s\n", msg)
}

func (s *Server) handleUser(msg *irc.Message, connection net.Conn) {
	resp_1 := irc.Message{
		Prefix:  s.prefix,
		Command: irc.RPL_WELCOME,
	}
	resp_2 := irc.Message{
		Prefix:   s.prefix,
		Command:  irc.RPL_YOURHOST,
		Trailing: "your host is localhost",
	}
	resp_3 := irc.Message{
		Prefix:   s.prefix,
		Command:  irc.RPL_CREATED,
		Trailing: "This server was created at the beginning of time",
	}
	resp_4 := irc.Message{
		Prefix:   s.prefix,
		Command:  irc.RPL_CREATED,
		Trailing: "go-irc v0.0.1 dioswkg biklmnopstv",
	}

	encoder := irc.NewEncoder(connection)

	encoder.Write(resp_1.Bytes())
	encoder.Write(resp_2.Bytes())
	encoder.Write(resp_3.Bytes())
	encoder.Write(resp_4.Bytes())
}

func (s *Server) handleUnknown(msg *irc.Message) {
	fmt.Printf("unknown message: %s\n", msg)
}
