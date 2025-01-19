package main

import (
	"fmt"

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
		QName:  "www.google.com",
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

	fmt.Println(encoded)

}
