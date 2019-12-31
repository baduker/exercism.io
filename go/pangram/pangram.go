// Package pangram implements logic to check if a sentence
// is a pangram. A pangram (Greek: παν γράμμα, pan gramma,
// "every letter") is a sentence using every letter of
// the alphabet at least once.
package pangram

import "strings"

// IsPangram determines if a sentence is a pangram.
func IsPangram(text string) bool {
	text = strings.ToLower(strings.Replace(text, " ", "", -1))
	for char := 'a'; char <= 'z'; char++ {
		if !strings.ContainsRune(text, char) {
			return false
		}
	}
	return true
}
