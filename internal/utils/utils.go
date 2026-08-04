package utils

import (
	"strconv"
	"strings"

	myErrors "github.com/andresilvase/go-quiz-cli/internal/errors"
)

func ToInt(value string) (int, error) {
	value = strings.TrimSpace(value)

	intValue, err := strconv.Atoi(value)

	if err != nil {
		return 0, myErrors.ExpectInteger{Message: "você precisa digitar um valor numérico. "}
	}

	return intValue, nil
}
