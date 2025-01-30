package models

type DNSHeader struct {
	ID      uint16
	Flags   HeaderFlags
	QDCount uint16 //Number of queries in the answer Section
	ANCount uint16 //Number of resource records in the Answer Section
	NSCount uint16 //Number of name servers in the Authority Records section
	ARCount uint16 //Number of resource records in the Additional Record section
}

type DNSQuestion struct {
	QName  string // Encoded Domain Name
	QType  uint16
	QClass uint16
}

type DNSMessage struct {
	Header   DNSHeader
	Question DNSQuestion
}
