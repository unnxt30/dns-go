package models

type ResponseStruct struct {
	Name string
	Type uint16 
	Class uint16 
	TTL uint32
	RDLength uint16
	RData string
}