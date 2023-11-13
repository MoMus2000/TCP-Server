package transport

import (
	"fmt"
	"net"
	"os"
)

type TCPFileServer struct {
	Listener net.Listener
	Address  string
	quitch   chan struct{}
}

func NewTCPFileServer(Address string) *TCPFileServer {
	return &TCPFileServer{
		Address: Address,
	}
}

func (t *TCPFileServer) StartServer() error {
	listener, err := net.Listen("tcp", t.Address)
	if err != nil {
		return err
	}
	t.Listener = listener

	go t.acceptAndServe()

	<-t.quitch

	return nil
}

func (t *TCPFileServer) acceptAndServe() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			conn.Close()
			fmt.Println("Closed Connection : ", conn.RemoteAddr().String())
		}
		go t.serveFile(conn)
	}
}

func (t *TCPFileServer) serveFile(conn net.Conn) {
	file, err := os.ReadFile("./dummyFiles/testFile.png")
	if err != nil {
		fmt.Println("Error importing the file from folder: ", err)
		return
	}
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}
	fmt.Println("Sending File as bytes")
	conn.Write(file)
	conn.Close()
}
