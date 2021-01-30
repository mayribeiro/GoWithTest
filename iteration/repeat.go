package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, replace string, count int) string {
	var newString string

	if count == 0 {
		count = repeatCount
	}

	if replace != "" {
		newString = strings.Replace(character, character, replace, count)
		return strings.Repeat(newString, count)
	}

	return strings.Repeat(character, count)
}
