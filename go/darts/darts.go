package darts

import "math"

// Score calcualtes dart points based on dart location (x, y)
func Score(x, y float64) int {
	targetPoint := math.Pow(x, 2) + math.Pow(y, 2)
	switch {
	case targetPoint <= GetSquare(1):
		return 10
	case targetPoint <= GetSquare(5):
		return 5
	case targetPoint <= GetSquare(10):
		return 1
	default:
		return 0
	}
}

// GetSquare returns radius value squared
func GetSquare(r float64) float64 {
	return math.Pow(r, 2)
}
