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

func Map(a any) bool {
	if _, ok := a.(map[string]any); ok {
		return true
	}
	return false
}
