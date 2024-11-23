package packet

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Address [6]byte

type Packet struct {
	Header  [6]byte
	Payload [16]Address
}

const header = 0xFF

func New(m string) (*Packet, error) {
	hw, err := net.ParseMAC(m)
	if err != nil {
		return nil, err
	}

	packet := Packet{}
	for idx := range packet.Header {
		packet.Header[idx] = header
	}

	addr := Address{}
	copy(addr[:], hw)

	for idx := range packet.Payload {
		packet.Payload[idx] = addr
	}

	return &packet, nil
}

func (p *Packet) Marshal() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.BigEndian, p); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
