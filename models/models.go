package models

import (
	"errors"
	"fmt"
	"math/bits"
)

func checkBitLength(desired int, got uint) error {
	if bits.Len(got) != desired {
		return errors.New(fmt.Sprintf("Wanted %v bits got %v bits", desired, bits.Len(got)))
	}
	return nil
}

type HeaderFlags struct {
	QR uint //Query(0) or Response(1)
	RD uint //Recursion Desired
	RA uint //Recursion Available
	TC uint //Truncation

}

type DNSHeader struct {
	ID      uint16
	Flags   HeaderFlags
	OPCode  uint //A 4-bit field
	AA      string
	QDCount uint16 //Number of queries in the answer Section
	ANCount uint16 //Number of resource records in the Answer Section
	NSCount uint16 //Number of name servers in the Authority Records section
	ARCount uint16 //Number of resource records in the Additional Record section
}

type DNSQuestion struct {
	QName  string // Encoded Domain Name
	QType  RecordType
	QClass ClassType
}
