package util

import (
	"fmt"
	"strings"
)

func ToChar(i int) string {
	//Get ascii characters
	return fmt.Sprintf("%s", string('A'+i))
}

func SanitizeInputs(arr []string) []string {
	var sanitized []string
	for _, v := range arr {
		sanitized = append(sanitized, strings.ToUpper(v))
	}
	return sanitized
}
