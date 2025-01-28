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
		RD: 0,
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

	addr, err := net.ResolveUDPAddr("udp", "198.41.0.4:53")
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

	_, err = decoder.DecodeAnswers(int(h.NSCount))

	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Name Servers: ")
	for i := 0; i< len(decoder.NSRecords); i++ {
		fmt.Println(decoder.NSRecords[i])
	}

	_, err = decoder.DecodeAnswers(int(h.ARCount))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IP Addresses : ")
	for i := 0; i< len(decoder.IPRecords); i++ {
		fmt.Println(decoder.IPRecords[i])
	}

	



}
