package cmd

import (
	"encoding/binary"
	"fmt"
	"net"
	"reflect"

	"github.com/unnxt30/dns-go/models"
)

type DNSDecoder struct {
    Encoded []byte
    Offset  int
	NSRecords []string
	IPRecords []string
}

func (d *DNSDecoder) print(x ...interface{}) {
    fmt.Println(x...)
}

func (d *DNSDecoder) DecodeHeader() (models.DNSHeader, error) {
    var header models.DNSHeader

    if len(d.Encoded) < d.Offset+12 {
        return header, fmt.Errorf("not enough bytes to decode header")
    }

    header.ID = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2
    flags := binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2
    header.Flags = models.UnpackFlags(flags)
    header.QDCount = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2
    header.ANCount = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2
    header.NSCount = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2
    header.ARCount = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2

    d.PrintStructFields(header)

    return header, nil
}

func (d *DNSDecoder) DecodeQuestion() (models.DNSQuestion, error) {
    var question models.DNSQuestion

    name := ""
    for d.Encoded[d.Offset] != 0 {
        length := int(d.Encoded[d.Offset])
        d.Offset++
        if len(name) > 0 {
            name += "."
        }
        name += string(d.Encoded[d.Offset : d.Offset+length])
        d.Offset += length
    }
    d.Offset++

    question.QName = name
    if len(d.Encoded) < d.Offset+4 {
        return question, fmt.Errorf("not enough bytes to decode question")
    }
    question.QType = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2
    question.QClass = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2

    d.PrintStructFields(question)

    return question, nil
}

func (d *DNSDecoder) DecodeAnswers(anCount int) ([]models.ResponseStruct, error) {
    var answers []models.ResponseStruct

    for i := 0; i < anCount; i++ {
        answer, err := d.DecodeAnswer()
        if err != nil {
            return nil, err
        }
        answers = append(answers, answer)
    }

    return answers, nil
}

func (d *DNSDecoder) DecodeAnswer() (models.ResponseStruct, error) {
	// fmt.Println(d.Encoded[d.Offset:])
    var answer models.ResponseStruct

	name, err := d.decodeName()
	if err != nil {
		return answer, err
	}

    answer.Name = name

    // Decode the rest of the fields
    if len(d.Encoded) < d.Offset+10 {
        return answer, fmt.Errorf("not enough bytes to decode answer")
    }
    answer.Type = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2
    answer.Class = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2
    answer.TTL = binary.BigEndian.Uint32(d.Encoded[d.Offset : d.Offset+4])
    d.Offset += 4
    answer.RDLength = binary.BigEndian.Uint16(d.Encoded[d.Offset : d.Offset+2])
    d.Offset += 2

    if len(d.Encoded) < d.Offset+int(answer.RDLength) {
        return answer, fmt.Errorf("not enough bytes to decode answer data")
    }

    // Decode the data based on the type
    switch answer.Type {
    case 1: // A record
        answer.RData = net.IP(d.Encoded[d.Offset : d.Offset+int(answer.RDLength)]).String()
        d.Offset += int(answer.RDLength)
		d.IPRecords = append(d.IPRecords, answer.RData)
    case 2: // NS Record
		nsName, err := d.decodeName()
		if err != nil {
			return answer, fmt.Errorf("failed to decode NS record: %v", err)
		}
		answer.RData = nsName
		d.NSRecords = append(d.NSRecords, nsName)
    default:
        answer.RData = string(d.Encoded[d.Offset : d.Offset+int(answer.RDLength)])
        d.Offset += int(answer.RDLength)
    }

    return answer, nil
}

func (d *DNSDecoder) PrintStructFields(v interface{}) {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)

    for i := 0; i < typ.NumField(); i++ {
        field := typ.Field(i)
        value := val.Field(i).Interface()

        if field.Type.Kind() == reflect.Struct {
            fmt.Printf("Nested Struct: %s\n", field.Name)
            d.PrintStructFields(value)
        } else {
            switch field.Name {
            case "QClass":
                value = models.ClassType(value.(uint16))
                d.print("ClassType", value)
            case "QType":
                value = models.RecordType(value.(uint16))
                d.print("RecordType", value)
            default:
                d.print(field.Name, value)
            }
        }
    }
}


func (d *DNSDecoder) decodeName() (string, error) {
    var name string
    originalOffset := d.Offset

    for {
        if d.Offset >= len(d.Encoded) {
            return "", fmt.Errorf("invalid name: offset out of bounds")
        }

        length := int(d.Encoded[d.Offset])
        if length == 0 {
            d.Offset++
            break
        }

        // Handle DNS name compression (pointer)
        if length >= 192 { // 192 = 0xC0
            if d.Offset+1 >= len(d.Encoded) {
                return "", fmt.Errorf("invalid compression pointer: offset out of bounds")
            }
            ptr := int(binary.BigEndian.Uint16(d.Encoded[d.Offset:d.Offset+2]) & 0x3FFF) 
            d.Offset += 2

            savedOffset := d.Offset
            d.Offset = ptr

            compressedName, err := d.decodeName()
            if err != nil {
                return "", err
            }
			if len(name) > 0 {
				name += "." + compressedName
			}else{
            	name += compressedName
			}

            d.Offset = savedOffset
            break
        }

        // Decode a label
        d.Offset++
        if d.Offset+length > len(d.Encoded) {
            return "", fmt.Errorf("invalid label: offset out of bounds")
        }
        if len(name) > 0 {
            name += "."
        }
        name += string(d.Encoded[d.Offset : d.Offset+length])
        d.Offset += length
    }

    // Reset the offset if no compression was encountered
    if name == "" {
        d.Offset = originalOffset
    }

    return name, nil
}
