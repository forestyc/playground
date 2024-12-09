package udp

import (
	"fmt"
	"net"
)

type Client struct {
	readBuf  chan []byte
	writeBuf chan []byte
	stop     chan int
	conn     *net.UDPConn
}

func NewClient(addr string) (*Client, error) {
	var client Client
	client.readBuf = make(chan []byte)
	client.writeBuf = make(chan []byte)
	client.stop = make(chan int)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	client.conn, err = net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (c *Client) Run() {
	go c.read()
	go c.write()
	<-c.stop
	err := c.conn.Close()
	if err != nil {
		return
	}
}

func (c *Client) Recv() []byte {
	return <-c.readBuf
}

func (c *Client) Send(buf []byte) {
	c.writeBuf <- buf
}

func (c *Client) read() []byte {
	buffer := make([]byte, 1024)
	for {
		n, err := c.conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response:", err)
			continue
		}
		frame := make([]byte, n)
		copy(frame, buffer[:n])
		c.readBuf <- frame
	}
}

func (c *Client) write() {
	for {
		data := <-c.writeBuf
		_, err := c.conn.Write(data)
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}
	}
}
