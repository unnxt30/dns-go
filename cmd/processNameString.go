package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func EncodeNameString(s string) (string, error) {
	split_string := strings.Split(s, ".")

	if len(split_string) <= 1 {
		return "", errors.New("Invalid Name-String")
	}

	var result_string string

	for _, v := range split_string {

		part_length := strconv.Itoa(len(v))

		result_string += part_length
		result_string += v

	}

	fmt.Println(result_string)

	return result_string, nil
}
