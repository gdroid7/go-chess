package util

import (
	"fmt"
	"strings"
)

func ToChar(i int) string {
	//Get ascii characters
	return fmt.Sprintf("%s", string('A'+i))
}

func ToInt(c byte) int {
	//Get ascii characters
	return int(byte(c)) - int('A')
}

func SanitizeInputs(arr []string) (string, string) {
	return strings.Trim(strings.ToUpper(arr[0]), ","), strings.Trim(strings.ToUpper(arr[1]), ",")
}

func CoordinatesForHumans(moves [][]int) string {
	var result string
	for i := 0; i < len(moves); i++ {
		result += fmt.Sprintf("%s%d,", ToChar(moves[i][1]), moves[i][0]+1)
	}
	return strings.TrimRight(result, ",")
}
