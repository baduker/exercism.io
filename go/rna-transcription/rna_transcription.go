// Package strand has functions to work with DNA strands
package strand

import "strings"

// ToRNA given a DNA strand, return its RNA complement
func ToRNA(dna string) string {
	RnaMap := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	var transcribed []string
	for _, nucleotide := range dna {
		transcribed = append(transcribed, RnaMap[string(nucleotide)])
	}
	return strings.Join(transcribed, "")
}
