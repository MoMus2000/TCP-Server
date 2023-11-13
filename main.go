package main

import (
	"fmt"

	"github.com/momus2000/dcas/transport"
)

func main() {
	server := transport.NewTCPFileStreamingServer(":4000")
	fmt.Println("Starting TCP Streaming Server")
	server.StartServer()
}
