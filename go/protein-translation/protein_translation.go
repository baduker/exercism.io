package protein

import "errors"

// Custom error for protein termination and invalid codon
var (
	ErrStop        = errors.New("protein translation terminated")
	ErrInvalidBase = errors.New("invalid codon")
)

// FromRNA translates RNA sequences into proteins
func FromRNA(strand string) ([]string, error) {
	proteins := make([]string, 0, len(strand)/3)
	for i := 0; i < len(strand); i += 3 {
		protein, err := FromCodon(strand[i : i+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

// FromCodon converts RNA codons into corresponding proteins
// Should the codon be a 'STOP' codon or the codon is invalid
// a correspodning error is returned
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
