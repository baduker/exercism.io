// Package raindrops deals with converting numbers to strings
package raindrops

import (
	"strconv"
	// "strings"
	// "github.com/AlasdairF/Sort/Int"
)

// Convert converts a number to a string based on the number's factors
func Convert(num int) (r string) {
	if num%3 == 0 {
		r += "Pling"
	}
	if num%5 == 0 {
		r += "Plang"
	}
	if num%7 == 0 {
		r += "Plong"
	}
	if r == "" {
		r = strconv.Itoa(num)
	}
	return
}
