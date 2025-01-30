package cmd

import (
	"log"
	"net"
)

func ExchangeMessage(conn *net.UDPConn, encoded []byte) ([]byte, error) {

	_, err := conn.Write(encoded)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return buf[:n], nil
}
