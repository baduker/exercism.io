// Package raindrops deals with converting numbers to strings
package raindrops

import (
	"strconv"
	"strings"
)

// RainDrops converts a number to a string based on the number's factors
func Convert(number int) string {
	sounds := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	var speak []string

	for key, value := range sounds {
		if number % key == 0 {
			speak = append(speak, value)
		}
	}

	if len(speak) != 0 {
		return strings.Join(speak, "")
	}

	return strconv.Itoa(number)
}
