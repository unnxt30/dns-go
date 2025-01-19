package cmd

import (
	"errors"
	"strings"
)

func EncodeNameString(s string) ([]byte, error) {
	split_string := strings.Split(s, ".")
	var encoded []byte
	for _, part := range split_string {
		if len(part) > 63 {
			return nil, errors.New("label too long")
		}

		encoded = append(encoded, byte(len(part)))
		encoded = append(encoded, part...)

	}

	encoded = append(encoded, 0)
	return encoded, nil
}
