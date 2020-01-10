// Package armstrong implements IsNumber to check if a number
// is an Armstrong number. An Armstrong number is a number that
// is the sum of its own digits each raised to the power of
// the number of digits.
package armstrong

import (
	"math"
	"strconv"
)

// IsNumber checks if candidate is an Armstrong number
func IsNumber(candidate int) (isArmstrong bool) {
	length := strconv.Itoa(candidate)
	power := len(length)
	result := 0
	for _, char := range length {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return
		}
		result += int(math.Pow(float64(digit), float64(power)))
	}
	if candidate == result {
		isArmstrong = true
	}
	return
}
