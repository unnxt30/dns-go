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
		QType: 1,
		QClass: 1,
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

	if err != nil {
		log.Fatal(err)
		return 
	}

	defer conn.Close()

	resp, err := cmd.ExchangeMessage(conn, encoded)

	if err != nil {
		fmt.Println("Error exchanging message")
		return
	}

	// _, err = cmd.DecodeHeader(resp)

	// if err != nil {
	// 	fmt.Println("Error decoding")
	// 	return
	// }


	decoder := cmd.DNSDecoder{
		Encoded: resp,
		Offset: 0,
	}	

	h, err := decoder.DecodeHeader()
	if err != nil {
		fmt.Println(err)
		return 
	}

	_, err = decoder.DecodeQuestion()

	if err != nil {
		fmt.Println("Error decoding")
		return
	}

	_, err = decoder.DecodeAnswers(int(h.ANCount))

	if err != nil {
		fmt.Println(err)
		return 
	}


}
