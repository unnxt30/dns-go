package models

type ClassType uint16

const (
	IN ClassType = 1 // Internet
	CS ClassType = 2 // CSNET class
	CH ClassType = 3 // CHAOS class
	HS ClassType = 4 // Hesiod
)

func (ct ClassType) String() string {
	switch ct {
	case IN:
		return "IN"
	case CS:
		return "CS"
	case CH:
		return "CH"
	case HS:
		return "HS"
	default:
		return "UNKNOWN"

	}
}
