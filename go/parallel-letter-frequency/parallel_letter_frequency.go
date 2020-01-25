// Package letter implements letter frequency counter using
// parallel computation.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// ConcurrentFrequency counts the frequency of letters
func ConcurrentFrequency(source []string) FreqMap {
	ch := make(chan FreqMap, 10)

	for _, text := range source {
		go func(word string) {
			ch <- Frequency(word)
		}(text)
	}

	countMap := FreqMap{}

	for range source {
		for letter, count := range <-ch {
			countMap[letter] += count
		}
	}
	return countMap
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
