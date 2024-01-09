package utils

import (
	"fmt"
	"strings"
)

func RemoveOx(str string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("please provide string to remove 0x")
	}
	index := strings.Index(str, "0x")
	if index == -1 {
		return str, nil
	}
	result := str[2:]
	return result, nil
}
