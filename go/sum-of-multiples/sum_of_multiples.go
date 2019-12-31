// Package summultiples implements SumMultiples
package summultiples

// SumMultiples finds the sum of all the unique multiples
// of particular numbers up to but not including that number.
func SumMultiples(limit int, divisors ... int) (sum int) {
	// Perform initial check
	if limit < 2 || len(divisors) == 0 {
		return 0
	}
	// Keep tabs on checked integers
	multiples := make(map[int]bool)
	for _, divisor := range divisors {
		if divisor > 0 {
			for integer := divisor; integer < limit; integer += divisor {
				if !multiples[integer] {
					sum += integer
					multiples[integer] = true
				}
			}
		}
	}
	return sum
}
