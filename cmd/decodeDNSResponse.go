package cmd

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/unnxt30/dns-go/models"
)

func DecodeHeader(encoded []byte) (string, error) {
	var header models.DNSHeader

	header.ID = binary.BigEndian.Uint16(encoded[:2])
	flags := binary.BigEndian.Uint16(encoded[2:4])
	header.Flags = models.UnpackFlags(flags)
	header.QDCount = binary.BigEndian.Uint16(encoded[4:6])
	header.ANCount = binary.BigEndian.Uint16(encoded[6:8])
	header.NSCount = binary.BigEndian.Uint16(encoded[8:10])
	header.ARCount = binary.BigEndian.Uint16(encoded[10:12])

	headerType := reflect.TypeOf(header)
	headerValue := reflect.ValueOf(header)

	for i := 0; i < headerType.NumField(); i++ {
		field := headerType.Field(i)
		if field.Name == "Flags" {
			flags := header.Flags
			flagType := reflect.TypeOf(flags)
			flagValue := reflect.ValueOf(flags)
			for j := 0; j < flagType.NumField(); j++ {
				flagField := flagType.Field(j)
				value := flagValue.Field(j).Interface()
				fmt.Println(flagField.Name, value)
			}
			continue
		}
		value := headerValue.Field(i).Interface()
		fmt.Println(field.Name, value)
	}

	return "", nil
}