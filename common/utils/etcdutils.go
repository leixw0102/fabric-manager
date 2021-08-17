package utils

import "strings"

func GetCmd(key []byte) string {
	// sample key: fabric-manager/server/reportIP
	event := string(key)
	parts := strings.Split(event, "/")
	return parts[len(parts)-1]
}

func GetParams(value []byte) string {
	return string(value)
}
