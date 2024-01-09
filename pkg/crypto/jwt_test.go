package cryptography

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("123", "123", "equal")

}
