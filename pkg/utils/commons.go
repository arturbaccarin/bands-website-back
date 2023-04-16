package utils

import "strings"

// ConcatenateStrings concatenate a slice of strings
func ConcatenateStrings(values ...string) string {
	var sb strings.Builder

	for _, str := range values {
		sb.WriteString(str)
	}

	return sb.String()
}
