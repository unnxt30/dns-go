package cmd

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/unnxt30/dns-go/models"
)

func print(x ...interface{}){
	fmt.Println(x...)
}

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

func DecodeQuestion(encoded []byte) (string, error) {
	var question models.DNSQuestion
	qnameEnd := 12
	for encoded[qnameEnd] != 0 {
		qnameEnd++
	}
	name := ""
	for i := 12; i < qnameEnd; i++ {
		length := int(encoded[i])
		if length == 0 {
			break
		}
		if len(name) > 0{
			name += "."
		}
		name += string(encoded[i:i+length+1])
		i += length
	}
	qnameEnd++	

	question.QName = name
	typeEnd := qnameEnd + 2 
	question.QType = binary.BigEndian.Uint16(encoded[qnameEnd:typeEnd])
	classEnd := typeEnd + 2
	question.QClass = binary.BigEndian.Uint16(encoded[typeEnd:classEnd])


	questionType := reflect.TypeOf(question)
	questionValue := reflect.ValueOf(question)

	for i:=0; i<questionType.NumField(); i++ {
		field := questionType.Field(i)
		value := questionValue.Field(i).Interface()
		if field.Name == "QClass"{
			value = models.ClassType(value.(uint16))
			print("ClassType", value)
			continue
		}
		if field.Name == "QType"{
			value = models.RecordType(value.(uint16))
			print("RecordType", value)
			continue
		}
		print(field.Name,value)
	}

	return "", nil
}
