package cmd

import (
	"bytes"
	"encoding/binary"

	"github.com/unnxt30/dns-go/models"
)

func EncodeMessage(msg models.DNSMessage) ([]byte, error) {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.BigEndian, msg.Header.ID)
	binary.Write(buff, binary.BigEndian, msg.Header.Flags.Pack())
	binary.Write(buff, binary.BigEndian, msg.Header.QDCount)
	binary.Write(buff, binary.BigEndian, msg.Header.ANCount)
	binary.Write(buff, binary.BigEndian, msg.Header.NSCount)
	binary.Write(buff, binary.BigEndian, msg.Header.ARCount)

	qname, err := EncodeNameString(msg.Question.QName)
	if err != nil {
		return nil, err
	}
	buff.Write(qname)
	binary.Write(buff, binary.BigEndian, msg.Question.QType)
	binary.Write(buff, binary.BigEndian, msg.Question.QClass)
	return buff.Bytes(), nil
}
