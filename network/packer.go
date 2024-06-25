package network

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

type NormalPacker struct {
	ByteOrder binary.ByteOrder
}

// Pack	|data 长度|id|data|
func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	buffer := make([]byte, 8+8+len(message.Data))
	p.ByteOrder.PutUint64(buffer[:8], uint64(len(buffer)))
	p.ByteOrder.PutUint64(buffer[8:16], message.ID)
	copy(buffer[16:], message.Data)
	return buffer, nil
}

func (p *NormalPacker) Unpack(reader io.Reader) (*Message, error) {
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second * 2))
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 8+8)
	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}
	totalLen := p.ByteOrder.Uint64(buffer[:8])
	id := p.ByteOrder.Uint64(buffer[8:])
	dataLen := totalLen - 16
	dataBuffer := make([]byte, dataLen)
	_, err = io.ReadFull(reader, dataBuffer)
	if err != nil {
		return nil, err
	}
	msg := &Message{
		ID:   id,
		Data: dataBuffer,
	}
	return msg, nil
}
