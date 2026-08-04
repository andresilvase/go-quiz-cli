package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ToInt(value string) (int, error) {
	if strings.Contains(value, "\n") {
		value = value[:len(value)-1]
	}

	intValue, err := strconv.Atoi(value)

	if err != nil {
		return 0, errors.New("você precisa digitar um valor numérico.\n")
	}

	return intValue, nil
}
