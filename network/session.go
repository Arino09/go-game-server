package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Session struct {
	UId            int64
	conn           net.Conn
	IsClose        bool
	packer         IPacker
	WriteCh        chan *SessionPacket
	IsPlayerOnline bool
	MessageHandler func(packet *SessionPacket)
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn: conn,
		packer: &NormalPacker{
			ByteOrder: binary.BigEndian,
		},
		WriteCh: make(chan *SessionPacket, 1),
	}
}

func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

func (s *Session) Read() {
	for {
		err := s.conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		if err != nil {
			fmt.Println(err)
			continue
		}
		message, err := s.packer.Unpack(s.conn)
		if _, ok := err.(net.Error); ok {
			continue
		}
		fmt.Println("Server receive message: ", string(message.Data))
		s.MessageHandler(&SessionPacket{
			Msg:  message,
			Sess: s,
		})
		s.WriteCh <- &SessionPacket{
			Msg: &Message{
				ID:   555,
				Data: []byte("Hello"),
			},
			Sess: s,
		}
	}
}

func (s *Session) Write() {
	for {
		select {
		case resp := <-s.WriteCh:
			s.send(resp.Msg)
		}
	}
}

func (s *Session) send(msg *Message) {
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := s.packer.Pack(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.conn.Write(bytes)
}
