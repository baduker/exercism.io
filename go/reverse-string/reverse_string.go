package reverse

import "bytes"

// String returns the input string in reverse mode
func Reverse(input string) string {
	var output bytes.Buffer

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		output.WriteRune(runes[i])
	}

	return output.String()
}
