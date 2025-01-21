package cmd

import (
	"net"
)

func SendMessage(conn *net.UDPConn, encoded []byte) error {
	addr := net.UDPAddr{
		Port: 53,
		IP:   net.ParseIP("8.8.8.8"),
	}
	conn.WriteToUDP(encoded, &addr)

	return nil
}
