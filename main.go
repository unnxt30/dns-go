package main

import (
	"fmt"
	"log"
	"net"

	"github.com/unnxt30/dns-go/cmd"
	"github.com/unnxt30/dns-go/models"
)

func main() {

	flags := models.HeaderFlags{
		RD: 1,
	}

	header := models.DNSHeader{
		ID:      22,
		Flags:   flags,
		QDCount: 1,
		ANCount: 0,
		NSCount: 0,
		ARCount: 0,
	}
	question := models.DNSQuestion{
		QName:  "dns.google.com",
		QType:  models.A,
		QClass: models.IN,
	}

	message := models.DNSMessage{
		Header:   header,
		Question: question,
	}

	encoded, err := cmd.EncodeMessage(message)

	if err != nil {
		fmt.Println("Error encoding")
		return
	}

	addr, err := net.ResolveUDPAddr("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println("Couldn't create a UDP address")
	}

	conn, err := net.DialUDP("udp", nil, addr)

	_, err = conn.Write(encoded)

	if err != nil {
		fmt.Println("Could not send message to server")
		return
	}

	buf := make([]byte, 1024)
	go func() {
		for {
			n, err := conn.Read(buf)

			if err != nil {
				log.Fatal(err)
				break
			}

			fmt.Println(buf[:n])
		}
	}()

}
