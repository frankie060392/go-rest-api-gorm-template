package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)
	test := RemoveOx("0x12345")
	assert.Equal("12345", test, "Remove 0x")
}
