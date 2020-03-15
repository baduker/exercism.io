package anagram

import (
	"strings"
	"unicode"
)


func countChars(word string) (count [26]rune) {
	word = strings.ToLower(word)
	for _, char := range word {
		if unicode.IsLetter(char) {
			count[char-'a']++
		}
	}
	return
}

func Detect(word string, candidates []string) (anagram []string) {
	source := countChars(word)
	for _, candidate := range candidates {
		if len(word) != len(candidate) || strings.EqualFold(word, candidate) {
			continue
		}
		candidateTwo := countChars(candidate)
		if source == candidateTwo {
			anagram = append(anagram, candidate)
		}
	}
	return
}
