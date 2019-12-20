package dna

import (
	"fmt"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]uint

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	count := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	d = DNA(strings.ToUpper(string(d)))
	for _, nucleotide := range d {
		_, IsValid := count[nucleotide]
		if !IsValid {
			return nil, fmt.Errorf("invalid nucleotide %v", nucleotide)
		}
		count[nucleotide]++
	}
	return count, nil
}
