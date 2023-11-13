package transport

import (
	"fmt"
	"net"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type MusicClient struct {
	Address  string
	Listener net.Listener
}

func NewMusicClient(Addr string) *MusicClient {
	return &MusicClient{
		Address: Addr,
	}
}

func (client *MusicClient) Listen() {
	conn, err := net.Dial("tcp", client.Address)
	if err != nil {
		fmt.Println("Unable to establish connection with server:", err)
		return
	}
	defer conn.Close()

	streamer, format, err := mp3.Decode(conn)
	if err != nil {
		fmt.Println("Error decoding mp3:", err)
		return
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		fmt.Println("Error initializing speaker:", err)
		return
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Block until done
	<-done
}
