// package hamming contains functions for working with DNA strands
package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strands are not equal!")
	}
	distance := 0
	strandA, strandB := []rune(a), []rune(b)
	for index := range strandA {
		if strandA[index] != strandB[index] {
			distance++
		}
	}
	return distance, nil
}
