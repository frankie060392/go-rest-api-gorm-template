package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)
	test, _ := RemoveOx("0x12345")
	assert.Equal("12345", test, "Remove 0x")
	_, err := RemoveOx("")
	assert.EqualErrorf(err, "please provide string to remove 0x", "Empty hex")
	fmt.Println("end test")
}
