// Package collatzconjecture solves the Collatz Conjecture challenge
package collatzconjecture

import "fmt"

// IsEven divides n by 2 to get n / 2, if n is even.
func IsEven(n int) int {
	return n / 2
}
// IsOdd multiplies n by 3 and adds 1 to get 3n + 1, if n is odd.
func IsOdd(n int) int {
	return n * 3 + 1
}
// CollatzConjecture given a number n,
// return the number of steps required to reach 1.
func CollatzConjecture(number int) (int, error) {
	countSteps := 0
	if number <= 0 {
		err := fmt.Errorf("invalid value: %q", number)
		return 0, err
	}
	for number > 1 {
		if number % 2 == 0 {
			number = IsEven(number)
		} else {
			number = IsOdd(number)
		}
		countSteps++
	}
	return countSteps, nil
}
