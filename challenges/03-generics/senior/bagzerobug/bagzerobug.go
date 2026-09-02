// Package bagzerobug — Gopher Workplace challenge.
package bagzerobug

// Count tallies how many times each value appears.
//
// Examples:
//
//	Count([]string{"a", "a"}) => map[a:2]
func Count[T comparable](s []T) map[T]int {
	// CHANGE CODE BELOW THIS LINE
	m := make(map[T]int, len(s))
	for _, v := range s {
		m[v] = 1
	}
	return m
	// CHANGE CODE ABOVE THIS LINE
}
