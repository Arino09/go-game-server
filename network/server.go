package network

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
	address  string
	network  string
}

func NewServer(network string, address string) *Server {
	return &Server{
		listener: nil,
		network:  network,
		address:  address,
	}
}

func (s *Server) Run() {
	addr, err := net.ResolveTCPAddr(s.network, s.address)
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpListener, err := net.ListenTCP(s.network, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.listener = tcpListener
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			go func() {
				session := NewSession(conn)
				session.Run()
			}()
		}
	}()
}
