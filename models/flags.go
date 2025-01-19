package models

type HeaderFlags struct {
	QR     uint16 //Query(0) or Response(1)
	Opcode uint16
	AA     uint16
	RD     uint16 //Recursion Desired
	RA     uint16 //Recursion Available
	TC     uint16 //Truncation
	Z      uint16
	RCode  uint16
}

func (hf HeaderFlags) Pack() uint16 {
	var flags uint16
	flags |= hf.QR << 15
	flags |= hf.Opcode << 11
	flags |= hf.AA << 10
	flags |= hf.TC << 9
	flags |= hf.RD << 8
	flags |= hf.RA << 7
	flags |= hf.Z << 4
	flags |= hf.RCode

	return flags
}
