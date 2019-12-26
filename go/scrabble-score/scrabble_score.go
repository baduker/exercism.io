// Package scrabble deals with proper word score counting
package scrabble

// Score count Scrabble score
func Score(word string) int {
	var score int
	for i := 0; i < len(word); i++ {
		score += int(scores[word[i]])
	}
	return score
}

var scores = [256]byte{
	' ': 0, // blank
	'A': 1, 'a': 1,
	'E': 1, 'e': 1,
	'I': 1, 'i': 1,
	'O': 1, 'o': 1,
	'U': 1, 'u': 1,
	'L': 1, 'l': 1,
	'N': 1, 'n': 1,
	'R': 1, 'r': 1,
	'S': 1, 's': 1,
	'T': 1, 't': 1,
	'D': 2, 'd': 2,
	'G': 2, 'g': 2,
	'B': 3, 'b': 3,
	'C': 3, 'c': 3,
	'M': 3, 'm': 3,
	'P': 3, 'p': 3,
	'F': 4, 'f': 4,
	'H': 4, 'h': 4,
	'V': 4, 'v': 4,
	'W': 4, 'w': 4,
	'Y': 4, 'y': 4,
	'K': 5, 'k': 5,
	'J': 8, 'j': 8,
	'X': 8, 'x': 8,
	'Q': 10, 'q': 10,
	'Z': 10, 'z': 10,
}
