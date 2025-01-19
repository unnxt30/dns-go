package cmd

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"unsafe"

	"github.com/unnxt30/dns-go/models"
)

func encodeMessage(Message models.DNSMessage) (*bytes.Buffer, error) {
	msgSize := int(unsafe.Sizeof(Message))
	buffer := bytes.NewBuffer(make([]byte, msgSize))
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(Message)

	if err != nil {
		return bytes.NewBuffer(make([]byte, 0)), errors.New("Could not encode message")
	}

	fmt.Println(buffer)

	return buffer, nil
}

func PerformEncoding() {
	flags := models.HeaderFlags{
		RD: 1,
	}

	header := models.DNSHeader{
		ID:      22,
		Flags:   flags,
		QDCount: 0,
		ANCount: 1,
		NSCount: 0,
		ARCount: 0,
	}

	encodedName, err := EncodeNameString("dns.google.com")
	if err != nil {
		fmt.Println(encodedName)
		return
	}

	question := models.DNSQuestion{
		QName: encodedName,
		QType: 1,
	}

	message := models.DNSMessage{
		Header:   header,
		Question: question,
	}

	buff, err := encodeMessage(message)

	if err != nil {
		return
	}
	fmt.Println(buff)
}
