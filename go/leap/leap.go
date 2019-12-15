package leap

// IsLeapYear checks if a year is leap or not
func IsLeapYear(year int) bool {
	return year % 100 != 0 && year % 4 == 0 || year % 400 == 0
}
