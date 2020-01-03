package anagram

import (
	"reflect"
	"strings"
)

func countChars(word string) (count map[rune]int) {
	count = make(map[rune]int)
	word = strings.ToLower(word)
	for _, value := range word {
		if _, ok := count[value]; ok {
			count[value] += 1
		} else {
			count[value] = 1
		}
	}
	return
}

func Detect(word string, candidates []string) (anagram []string) {
	candidateOne := countChars(word)
	for _, item := range candidates {
		candidateTwo := countChars(item)
		if !reflect.DeepEqual(candidateOne, candidateTwo) {
			continue
		}
		anagram = append(anagram, item)
	}
	return
}
