package models

type RecordType uint16

const (
	A     RecordType = 1  // Host address
	NS    RecordType = 2  // Authoritative name server
	MD    RecordType = 3  // Mail destination (Obsolete - use MX)
	MF    RecordType = 4  // Mail forwarder (Obsolete - use MX)
	CNAME RecordType = 5  // Canonical name for an alias
	SOA   RecordType = 6  // Start of authority
	MB    RecordType = 7  // Mailbox domain name (EXPERIMENTAL)
	MG    RecordType = 8  // Mail group member (EXPERIMENTAL)
	MR    RecordType = 9  // Mail rename domain name (EXPERIMENTAL)
	NULL  RecordType = 10 // Null RR (EXPERIMENTAL)
	WKS   RecordType = 11 // Well known service description
	PTR   RecordType = 12 // Domain name pointer
	HINFO RecordType = 13 // Host information
	MINFO RecordType = 14 // Mailbox or mail list information
	MX    RecordType = 15 // Mail exchange
	TXT   RecordType = 16 // Text strings
)

func (rt RecordType) String() string {
	switch rt {
	case A:
		return "A"
	case NS:
		return "NS"
	case MD:
		return "MD"
	case MF:
		return "MF"
	case CNAME:
		return "CNAME"
	case SOA:
		return "SOA"
	case MB:
		return "MB"
	case MG:
		return "MG"
	case MR:
		return "MR"
	case NULL:
		return "NULL"
	case WKS:
		return "WKS"
	case PTR:
		return "PTR"
	case HINFO:
		return "HINFO"
	case MINFO:
		return "MINFO"
	case MX:
		return "MX"
	case TXT:
		return "TXT"
	default:
		return "UNKNOWN"
	}
}
