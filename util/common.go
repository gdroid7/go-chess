package util

import (
	"fmt"
	"strings"
)

func ToChar(i int) string {
	//Get ascii characters
	return fmt.Sprintf("%s", string('A'+i))
}

func SanitizeInputs(arr []string) (string, string) {
	return strings.Trim(strings.ToUpper(arr[0]), ","), strings.Trim(strings.ToUpper(arr[1]), ",")
}
