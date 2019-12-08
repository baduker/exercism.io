package twofer

import "strings"

func ShareWith(name string) string {
	result := "One for you, one for me."
	if name != "" {
		result = strings.Replace(result, "you", name, -1)
	}
	return result
}
