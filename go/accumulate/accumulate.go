// Package accumulate implements the Accumulate function
package accumulate

// Accumulate applies the converter function to each element
// of an array of strings and returns a new slice.
func Accumulate(collection []string, converter func(string) string) []string {
	var result []string
	for _, item := range collection {
		result = append(result, converter(string(item)))
	}
	return result
}
