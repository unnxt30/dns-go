// package cmd

// import (
// 	"encoding/binary"
// 	"fmt"
// 	"reflect"

// 	"github.com/unnxt30/dns-go/models"
// )

// func print(x ...interface{}){
// 	fmt.Println(x...)
// }

// func printStructFields(v interface{}) {
//     val := reflect.ValueOf(v)
//     typ := reflect.TypeOf(v)

//     for i := 0; i < typ.NumField(); i++ {
//         field := typ.Field(i)
//         value := val.Field(i).Interface()

//         if field.Type.Kind() == reflect.Struct {
//             fmt.Printf("Nested Struct: %s\n", field.Name)
//             printStructFields(value)
//         } else {
//             switch field.Name {
//             case "QClass":
//                 value = models.ClassType(value.(uint16))
//                 print("ClassType", value)
//             case "QType":
//                 value = models.RecordType(value.(uint16))
//                 print("RecordType", value)
//             default:
//                 print(field.Name, value)
//             }
//         }
//     }
// }

// var Offset int = 0;

// func DecodeHeader(Encoded []byte) (string, error) {
// 	var header models.DNSHeader

// 	header.ID = binary.BigEndian.Uint16(Encoded[Offset:Offset + 2])
// 	Offset += 2
// 	flags := binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2
// 	header.Flags = models.UnpackFlags(flags)
// 	header.QDCount = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2
// 	header.ANCount = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2
// 	header.NSCount = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2
// 	header.ARCount = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2

// 	// headerType := reflect.TypeOf(header)
// 	// headerValue := reflect.ValueOf(header)

// 	// for i := 0; i < headerType.NumField(); i++ {
// 	// 	field := headerType.Field(i)
// 	// 	if field.Name == "Flags" {
// 	// 		flags := header.Flags
// 	// 		flagType := reflect.TypeOf(flags)
// 	// 		flagValue := reflect.ValueOf(flags)
// 	// 		for j := 0; j < flagType.NumField(); j++ {
// 	// 			flagField := flagType.Field(j)
// 	// 			value := flagValue.Field(j).Interface()
// 	// 			fmt.Println(flagField.Name, value)
// 	// 		}
// 	// 		continue
// 	// 	}
// 	// 	value := headerValue.Field(i).Interface()
// 	// 	fmt.Println(field.Name, value)
// 	// }

// 	printStructFields(header)

// 	return "", nil
// }

// func DecodeQuestion(Encoded []byte) (string, error) {
// 	var question models.DNSQuestion
// 	curr := Offset
// 	for Encoded[Offset] != 0 {
// 		Offset++
// 	}
// 	name := ""
// 	for i := curr; i < Offset; i++ {
// 		length := int(Encoded[i])
// 		if length == 0 {
// 			break
// 		}
// 		if len(name) > 0{
// 			name += "."
// 		}
// 		name += string(Encoded[i:i+length+1])
// 		i += length
// 	}
// 	Offset++

// 	question.QName = name
// 	question.QType = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2
// 	question.QClass = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2

// 	// questionType := reflect.TypeOf(question)
// 	// questionValue := reflect.ValueOf(question)

// 	// for i:=0; i<questionType.NumField(); i++ {
// 	// 	field := questionType.Field(i)
// 	// 	value := questionValue.Field(i).Interface()
// 	// 	if field.Name == "QClass"{
// 	// 		value = models.ClassType(value.(uint16))
// 	// 		print("ClassType", value)
// 	// 		continue
// 	// 	}
// 	// 	if field.Name == "QType"{
// 	// 		value = models.RecordType(value.(uint16))
// 	// 		print("RecordType", value)
// 	// 		continue
// 	// 	}
// 	// 	print(field.Name,value)
// 	// }

// 	printStructFields(question)

// 	return "", nil
// }

// func DecodeAnswer(Encoded []byte) (string, error) {

// 	//DNS Name encoding comes into place -> two octet encoding, starting with 11, the first octet and the following octet gives the Offset from the begining of the Encoded string.
// 	// which becomes 11000000 xxxxxxxx which is equivalent to [192 Offset]
// 	var answer models.ResponseStruct

// 	if int(Encoded[Offset]) == int(192){
// 		Offset ++
// 	}

// 	ptr := int(Encoded[Offset])

// 	curr := ptr

// 	for Encoded[ptr] != 0 {
// 		ptr++
// 	}
// 	ptr++
// 	name := ""

// 	for i:= curr; i<ptr; i++ {
// 		length := int(Encoded[i])
// 		if length == 0 {
// 			break
// 		}
// 		if len(name) > 0{
// 			name += "."
// 		}
// 		name += string(Encoded[i:i+length+1])
// 		i += length
// 	}
// 	Offset++
// 	answer.Name = name

// 	answer.Class = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2
// 	answer.Type = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2
// 	answer.TTL = binary.BigEndian.Uint32(Encoded[Offset:Offset+4])
// 	Offset += 4
// 	answer.RDLength = binary.BigEndian.Uint16(Encoded[Offset:Offset+2])
// 	Offset += 2
// 	fmt.Println(name)
// 	ip := ""

// 	// for i:=0; i < int(answer.RDLength); i++ {
// 	// 	if len(ip) > 0{
// 	// 		ip += "."
// 	// 	}
// 	// 	ip += strconv.Itoa(int(Encoded[Offset+i]))
// 	// }

// 	answer.RData = ip

//     // Decode the rest of the fields
// 	printStructFields(answer)
// 	return "", nil
// }

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

    d.printStructFields(header)

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

    d.printStructFields(question)

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

	for i := 0; i < len(answers); i++ {
		d.printStructFields(answers[i])
	}

    return answers, nil
}

func (d *DNSDecoder) DecodeAnswer() (models.ResponseStruct, error) {
    var answer models.ResponseStruct

    // Handle DNS name compression
    if int(d.Encoded[d.Offset]) == 192 {
        d.Offset++
    }

    ptr := int(d.Encoded[d.Offset])
	curr := ptr

	for d.Encoded[ptr] != 0 {
		ptr++
	}
	ptr++
	name := ""

	for i:= curr; i<ptr; i++ {
		length := int(d.Encoded[i])
		if length == 0 {
			break
		}
		if len(name) > 0{
			name += "."
		}
		name += string(d.Encoded[i:i+length+1])
		i += length
	}

    d.Offset++

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
    default:
        answer.RData = string(d.Encoded[d.Offset : d.Offset+int(answer.RDLength)])
    }
    d.Offset += int(answer.RDLength)

    d.printStructFields(answer)

    return answer, nil
}

func (d *DNSDecoder) printStructFields(v interface{}) {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)

    for i := 0; i < typ.NumField(); i++ {
        field := typ.Field(i)
        value := val.Field(i).Interface()

        if field.Type.Kind() == reflect.Struct {
            fmt.Printf("Nested Struct: %s\n", field.Name)
            d.printStructFields(value)
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
