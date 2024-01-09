package utils

import "strings"

func RemoveOx(str string) string {
	if str == "" {
		panic("Please provide string to remove 0x")
	}
	index := strings.Index(str, "0x")
	if index != -1 {
		return str
	}
	result := str[1:]
	return result
}
