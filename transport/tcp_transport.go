package transport

import (
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
	Listener net.Listener
	Address  string
}

func NewTCPServer(Address string) *TCPServer {
	return &TCPServer{
		Address: Address,
	}
}

func (t *TCPServer) StartServer() {
	Listener, err := net.Listen("tcp", t.Address)
	if err != nil {
		log.Fatal(err)
	}
	t.Listener = Listener
	go t.ListenAndAccept()
}

func (t *TCPServer) ListenAndAccept() error {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			return err
		}
		fmt.Println("Connected to ", conn.RemoteAddr().String())
		go t.acceptConnection(conn)
	}
}

func (t *TCPServer) acceptConnection(conn net.Conn) {
	buffer := make([]byte, 1028)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		fmt.Println(string(buffer[:n]))
	}
}
