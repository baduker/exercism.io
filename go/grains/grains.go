// Package grains implements logic for calculating the number
// of grains on a chessboard.
package grains

import "errors"

// Upper limit of grains on a chessboard
const maxGrains = uint64((1 << 64) - 1)

// Square returns the number of grains on a square of a chessboard
// given that the number of grains doubles on each square starting
// with one.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("input must be in range [1..64] only")
	}
	return 1 << uint(number - 1), nil
}

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
	return maxGrains
}
