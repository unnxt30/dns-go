package models

type ResponseStruct struct {
	Name string
	Type RecordType
	ClassType ClassType
	TTL uint32
	RDLength uint16
	RData string
}