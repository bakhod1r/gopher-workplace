// Package lastindexbug — Gopher Workplace challenge.
package lastindexbug

// LastIndex returns the index of the last element equal to v, or -1.
//
// Examples:
//
//	LastIndex([]int{1, 2, 1}, 1) => 2
func LastIndex[T comparable](s []T, v T) int {
	// CHANGE CODE BELOW THIS LINE
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return i
		}
	}
	return -1
	// CHANGE CODE ABOVE THIS LINE
}
