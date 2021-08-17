package utils

import "strings"

func GetCmd(key []byte) string {
	// event: fabric-manager/server/reportIP
	event := string(key)
	parts := strings.Split(event, "/")
	return parts[len(parts)-1] //createOrg
}

func GetParams(value []byte) string {
	return string(value)
}
