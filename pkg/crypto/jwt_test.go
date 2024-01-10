package cryptography

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	assert := assert.New(t)
	var durationTime time.Duration = 2 * time.Minute
	key := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlCUEFJQkFBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VXNjbGFFKzlaUUg5Q2VpOGIxcUVmCnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUUpCQUw4ZjRBMUlDSWEvQ2ZmdWR3TGMKNzRCdCtwOXg0TEZaZXMwdHdtV3Vha3hub3NaV0w4eVpSTUJpRmI4a25VL0hwb3piTnNxMmN1ZU9wKzVWdGRXNApiTlVDSVFENm9JdWxqcHdrZTFGY1VPaldnaXRQSjNnbFBma3NHVFBhdFYwYnJJVVI5d0loQVBOanJ1enB4ckhsCkUxRmJxeGtUNFZ5bWhCOU1HazU0Wk1jWnVjSmZOcjBUQWlFQWhML3UxOVZPdlVBWVd6Wjc3Y3JxMTdWSFBTcXoKUlhsZjd2TnJpdEg1ZGdjQ0lRRHR5QmFPdUxuNDlIOFIvZ2ZEZ1V1cjg3YWl5UHZ1YStxeEpXMzQrb0tFNXdJZwpQbG1KYXZsbW9jUG4rTkVRdGhLcTZuZFVYRGpXTTlTbktQQTVlUDZSUEs0PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ=="
	pubKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VQpzY2xhRSs5WlFIOUNlaThiMXFFZnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ=="

	result, err := CreateToken(durationTime, "Test", key)
	assert.Nil(err)
	assert.NotNil(result)

	data, err := ValidateToken(result, pubKey)
	assert.Nil(err)
	assert.Equal(data, "Test", "Test verify signature")
}
