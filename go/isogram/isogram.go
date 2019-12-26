// Package isogram checks if a string is an isogram.
// An isogram is a word or phrase without a repeating letter.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(text string) bool {
	charMap := make(map[rune]int)
	for _, char := range strings.ToLower(text) {
		if unicode.IsLetter(char) {
			charMap[char]++
			if charMap[char] > 1 {
				return false
			}
		}
	}
	return true
}
