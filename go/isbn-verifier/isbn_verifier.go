// Package isbn implements logic to verify ISBN-10 strings
package isbn

// IsValidISBN performs the ISBN-10 verification process
func IsValidISBN(isbn string) bool {
	weight, sum := 10, 0
	for index, digit := range isbn {
		if digit >= '0' && digit <= '9' {
			sum += int(digit-'0') * weight
			weight--
		} else if index == len(isbn)-1 && digit == 'X' {
			sum += 10
			weight--
		}
	}
	return sum%11 == 0 && weight == 0
}
