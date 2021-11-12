package util

import (
	"fmt"
	"strings"
)

func ToChar(i int, j int) string {
	//Get ascii characters
	return fmt.Sprintf("%s%d", string('A'+i), j+1)
}

func SanitizeInputs(arr []string) []string {
	var sanitized []string
	for _, v := range arr {
		sanitized = append(sanitized, strings.ToUpper(v))
	}
	return sanitized
}
