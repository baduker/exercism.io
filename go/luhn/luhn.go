// Package luhn implements the Luhn algotithm
package luhn

import (
	"strconv"
	"strings"
)

// Valid determines whether or not a string of digits
// is valid per the Luhn formula
func Valid(input string) bool {
	var sum int
	var odd bool

	// Sanitize the input by removing spaces
	noWhiteSpace := strings.Replace(input, " ", "", -1)
	cardNumber := []rune(noWhiteSpace)

	// Anything shorter than two chars is invalid
	if len(cardNumber) < 2 {
		return false
	}

	// Start looping the string from the right side
	for i := len(cardNumber) - 1; i > -1; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		// Error means there's a char that's not a digit, bail out
		if err != nil {
			return false
		}
		if odd {
			digit *= 2
			if digit > 9 {
				digit = (digit % 10) + 1
			}
		}
		// Turns the odd swtich on & off
		odd = !odd
		sum += digit
	}
	return sum%10 == 0
}
