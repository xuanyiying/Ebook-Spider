package utils

import "strings"

func ContainsAny(s string, arr []string) bool {
	for _, e := range arr {
		if strings.Contains(s, e) {
			return true
		}
	}
	return false
}
func HasPrefixAny(s string, arr []string) bool {
	for _, e := range arr {
		if strings.HasPrefix(s, e) {
			return true
		}
	}
	return false
}
