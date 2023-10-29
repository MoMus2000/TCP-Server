package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Port     string
	Listener net.Listener
	quitch   chan struct{}
}

func NewServer(Port string) *Server {
	return &Server{
		Port: Port,
	}
}

func (s *Server) StartServer() error {
	listener, err := net.Listen("tcp", s.Port)
	if err != nil {
		log.Fatal(err)
	}
	s.Listener = listener

	// Accept any incoming connections

	go s.AcceptConnection()

	<-s.quitch
	return nil
}

func (s *Server) AcceptConnection() error {
	for {
		con, err := s.Listener.Accept()
		if err != nil {
			return err
		}
		go s.ReadFromConnection(con)
	}
}

func (s *Server) ReadFromConnection(Con net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := Con.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Message From the server, ", string(buffer[:n]))
	}
}

func main() {
	fmt.Println("Starting the TCP Server")
	Server := NewServer(":8080")
	Server.StartServer()
}
