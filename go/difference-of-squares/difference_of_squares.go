// Package diffsquares finds the difference between
// the square of the sum and the sum of the squares
// of the first N natural numbers.
package diffsquares

// SquareOfSum calculates the square of the sum
// of the first ten natural numbers
func SquareOfSum(number int) int {
	sum := (number * (number + 1)) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares
// of the first ten natural numbers
func SumOfSquares(number int) int {
	return (number * (number + 1) * (2 * number + 1)) / 6
}

// Difference returns the difference between the square
// of the sum of the first ten natural numbers and the
// sum of the squares of the first ten natural numbers.
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
