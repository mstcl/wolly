package packet

import (
	"fmt"
	"net"
)

const (
	broadcastIP = "255.255.255.255"
	port        = "9"
)

func Send(stream []byte) error {
	broadcast := fmt.Sprintf("%s:%s", broadcastIP, port)
	local := &net.UDPAddr{}
	remote, err := net.ResolveUDPAddr("udp", broadcast)
	if err != nil {
		return fmt.Errorf("resolve udp: %v", err)
	}

	conn, err := net.DialUDP("udp", local, remote)
	if err != nil {
		return fmt.Errorf("dial udp: %v", err)
	}

	defer conn.Close()

	if _, err = conn.Write(stream); err != nil {
		return fmt.Errorf("write to udp connection: %v", err)
	}

	return nil
}
