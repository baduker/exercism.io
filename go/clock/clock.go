// Package clock implements a simple clock that handles
// time without dates but supports addition and subtraction
package clock

import "fmt"

// Clock is a data container that abstracts time ;)
type Clock struct {
	minutes int
}

// New initializes an instance of Clock with hours & minutes
func New(hours, minutes int) Clock {
	var c Clock
	timeInMinutes := hours*60 + minutes
	return c.Add(timeInMinutes)
}

// String shows time in the HH:MM format
func (c Clock) String() string {
	hours, minutes := c.minutes/60, c.minutes%60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// Add adds minutes to Clock
func (c Clock) Add(minutes int) Clock {
	c.minutes += minutes
	c.minutes %= 24 * 60
	// handle negative time
	if c.minutes < 0 {
		c.minutes += 24 * 60
	}
	return c
}

// Subtract subtracts minutes from Clock
// using Add's logic but with negative sign
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
