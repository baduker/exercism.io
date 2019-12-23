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
    'a': 1,
    'e': 1,
    'i': 1,
    'o': 1,
    'u': 1,
    'l': 1,
    'n': 1,
    'r': 1,
    's': 1,
    't': 1,
    'A': 1,
    'E': 1,
    'I': 1,
    'O': 1,
    'U': 1,
    'L': 1,
    'N': 1,
    'R': 1,
    'S': 1,
    'T': 1,

    'd': 2,
    'g': 2,
    'D': 2,
    'G': 2,

    'b': 3,
    'c': 3,
    'm': 3,
    'p': 3,
    'B': 3,
    'C': 3,
    'M': 3,
    'P': 3,

    'f': 4,
    'h': 4,
    'v': 4,
    'w': 4,
    'y': 4,
    'F': 4,
    'H': 4,
    'V': 4,
    'W': 4,
    'Y': 4,

    'k': 5,
    'K': 5,

    'j': 8,
    'x': 8,
    'J': 8,
    'X': 8,

    'q': 10,
    'z': 10,
    'Q': 10,
    'Z': 10,
}
