package transport

import (
	"fmt"
	"net"
	"sync"
)

// Manage concurrency by using locks properly
// Implement a graceful shutdown
// Watch Tsoding's video to see some other stuff you can do here

type TCPBroadCastServer struct {
	Listener    net.Listener
	Address     string
	Connections []net.Conn
	Lock        sync.Mutex
	quicth      chan struct{}
}

func InitializeTCPBroadCastServer(Address string) *TCPBroadCastServer {
	return &TCPBroadCastServer{
		Address:     Address,
		Connections: []net.Conn{},
		Listener:    nil,
	}
}

func (t *TCPBroadCastServer) StartServer() {
	Listener, err := net.Listen("tcp", t.Address)
	if err != nil {
		fmt.Println(err)
	}
	t.Listener = Listener
	go t.listenForConnections()
	<-t.quicth
}

func (t *TCPBroadCastServer) listenForConnections() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		t.Lock.Lock()
		t.Connections = append(t.Connections, conn)
		t.Lock.Unlock()
		go t.broadcast(conn)
	}
}

func (t *TCPBroadCastServer) broadcast(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
		}
		// Broadcast
		for _, connections := range t.Connections {
			if connections == conn {
				continue
			}
			connections.Write(buffer[:n])
		}
	}
}
