// Package wordcount deals with string parsing
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a container for word occurance in a string
type Frequency map[string]int

// WordCount counts word occurrence in a string
func WordCount(text string) Frequency {
	count := Frequency{}
	pattern := regexp.MustCompile("\\w+('\\w+)?")
	for _, word := range pattern.FindAllString(text, -1) {
		count[strings.ToLower(word)]++
	}
	return count
}
