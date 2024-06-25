package network

import (
	"fmt"
	"net"
)

type Server struct {
	tcpListener     net.Listener
	OnSessionPacket func(packet *SessionPacket)
}

func NewServer(address string) *Server {
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	tcpListener, err := net.ListenTCP("tcp", resolveTCPAddr)
	if err != nil {
		panic(err)
	}
	s := &Server{
		tcpListener: tcpListener,
	}
	return s
}

func (s *Server) Run() {
	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				continue
			}
		}
		go func() {
			fmt.Println("New connection from " + conn.RemoteAddr().String())
			newSession := NewSession(conn)
			SessionMgrInstance.AddSession(newSession)
			newSession.Run()
			SessionMgrInstance.DelSession(newSession)
		}()
	}
}
