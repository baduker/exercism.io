// Package strain implements the keep and discard operation on collections.
package strain

// Ints is a container for integer values
type Ints []int

// Lists is an array container of a list of integer values
type Lists [][]int

// Strings is a container for string values
type Strings []string

// Keep returns a list of items (ints) that were kept
// by predicate function
func (input Ints) Keep(predicate func(int) bool) (output Ints) {
	for _, item := range input {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return output
}

// Discard returns a list of items (ints) that were removed
// by predicate function
func (input Ints) Discard(predicate func(int) bool) (output Ints) {
	for _, item := range input {
		if !predicate(item) {
			output = append(output, item)
		}
	}
	return output
}

// Keep returns a list of lists with item (ints) that were kept
// by predicate function
func (input Lists) Keep(predicate func([]int) bool) (output Lists) {
	for _, item := range input {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return output
}

// Keep returns a string with items that were kept by predicate function
func (input Strings) Keep(predicate func(string) bool) (output Strings) {
	for _, item := range input {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return output
}
