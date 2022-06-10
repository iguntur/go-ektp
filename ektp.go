package ektp

import (
	"strconv"
)

func isValidRange(id string) bool {
	if len(id) == 16 {
		return true
	}

	return false
}

func isDigit(id string) bool {
	if _, err := strconv.Atoi(id); err != nil {
		return false
	}

	return true
}
