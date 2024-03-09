package common

import "strings"

func Object(s string) bool {
	if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		return true
	}
	return false
}

func Array(s string) bool {
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		return true
	}
	return false
}
