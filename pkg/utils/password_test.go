package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPass(t *testing.T) {
	pass := "testpassword"
	assert := assert.New(t)
	hashedPassword, err := HashPassword(pass)
	assert.Nil(err, "Hash success")
	assert.NotNil(hashedPassword, "Hasd password result")

	err = VerifyPassword(hashedPassword, pass)
	assert.Nil(err, "verify password")
}
