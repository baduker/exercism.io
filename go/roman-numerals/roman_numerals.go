// Package romannumerals deals with encoding arabic numerlas to roman numerals
package romannumerals

import "fmt"

var (
	r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	r3 = []string{"", "M", "MM", "MMM"}
)

// getRoman checks if an arabic numeral can be represented as a roman numeral
// then it calculates the result
func getRoman(num int) (string, bool) {
	if num <= 0 || num > 3000 {
		return "", false
	}
	result := r3[num%10000/1000] + r2[num%1000/100] + r1[num%100/10] + r0[num%10]
	return result, true
}

// ToRomanNumeral coverts arabic numerlas to roman representation
func ToRomanNumeral(arabic int) (string, error) {
	romanNumeral, valid := getRoman(arabic)
	if valid {
		return romanNumeral, nil
	} else {
		err := fmt.Errorf("impossible to represent: %v", arabic)
		return "", err
	}
}
