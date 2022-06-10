package ektp

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidRangeID(t *testing.T) {
	assert.False(t, isValidRange(""), "Invalid when empty")

	var ids []string

	for i := 1; i < 16; i++ {
		ids = append(ids, strings.Repeat("1", i))
	}

	for _, id := range ids {
		assert.False(t, isValidRange(id), "Invalid range of "+strconv.Itoa(len(id)))
	}
}

func TestInvalidOutOfRangeID(t *testing.T) {
	var ids []string

	for i := 17; i < 20; i++ {
		ids = append(ids, strings.Repeat("1", i))
	}

	for _, id := range ids {
		assert.False(t, isValidRange(id), "Invalid out of range of "+strconv.Itoa(len(id)))
	}
}

func TestValidRange(t *testing.T) {
	assert.True(t, isValidRange(strings.Repeat("1", 16)), "Valid range")
}

func TestInvalidNonDigit(t *testing.T) {
	assert.False(t, isDigit("a"), "Invalid ID contains letter")
	assert.False(t, isDigit("a1"), "Invalid ID contains letter")
	assert.False(t, isDigit("1b"), "Invalid ID contains letter")
	assert.False(t, isDigit("1 8"), "Invalid ID contains space")
	assert.False(t, isDigit("1-9"), "Invalid ID contains symbol")
	assert.False(t, isDigit("19,1"), "Invalid ID contains symbol")
	assert.False(t, isDigit("19.2"), "Invalid ID contains symbol")
}

func TestValidDigit(t *testing.T) {
	assert.True(t, isDigit("1123"), "Valid input type")
}
