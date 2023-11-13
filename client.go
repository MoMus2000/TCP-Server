package main

import (
	"fmt"
	"time"

	"github.com/momus2000/dcas/transport"
)

func main() {
	time.Sleep(time.Second * 5)
	fmt.Println("Starting Music Client")
	client := transport.NewMusicClient("localhost:4000")
	fmt.Println("Listening over localhost port 4000")
	client.Listen()
}
