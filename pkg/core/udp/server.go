package udp

import (
	"fmt"
	"log"
	"net"
)

type Data struct {
	Addr *net.UDPAddr
	Buf  []byte
}

type Server struct {
	readBuf  chan Data
	writeBuf chan Data
	stop     chan int
	conn     *net.UDPConn
}

func NewServer(address string) (*Server, error) {
	var server Server
	server.readBuf = make(chan Data, 1024)
	server.writeBuf = make(chan Data, 1024)
	server.stop = make(chan int)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	server.conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}
	return &server, nil
}

func (s *Server) Run() {
	go s.read()
	go s.write()
	<-s.stop
	if err := s.conn.Close(); err != nil {
		return
	}
}

func (s *Server) Recv() Data {
	return <-s.readBuf
}

func (s *Server) Send(data Data) {
	s.writeBuf <- data
}

func (s *Server) read() {
	buffer := make([]byte, 1024)
	for {
		n, addr, err := s.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading from UDP:", err)
			continue
		}
		frame := Data{
			Addr: addr,
			Buf:  make([]byte, n),
		}
		copy(frame.Buf, buffer[:n])
		s.readBuf <- frame
	}
}

func (s *Server) write() {
	for {
		data := <-s.writeBuf
		_, err := s.conn.WriteToUDP(data.Buf, data.Addr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}
