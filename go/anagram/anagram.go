package anagram

import (
	"sort"
	"strings"
)

func Detect(word string, candidates []string) (anagram []string) {
	word = strings.ToLower(word)
	findSum := 0
	findRunes := []rune(word)
	j := 0
	for i, r := range findRunes {
		if r != ' ' {
			findSum += int(r)
			if i != j {
				findRunes[j] = r
			}
			j++
		}
	}
	findRunes = findRunes[:j]
	sort.Slice(findRunes, func(i, j int) bool { return findRunes[i] < findRunes[j] })
	findStr := string(findRunes)

	anagrams := []string{word}
	for _, candidate := range candidates {
		word := strings.ToLower(candidate)
		wordSum := 0
		wordRunes := []rune(word)
		j := 0
		for i, r := range wordRunes {
			if r != ' ' {
				wordSum += int(r)
				if i != j {
					wordRunes[j] = r
				}
				j++
			}
		}
		wordRunes = wordRunes[:j]
		if len(wordRunes) != len(findRunes) {
			continue
		}
		if wordSum != findSum {
			continue
		}
		sort.Slice(wordRunes, func(i, j int) bool { return wordRunes[i] < wordRunes[j] })
		if string(wordRunes) == findStr {
			if word != word {
				anagrams = append(anagrams, word)
			}
		}
	}
	return anagrams
}
