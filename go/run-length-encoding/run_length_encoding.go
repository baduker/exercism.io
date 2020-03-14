package encode

import (
	"strconv"
	"strings"
	"unicode"
)

type runLength struct {
	enc, dec func(string) string
}

func newRL() *runLength {
	return &runLength{
		enc: func(input string) string {
			var result strings.Builder
			for len(input) > 0 {
				letter := input[0]
				inputLength := len(input)
				input = strings.TrimLeft(input, string(letter))
				if counter := inputLength - len(input); counter > 1 {
					result.WriteString(strconv.Itoa(counter))
				}
				result.WriteString(string(letter))
			}
			return result.String()
		},
		dec: func(input string) string {
			var result strings.Builder
			for len(input) > 0 {
				index := strings.IndexFunc(input, func(r rune) bool {return !unicode.IsDigit(r)})
				n := 1
				if index != 0 {
					n, _ = strconv.Atoi(input[:index])
				}
				result.WriteString(strings.Repeat(string(input[index]), n))
				input = input[index+1:]
			}
			return result.String()
		}}
}

func (rle runLength) encode(input string) string {
	return rle.enc(input)
}

func (rle runLength) decode(input string) string {
	return rle.dec(input)
}

func RunLengthEncode(input string) string {
	return newRL().encode(input)
}

func RunLengthDecode(input string) string {
	return newRL().decode(input)
}
