package main

import (
	"fmt"
	"github.com/forestyc/playground/pkg/core/udp"
	"log"
)

func main() {
	go client()
	server()
}

func server() {
	udpServer, err := udp.NewServer("127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			data := udpServer.Recv()
			fmt.Println(data.Addr, string(data.Buf))
		}
	}()
	udpServer.Run()
}

func client() {
	udpClient, err := udp.NewClient("127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			udpClient.Send([]byte("hello"))
		}
	}()

	udpClient.Run()
}
