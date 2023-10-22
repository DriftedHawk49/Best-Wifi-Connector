package utilities

import (
	"strings"
)

// This function splits the string at splitter wherever it is not escaped in the string
func SplitEscapeSafeString(str string, splitter string) []string {
	result := make([]string, 0)
	j := 0

	for i, c := range str {
		if string(c) != splitter {

			if len(str)-1 == i {
				result = append(result, str[j:i+1])
			}
			continue
		}

		if i == 0 {
			result = append(result, "")
		}

		if string(str[i-1]) == ("\\") {
			continue
		}

		result = append(result, str[j:i])
		if i != len(str)-1 {
			j = i + 1
			continue
		}
		result = append(result, "")

	}

	return result
}

// removes escape characters from a string
func TrimEscapeCharacters(str string) string {
	sb := strings.Builder{}

	for _, s := range str {
		if string(s) == "\\" {
			continue
		}

		sb.WriteRune(s)
	}

	return sb.String()
}
