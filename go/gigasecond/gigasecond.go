package gigasecond

import "time"

// AddGigasecond determins the moment that would be after
// a gigasecond has passed. A gigasecond is 10^9 (1,000,000,000) seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000	* time.Second)
}
