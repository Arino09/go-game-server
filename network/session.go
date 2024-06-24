package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Session struct {
	conn    net.Conn
	packer  *NormalPacker
	chWrite chan *Message
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPacker(binary.BigEndian),
		chWrite: make(chan *Message, 1),
	}
}

func (s *Session) Run() {
	go s.Read()
	go s.Write()

}

func (s *Session) Read() {
	err := s.conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		fmt.Println(err)
	}
	for {
		message, err := s.packer.Unpack(s.conn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Server receive message: ", string(message.Data))
		s.chWrite <- &Message{
			Id:   999,
			Data: []byte("Hello"),
		}
	}
}

func (s *Session) Write() {
	for {
		select {
		case msg := <-s.chWrite:
			s.send(msg)
		}
	}
}

func (s *Session) send(msg *Message) {
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := s.packer.Pack(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = s.conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
